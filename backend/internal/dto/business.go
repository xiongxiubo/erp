package dto

import "time"

type DocumentLineRequest struct {
	ID          uint    `json:"id"`
	OrderLineID uint    `json:"orderLineId"`
	ProductID   uint    `json:"productId" binding:"required"`
	WarehouseID uint    `json:"warehouseId"`
	Quantity    float64 `json:"quantity" binding:"required"`
	Price       float64 `json:"price"`
	Remark      string  `json:"remark"`
}

type PurchaseOrderRequest struct {
	Code       string                `json:"code"`
	SupplierID uint                  `json:"supplierId" binding:"required"`
	OrderDate  *time.Time            `json:"orderDate"`
	Remark     string                `json:"remark"`
	Lines      []DocumentLineRequest `json:"lines" binding:"required"`
}

type PurchaseInboundRequest struct {
	Code        string                `json:"code"`
	OrderID     uint                  `json:"orderId" binding:"required"`
	SupplierID  uint                  `json:"supplierId" binding:"required"`
	InboundDate *time.Time            `json:"inboundDate"`
	Remark      string                `json:"remark"`
	Lines       []DocumentLineRequest `json:"lines" binding:"required"`
}

type SalesOrderRequest struct {
	Code       string                `json:"code"`
	CustomerID uint                  `json:"customerId" binding:"required"`
	OrderDate  *time.Time            `json:"orderDate"`
	Remark     string                `json:"remark"`
	Lines      []DocumentLineRequest `json:"lines" binding:"required"`
}

type SalesOutboundRequest struct {
	Code         string                `json:"code"`
	OrderID      uint                  `json:"orderId" binding:"required"`
	CustomerID   uint                  `json:"customerId" binding:"required"`
	OutboundDate *time.Time            `json:"outboundDate"`
	Remark       string                `json:"remark"`
	Lines        []DocumentLineRequest `json:"lines" binding:"required"`
}
