package service

import (
	"time"

	"gorm.io/gorm"
)

type DashboardService struct {
	db *gorm.DB
}

type DashboardSummary struct {
	TodaySales       float64 `json:"todaySales"`
	MonthSales       float64 `json:"monthSales"`
	MonthPurchase    float64 `json:"monthPurchase"`
	ReceivableAmount float64 `json:"receivableAmount"`
	PayableAmount    float64 `json:"payableAmount"`
	InventoryValue   float64 `json:"inventoryValue"`
	LowStockCount    int64   `json:"lowStockCount"`
}

type RecentDocument struct {
	Type      string    `json:"type"`
	Code      string    `json:"code"`
	Status    string    `json:"status"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewDashboardService(db *gorm.DB) *DashboardService {
	return &DashboardService{db: db}
}

func (s *DashboardService) Summary() (DashboardSummary, error) {
	now := time.Now()
	dayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	summary := DashboardSummary{}
	s.db.Table("sales_outbounds").Where("status = ? and created_at >= ?", "confirmed", dayStart).Select("coalesce(sum(total_amount), 0)").Scan(&summary.TodaySales)
	s.db.Table("sales_outbounds").Where("status = ? and created_at >= ?", "confirmed", monthStart).Select("coalesce(sum(total_amount), 0)").Scan(&summary.MonthSales)
	s.db.Table("purchase_inbounds").Where("status = ? and created_at >= ?", "confirmed", monthStart).Select("coalesce(sum(total_amount), 0)").Scan(&summary.MonthPurchase)
	s.db.Table("fin_receivables").Where("status != ?", "voided").Select("coalesce(sum(unpaid_amount), 0)").Scan(&summary.ReceivableAmount)
	s.db.Table("fin_payables").Where("status != ?", "voided").Select("coalesce(sum(unpaid_amount), 0)").Scan(&summary.PayableAmount)
	s.db.Table("inv_stocks").Select("coalesce(sum(quantity * average_cost), 0)").Scan(&summary.InventoryValue)
	s.db.Table("inv_stocks").Joins("join md_products on md_products.id = inv_stocks.product_id").Where("inv_stocks.quantity <= md_products.stock_warning_qty").Count(&summary.LowStockCount)
	return summary, nil
}

func (s *DashboardService) RecentDocuments() ([]RecentDocument, error) {
	var docs []RecentDocument
	err := s.db.Raw(`
		(select 'purchase_order' as type, code, status, total_amount as amount, created_at from purchase_orders)
		union all
		(select 'purchase_inbound' as type, code, status, total_amount as amount, created_at from purchase_inbounds)
		union all
		(select 'sales_order' as type, code, status, total_amount as amount, created_at from sales_orders)
		union all
		(select 'sales_outbound' as type, code, status, total_amount as amount, created_at from sales_outbounds)
		order by created_at desc
		limit 12
	`).Scan(&docs).Error
	return docs, err
}
