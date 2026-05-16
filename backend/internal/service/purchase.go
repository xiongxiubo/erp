package service

import (
	"erp/internal/dto"
	"erp/internal/model"
	"erp/internal/response"
	"erp/internal/utils"
	"time"

	"gorm.io/gorm"
)

type PurchaseService struct {
	db    *gorm.DB
	stock *StockService
}

func NewPurchaseService(db *gorm.DB, stock *StockService) *PurchaseService {
	return &PurchaseService{db: db, stock: stock}
}

func (s *PurchaseService) CreateOrder(req dto.PurchaseOrderRequest) (model.PurchaseOrder, error) {
	orderDate := time.Now()
	if req.OrderDate != nil {
		orderDate = *req.OrderDate
	}
	order := model.PurchaseOrder{Code: req.Code, SupplierID: req.SupplierID, OrderDate: orderDate, Status: model.DocumentDraft}
	order.Remark = req.Remark
	if order.Code == "" {
		order.Code = utils.NewCode("PO")
	}
	for _, line := range req.Lines {
		amount := line.Quantity * line.Price
		order.TotalAmount += amount
		order.Lines = append(order.Lines, model.PurchaseOrderLine{ProductID: line.ProductID, WarehouseID: line.WarehouseID, Quantity: line.Quantity, Price: line.Price, Amount: amount})
	}
	return order, s.db.Create(&order).Error
}

func (s *PurchaseService) UpdateOrder(id uint, req dto.PurchaseOrderRequest) error {
	var existing model.PurchaseOrder
	if err := s.db.Preload("Lines").First(&existing, id).Error; err != nil {
		return response.NotFound("采购订单不存在")
	}
	if existing.Status != model.DocumentDraft {
		return response.BadRequest("只有草稿采购订单可以编辑")
	}
	return s.db.Transaction(func(tx *gorm.DB) error {
		total := 0.0
		updates := map[string]any{"supplier_id": req.SupplierID, "remark": req.Remark}
		if req.OrderDate != nil {
			updates["order_date"] = *req.OrderDate
		}
		if err := tx.Model(&existing).Updates(updates).Error; err != nil {
			return err
		}
		if err := tx.Where("order_id = ?", id).Delete(&model.PurchaseOrderLine{}).Error; err != nil {
			return err
		}
		lines := make([]model.PurchaseOrderLine, 0, len(req.Lines))
		for _, line := range req.Lines {
			amount := line.Quantity * line.Price
			total += amount
			lines = append(lines, model.PurchaseOrderLine{OrderID: id, ProductID: line.ProductID, WarehouseID: line.WarehouseID, Quantity: line.Quantity, Price: line.Price, Amount: amount})
		}
		if len(lines) > 0 {
			if err := tx.Create(&lines).Error; err != nil {
				return err
			}
		}
		return tx.Model(&existing).Update("total_amount", total).Error
	})
}

func (s *PurchaseService) ApproveOrder(id uint) error {
	return s.updateOrderStatus(id, model.DocumentDraft, model.DocumentApproved, "只有草稿采购订单可以审批")
}

func (s *PurchaseService) CloseOrder(id uint) error {
	return s.db.Model(&model.PurchaseOrder{}).Where("id = ? and status in ?", id, []string{model.DocumentApproved, model.DocumentPartiallyInbound}).Update("status", model.DocumentClosed).Error
}

func (s *PurchaseService) VoidOrder(id uint) error {
	return s.updateOrderStatus(id, model.DocumentDraft, model.DocumentVoided, "只有草稿采购订单可以作废")
}

func (s *PurchaseService) CreateInbound(req dto.PurchaseInboundRequest) (model.PurchaseInbound, error) {
	date := time.Now()
	if req.InboundDate != nil {
		date = *req.InboundDate
	}
	inbound := model.PurchaseInbound{Code: req.Code, OrderID: req.OrderID, SupplierID: req.SupplierID, InboundDate: date, Status: model.DocumentDraft}
	inbound.Remark = req.Remark
	if inbound.Code == "" {
		inbound.Code = utils.NewCode("PI")
	}
	for _, line := range req.Lines {
		amount := line.Quantity * line.Price
		inbound.TotalAmount += amount
		inbound.Lines = append(inbound.Lines, model.PurchaseInboundLine{OrderLineID: line.OrderLineID, ProductID: line.ProductID, WarehouseID: line.WarehouseID, Quantity: line.Quantity, Price: line.Price, Amount: amount})
	}
	return inbound, s.db.Create(&inbound).Error
}

func (s *PurchaseService) ConfirmInbound(id uint, operatorID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var inbound model.PurchaseInbound
		if err := tx.Preload("Lines").First(&inbound, id).Error; err != nil {
			return response.NotFound("采购入库单不存在")
		}
		if inbound.Status != model.DocumentDraft {
			return response.BadRequest("只有草稿入库单可以确认")
		}
		var order model.PurchaseOrder
		if err := tx.Preload("Lines").First(&order, inbound.OrderID).Error; err != nil {
			return response.NotFound("采购订单不存在")
		}
		if order.Status != model.DocumentApproved && order.Status != model.DocumentPartiallyInbound {
			return response.BadRequest("采购订单未审批或已关闭")
		}
		for _, line := range inbound.Lines {
			if line.WarehouseID == 0 {
				return response.BadRequest("入库明细必须选择仓库")
			}
			if err := s.stock.Increase(tx, line.WarehouseID, line.ProductID, line.Quantity, line.Price, model.LedgerBizPurchaseInbound, inbound.ID, inbound.Code, operatorID); err != nil {
				return err
			}
			if line.OrderLineID != 0 {
				if err := tx.Model(&model.PurchaseOrderLine{}).Where("id = ?", line.OrderLineID).Update("inbound_qty", gorm.Expr("inbound_qty + ?", line.Quantity)).Error; err != nil {
					return err
				}
			}
		}
		status := model.DocumentCompleted
		for _, line := range order.Lines {
			var refreshed model.PurchaseOrderLine
			if err := tx.First(&refreshed, line.ID).Error; err == nil && refreshed.InboundQty < refreshed.Quantity {
				status = model.DocumentPartiallyInbound
			}
		}
		if err := tx.Model(&inbound).Update("status", model.DocumentConfirmed).Error; err != nil {
			return err
		}
		if err := tx.Model(&order).Updates(map[string]any{"status": status, "inbound_amount": gorm.Expr("inbound_amount + ?", inbound.TotalAmount)}).Error; err != nil {
			return err
		}
		return tx.Create(&model.Payable{SourceType: model.LedgerBizPurchaseInbound, SourceID: inbound.ID, SourceCode: inbound.Code, SupplierID: inbound.SupplierID, Amount: inbound.TotalAmount, UnpaidAmount: inbound.TotalAmount, Status: model.FinanceUnpaid}).Error
	})
}

func (s *PurchaseService) VoidInbound(id uint) error {
	return s.db.Model(&model.PurchaseInbound{}).Where("id = ? and status = ?", id, model.DocumentDraft).Update("status", model.DocumentVoided).Error
}

func (s *PurchaseService) updateOrderStatus(id uint, from string, to string, message string) error {
	result := s.db.Model(&model.PurchaseOrder{}).Where("id = ? and status = ?", id, from).Update("status", to)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return response.BadRequest(message)
	}
	return nil
}
