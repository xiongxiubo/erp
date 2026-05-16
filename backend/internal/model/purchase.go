package model

import "time"

type PurchaseOrder struct {
	BaseModel
	Code          string              `gorm:"size:80;uniqueIndex;not null" json:"code"`
	SupplierID    uint                `gorm:"index;not null" json:"supplierId"`
	OrderDate     time.Time           `json:"orderDate"`
	Status        string              `gorm:"size:32;index;default:draft" json:"status"`
	TotalAmount   float64             `json:"totalAmount"`
	InboundAmount float64             `json:"inboundAmount"`
	Lines         []PurchaseOrderLine `gorm:"foreignKey:OrderID" json:"lines"`
}

func (PurchaseOrder) TableName() string { return "purchase_orders" }

type PurchaseOrderLine struct {
	BaseModel
	OrderID     uint    `gorm:"index;not null" json:"orderId"`
	ProductID   uint    `gorm:"index;not null" json:"productId"`
	WarehouseID uint    `gorm:"index" json:"warehouseId"`
	Quantity    float64 `json:"quantity"`
	InboundQty  float64 `json:"inboundQty"`
	Price       float64 `json:"price"`
	Amount      float64 `json:"amount"`
}

func (PurchaseOrderLine) TableName() string { return "purchase_order_lines" }

type PurchaseInbound struct {
	BaseModel
	Code        string                `gorm:"size:80;uniqueIndex;not null" json:"code"`
	OrderID     uint                  `gorm:"index;not null" json:"orderId"`
	SupplierID  uint                  `gorm:"index;not null" json:"supplierId"`
	InboundDate time.Time             `json:"inboundDate"`
	Status      string                `gorm:"size:32;index;default:draft" json:"status"`
	TotalAmount float64               `json:"totalAmount"`
	Lines       []PurchaseInboundLine `gorm:"foreignKey:InboundID" json:"lines"`
}

func (PurchaseInbound) TableName() string { return "purchase_inbounds" }

type PurchaseInboundLine struct {
	BaseModel
	InboundID   uint    `gorm:"index;not null" json:"inboundId"`
	OrderLineID uint    `gorm:"index" json:"orderLineId"`
	ProductID   uint    `gorm:"index;not null" json:"productId"`
	WarehouseID uint    `gorm:"index;not null" json:"warehouseId"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	Amount      float64 `json:"amount"`
}

func (PurchaseInboundLine) TableName() string { return "purchase_inbound_lines" }
