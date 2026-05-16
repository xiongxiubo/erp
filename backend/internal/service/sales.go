package service

import (
	"erp/internal/dto"
	"erp/internal/model"
	"erp/internal/response"
	"erp/internal/utils"
	"time"

	"gorm.io/gorm"
)

type SalesService struct {
	db    *gorm.DB
	stock *StockService
}

func NewSalesService(db *gorm.DB, stock *StockService) *SalesService {
	return &SalesService{db: db, stock: stock}
}

func (s *SalesService) CreateOrder(req dto.SalesOrderRequest) (model.SalesOrder, error) {
	orderDate := time.Now()
	if req.OrderDate != nil {
		orderDate = *req.OrderDate
	}
	order := model.SalesOrder{Code: req.Code, CustomerID: req.CustomerID, OrderDate: orderDate, Status: model.DocumentDraft}
	order.Remark = req.Remark
	if order.Code == "" {
		order.Code = utils.NewCode("SO")
	}
	for _, line := range req.Lines {
		amount := line.Quantity * line.Price
		order.TotalAmount += amount
		order.Lines = append(order.Lines, model.SalesOrderLine{ProductID: line.ProductID, WarehouseID: line.WarehouseID, Quantity: line.Quantity, Price: line.Price, Amount: amount})
	}
	return order, s.db.Create(&order).Error
}

func (s *SalesService) ApproveOrder(id uint) error {
	return s.updateOrderStatus(id, model.DocumentDraft, model.DocumentApproved, "只有草稿销售订单可以审批")
}

func (s *SalesService) CloseOrder(id uint) error {
	return s.db.Model(&model.SalesOrder{}).Where("id = ? and status in ?", id, []string{model.DocumentApproved, model.DocumentPartiallyOutbound}).Update("status", model.DocumentClosed).Error
}

func (s *SalesService) VoidOrder(id uint) error {
	return s.updateOrderStatus(id, model.DocumentDraft, model.DocumentVoided, "只有草稿销售订单可以作废")
}

func (s *SalesService) CreateOutbound(req dto.SalesOutboundRequest) (model.SalesOutbound, error) {
	date := time.Now()
	if req.OutboundDate != nil {
		date = *req.OutboundDate
	}
	outbound := model.SalesOutbound{Code: req.Code, OrderID: req.OrderID, CustomerID: req.CustomerID, OutboundDate: date, Status: model.DocumentDraft}
	outbound.Remark = req.Remark
	if outbound.Code == "" {
		outbound.Code = utils.NewCode("SOH")
	}
	for _, line := range req.Lines {
		amount := line.Quantity * line.Price
		outbound.TotalAmount += amount
		outbound.Lines = append(outbound.Lines, model.SalesOutboundLine{OrderLineID: line.OrderLineID, ProductID: line.ProductID, WarehouseID: line.WarehouseID, Quantity: line.Quantity, Price: line.Price, Amount: amount})
	}
	return outbound, s.db.Create(&outbound).Error
}

func (s *SalesService) ConfirmOutbound(id uint, operatorID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var outbound model.SalesOutbound
		if err := tx.Preload("Lines").First(&outbound, id).Error; err != nil {
			return response.NotFound("销售出库单不存在")
		}
		if outbound.Status != model.DocumentDraft {
			return response.BadRequest("只有草稿出库单可以确认")
		}
		var order model.SalesOrder
		if err := tx.Preload("Lines").First(&order, outbound.OrderID).Error; err != nil {
			return response.NotFound("销售订单不存在")
		}
		if order.Status != model.DocumentApproved && order.Status != model.DocumentPartiallyOutbound {
			return response.BadRequest("销售订单未审批或已关闭")
		}
		for _, line := range outbound.Lines {
			if line.WarehouseID == 0 {
				return response.BadRequest("出库明细必须选择仓库")
			}
			if err := s.stock.Decrease(tx, line.WarehouseID, line.ProductID, line.Quantity, line.Price, model.LedgerBizSalesOutbound, outbound.ID, outbound.Code, operatorID); err != nil {
				return err
			}
			if line.OrderLineID != 0 {
				if err := tx.Model(&model.SalesOrderLine{}).Where("id = ?", line.OrderLineID).Update("outbound_qty", gorm.Expr("outbound_qty + ?", line.Quantity)).Error; err != nil {
					return err
				}
			}
		}
		status := model.DocumentCompleted
		for _, line := range order.Lines {
			var refreshed model.SalesOrderLine
			if err := tx.First(&refreshed, line.ID).Error; err == nil && refreshed.OutboundQty < refreshed.Quantity {
				status = model.DocumentPartiallyOutbound
			}
		}
		if err := tx.Model(&outbound).Update("status", model.DocumentConfirmed).Error; err != nil {
			return err
		}
		if err := tx.Model(&order).Updates(map[string]any{"status": status, "outbound_amount": gorm.Expr("outbound_amount + ?", outbound.TotalAmount)}).Error; err != nil {
			return err
		}
		return tx.Create(&model.Receivable{SourceType: model.LedgerBizSalesOutbound, SourceID: outbound.ID, SourceCode: outbound.Code, CustomerID: outbound.CustomerID, Amount: outbound.TotalAmount, UnpaidAmount: outbound.TotalAmount, Status: model.FinanceUnpaid}).Error
	})
}

func (s *SalesService) VoidOutbound(id uint) error {
	return s.db.Model(&model.SalesOutbound{}).Where("id = ? and status = ?", id, model.DocumentDraft).Update("status", model.DocumentVoided).Error
}

func (s *SalesService) updateOrderStatus(id uint, from string, to string, message string) error {
	result := s.db.Model(&model.SalesOrder{}).Where("id = ? and status = ?", id, from).Update("status", to)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return response.BadRequest(message)
	}
	return nil
}
