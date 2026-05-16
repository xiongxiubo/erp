package model

import "time"

type SalesOrder struct {
	BaseModel
	Code           string           `gorm:"size:80;uniqueIndex;not null" json:"code"`
	CustomerID     uint             `gorm:"index;not null" json:"customerId"`
	OrderDate      time.Time        `json:"orderDate"`
	Status         string           `gorm:"size:32;index;default:draft" json:"status"`
	TotalAmount    float64          `json:"totalAmount"`
	OutboundAmount float64          `json:"outboundAmount"`
	Lines          []SalesOrderLine `gorm:"foreignKey:OrderID" json:"lines"`
}

func (SalesOrder) TableName() string { return "sales_orders" }

type SalesOrderLine struct {
	BaseModel
	OrderID     uint    `gorm:"index;not null" json:"orderId"`
	ProductID   uint    `gorm:"index;not null" json:"productId"`
	WarehouseID uint    `gorm:"index" json:"warehouseId"`
	Quantity    float64 `json:"quantity"`
	OutboundQty float64 `json:"outboundQty"`
	Price       float64 `json:"price"`
	Amount      float64 `json:"amount"`
}

func (SalesOrderLine) TableName() string { return "sales_order_lines" }

type SalesOutbound struct {
	BaseModel
	Code         string              `gorm:"size:80;uniqueIndex;not null" json:"code"`
	OrderID      uint                `gorm:"index;not null" json:"orderId"`
	CustomerID   uint                `gorm:"index;not null" json:"customerId"`
	OutboundDate time.Time           `json:"outboundDate"`
	Status       string              `gorm:"size:32;index;default:draft" json:"status"`
	TotalAmount  float64             `json:"totalAmount"`
	Lines        []SalesOutboundLine `gorm:"foreignKey:OutboundID" json:"lines"`
}

func (SalesOutbound) TableName() string { return "sales_outbounds" }

type SalesOutboundLine struct {
	BaseModel
	OutboundID  uint    `gorm:"index;not null" json:"outboundId"`
	OrderLineID uint    `gorm:"index" json:"orderLineId"`
	ProductID   uint    `gorm:"index;not null" json:"productId"`
	WarehouseID uint    `gorm:"index;not null" json:"warehouseId"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	Amount      float64 `json:"amount"`
}

func (SalesOutboundLine) TableName() string { return "sales_outbound_lines" }
