package model

import "time"

type Receivable struct {
	BaseModel
	SourceType   string     `gorm:"size:80;index;not null" json:"sourceType"`
	SourceID     uint       `gorm:"index;not null" json:"sourceId"`
	SourceCode   string     `gorm:"size:80;index" json:"sourceCode"`
	CustomerID   uint       `gorm:"index;not null" json:"customerId"`
	Amount       float64    `json:"amount"`
	PaidAmount   float64    `json:"paidAmount"`
	UnpaidAmount float64    `json:"unpaidAmount"`
	Status       string     `gorm:"size:32;index;default:unpaid" json:"status"`
	DueDate      *time.Time `json:"dueDate"`
}

func (Receivable) TableName() string { return "fin_receivables" }

type Payable struct {
	BaseModel
	SourceType   string     `gorm:"size:80;index;not null" json:"sourceType"`
	SourceID     uint       `gorm:"index;not null" json:"sourceId"`
	SourceCode   string     `gorm:"size:80;index" json:"sourceCode"`
	SupplierID   uint       `gorm:"index;not null" json:"supplierId"`
	Amount       float64    `json:"amount"`
	PaidAmount   float64    `json:"paidAmount"`
	UnpaidAmount float64    `json:"unpaidAmount"`
	Status       string     `gorm:"size:32;index;default:unpaid" json:"status"`
	DueDate      *time.Time `json:"dueDate"`
}

func (Payable) TableName() string { return "fin_payables" }
