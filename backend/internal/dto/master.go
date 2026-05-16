package dto

type EntityRequest struct {
	Code             string  `json:"code"`
	Name             string  `json:"name" binding:"required"`
	ContactName      string  `json:"contactName"`
	Phone            string  `json:"phone"`
	Address          string  `json:"address"`
	SettlementMethod string  `json:"settlementMethod"`
	CreditLimit      float64 `json:"creditLimit"`
	Status           string  `json:"status"`
	Remark           string  `json:"remark"`
}

type ProductRequest struct {
	SKU             string  `json:"sku"`
	Name            string  `json:"name" binding:"required"`
	CategoryID      uint    `json:"categoryId"`
	UnitID          uint    `json:"unitId"`
	Spec            string  `json:"spec"`
	Barcode         string  `json:"barcode"`
	PurchasePrice   float64 `json:"purchasePrice"`
	SalePrice       float64 `json:"salePrice"`
	StockWarningQty float64 `json:"stockWarningQty"`
	Status          string  `json:"status"`
	Remark          string  `json:"remark"`
}

type WarehouseRequest struct {
	Code      string `json:"code"`
	Name      string `json:"name" binding:"required"`
	Address   string `json:"address"`
	ManagerID uint   `json:"managerId"`
	Status    string `json:"status"`
	Remark    string `json:"remark"`
}
