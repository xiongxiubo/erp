package model

import "time"

type Stock struct {
	BaseModel
	WarehouseID    uint    `gorm:"uniqueIndex:idx_stock_product_warehouse;not null" json:"warehouseId"`
	ProductID      uint    `gorm:"uniqueIndex:idx_stock_product_warehouse;not null" json:"productId"`
	Quantity       float64 `json:"quantity"`
	LockedQuantity float64 `json:"lockedQuantity"`
	AverageCost    float64 `json:"averageCost"`
}

func (Stock) TableName() string { return "inv_stocks" }

type StockLedger struct {
	BaseModel
	BizType        string    `gorm:"size:80;index;not null" json:"bizType"`
	BizID          uint      `gorm:"index;not null" json:"bizId"`
	BizCode        string    `gorm:"size:80;index" json:"bizCode"`
	Direction      string    `gorm:"size:16;index;not null" json:"direction"`
	WarehouseID    uint      `gorm:"index;not null" json:"warehouseId"`
	ProductID      uint      `gorm:"index;not null" json:"productId"`
	Quantity       float64   `json:"quantity"`
	UnitCost       float64   `json:"unitCost"`
	Amount         float64   `json:"amount"`
	BeforeQuantity float64   `json:"beforeQuantity"`
	AfterQuantity  float64   `json:"afterQuantity"`
	OccurredAt     time.Time `gorm:"index" json:"occurredAt"`
	OperatorID     uint      `gorm:"index" json:"operatorId"`
}

func (StockLedger) TableName() string { return "inv_stock_ledgers" }
