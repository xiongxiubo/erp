package model

type Customer struct {
	BaseModel
	Code             string  `gorm:"size:80;uniqueIndex;not null" json:"code"`
	Name             string  `gorm:"size:160;not null;index" json:"name"`
	ContactName      string  `gorm:"size:80" json:"contactName"`
	Phone            string  `gorm:"size:40" json:"phone"`
	Address          string  `gorm:"size:255" json:"address"`
	SettlementMethod string  `gorm:"size:80" json:"settlementMethod"`
	CreditLimit      float64 `json:"creditLimit"`
	Status           string  `gorm:"size:24;index;default:enabled" json:"status"`
}

func (Customer) TableName() string { return "md_customers" }

type Supplier struct {
	BaseModel
	Code             string  `gorm:"size:80;uniqueIndex;not null" json:"code"`
	Name             string  `gorm:"size:160;not null;index" json:"name"`
	ContactName      string  `gorm:"size:80" json:"contactName"`
	Phone            string  `gorm:"size:40" json:"phone"`
	Address          string  `gorm:"size:255" json:"address"`
	SettlementMethod string  `gorm:"size:80" json:"settlementMethod"`
	CreditLimit      float64 `json:"creditLimit"`
	Status           string  `gorm:"size:24;index;default:enabled" json:"status"`
}

func (Supplier) TableName() string { return "md_suppliers" }

type ProductCategory struct {
	BaseModel
	ParentID uint   `gorm:"index" json:"parentId"`
	Code     string `gorm:"size:80;uniqueIndex;not null" json:"code"`
	Name     string `gorm:"size:120;not null" json:"name"`
	Sort     int    `gorm:"index" json:"sort"`
	Status   string `gorm:"size:24;index;default:enabled" json:"status"`
}

func (ProductCategory) TableName() string { return "md_product_categories" }

type Unit struct {
	BaseModel
	Code   string `gorm:"size:80;uniqueIndex;not null" json:"code"`
	Name   string `gorm:"size:80;not null" json:"name"`
	Status string `gorm:"size:24;index;default:enabled" json:"status"`
}

func (Unit) TableName() string { return "md_units" }

type Product struct {
	BaseModel
	SKU             string  `gorm:"size:80;uniqueIndex;not null" json:"sku"`
	Name            string  `gorm:"size:160;not null;index" json:"name"`
	CategoryID      uint    `gorm:"index" json:"categoryId"`
	UnitID          uint    `gorm:"index" json:"unitId"`
	Spec            string  `gorm:"size:160" json:"spec"`
	Barcode         string  `gorm:"size:120;index" json:"barcode"`
	PurchasePrice   float64 `json:"purchasePrice"`
	SalePrice       float64 `json:"salePrice"`
	StockWarningQty float64 `json:"stockWarningQty"`
	Status          string  `gorm:"size:24;index;default:enabled" json:"status"`
}

func (Product) TableName() string { return "md_products" }

type Warehouse struct {
	BaseModel
	Code      string `gorm:"size:80;uniqueIndex;not null" json:"code"`
	Name      string `gorm:"size:120;not null;index" json:"name"`
	Address   string `gorm:"size:255" json:"address"`
	ManagerID uint   `gorm:"index" json:"managerId"`
	Status    string `gorm:"size:24;index;default:enabled" json:"status"`
}

func (Warehouse) TableName() string { return "md_warehouses" }
