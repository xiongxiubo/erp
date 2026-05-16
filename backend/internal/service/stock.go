package service

import (
	"erp/internal/model"
	"erp/internal/response"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StockService struct {
	db *gorm.DB
}

func NewStockService(db *gorm.DB) *StockService {
	return &StockService{db: db}
}

func (s *StockService) Increase(tx *gorm.DB, warehouseID uint, productID uint, quantity float64, unitCost float64, bizType string, bizID uint, bizCode string, operatorID uint) error {
	if quantity <= 0 {
		return response.BadRequest("入库数量必须大于 0")
	}
	stock, err := s.lockStock(tx, warehouseID, productID)
	if err != nil {
		return err
	}
	before := stock.Quantity
	amountBefore := stock.AverageCost * stock.Quantity
	amountIn := unitCost * quantity
	stock.Quantity += quantity
	if stock.Quantity > 0 {
		stock.AverageCost = (amountBefore + amountIn) / stock.Quantity
	}
	if err := tx.Save(&stock).Error; err != nil {
		return err
	}
	return tx.Create(&model.StockLedger{
		BizType:        bizType,
		BizID:          bizID,
		BizCode:        bizCode,
		Direction:      model.LedgerDirectionIn,
		WarehouseID:    warehouseID,
		ProductID:      productID,
		Quantity:       quantity,
		UnitCost:       unitCost,
		Amount:         amountIn,
		BeforeQuantity: before,
		AfterQuantity:  stock.Quantity,
		OccurredAt:     time.Now(),
		OperatorID:     operatorID,
	}).Error
}

func (s *StockService) Decrease(tx *gorm.DB, warehouseID uint, productID uint, quantity float64, unitCost float64, bizType string, bizID uint, bizCode string, operatorID uint) error {
	if quantity <= 0 {
		return response.BadRequest("出库数量必须大于 0")
	}
	stock, err := s.lockStock(tx, warehouseID, productID)
	if err != nil {
		return err
	}
	available := stock.Quantity - stock.LockedQuantity
	if available < quantity {
		return response.BadRequest("库存不足")
	}
	before := stock.Quantity
	stock.Quantity -= quantity
	if err := tx.Save(&stock).Error; err != nil {
		return err
	}
	cost := unitCost
	if cost == 0 {
		cost = stock.AverageCost
	}
	return tx.Create(&model.StockLedger{
		BizType:        bizType,
		BizID:          bizID,
		BizCode:        bizCode,
		Direction:      model.LedgerDirectionOut,
		WarehouseID:    warehouseID,
		ProductID:      productID,
		Quantity:       quantity,
		UnitCost:       cost,
		Amount:         cost * quantity,
		BeforeQuantity: before,
		AfterQuantity:  stock.Quantity,
		OccurredAt:     time.Now(),
		OperatorID:     operatorID,
	}).Error
}

func (s *StockService) lockStock(tx *gorm.DB, warehouseID uint, productID uint) (model.Stock, error) {
	var stock model.Stock
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("warehouse_id = ? and product_id = ?", warehouseID, productID).First(&stock).Error
	if err == nil {
		return stock, nil
	}
	if err != gorm.ErrRecordNotFound {
		return stock, err
	}
	stock = model.Stock{WarehouseID: warehouseID, ProductID: productID}
	return stock, tx.Create(&stock).Error
}
