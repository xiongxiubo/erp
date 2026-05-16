package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy uint           `json:"createdBy"`
	UpdatedBy uint           `json:"updatedBy"`
	Remark    string         `gorm:"size:500" json:"remark"`
}

const (
	StatusEnabled  = "enabled"
	StatusDisabled = "disabled"

	DocumentDraft             = "draft"
	DocumentApproved          = "approved"
	DocumentPartiallyInbound  = "partially_inbound"
	DocumentPartiallyOutbound = "partially_outbound"
	DocumentCompleted         = "completed"
	DocumentClosed            = "closed"
	DocumentVoided            = "voided"
	DocumentConfirmed         = "confirmed"
	FinanceUnpaid             = "unpaid"
	FinancePartial            = "partial"
	FinancePaid               = "paid"
	FinanceVoided             = "voided"
	LedgerDirectionIn         = "in"
	LedgerDirectionOut        = "out"
	LedgerBizPurchaseInbound  = "purchase_inbound"
	LedgerBizSalesOutbound    = "sales_outbound"
)
