package router

import (
	"erp/internal/config"
	"erp/internal/controller"
	"erp/internal/middleware"
	"erp/internal/model"
	"erp/internal/response"
	"erp/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(db *gorm.DB, cfg config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.CORS())

	stockSvc := service.NewStockService(db)
	authSvc := service.NewAuthService(db, cfg.JWT.Secret, cfg.JWT.ExpiresIn)
	purchaseSvc := service.NewPurchaseService(db, stockSvc)
	salesSvc := service.NewSalesService(db, stockSvc)
	dashboardSvc := service.NewDashboardService(db)

	authCtl := controller.NewAuthController(authSvc)
	purchaseCtl := controller.NewPurchaseController(purchaseSvc)
	salesCtl := controller.NewSalesController(salesSvc)
	dashboardCtl := controller.NewDashboardController(dashboardSvc)

	api := r.Group("/api/v1")
	api.GET("/health", func(c *gin.Context) { response.OK(c, gin.H{"status": "ok"}) })
	api.POST("/auth/login", authCtl.Login)

	protected := api.Group("")
	protected.Use(middleware.Auth(cfg.JWT.Secret))
	protected.POST("/auth/logout", func(c *gin.Context) { response.OK(c, gin.H{"ok": true}) })
	protected.GET("/auth/profile", authCtl.Profile)
	protected.GET("/auth/menus", authCtl.Menus)
	protected.GET("/auth/permissions", authCtl.Permissions)

	registerCRUD(protected.Group("/system/users"), controller.NewCRUDController(service.NewCRUDService[model.User](db, []string{"username", "name", "phone"}, "")))
	registerCRUD(protected.Group("/system/roles"), controller.NewCRUDController(service.NewCRUDService[model.Role](db, []string{"code", "name"}, "")))
	registerCRUD(protected.Group("/system/menus"), controller.NewCRUDController(service.NewCRUDService[model.Menu](db, []string{"name", "path"}, "")))
	registerCRUD(protected.Group("/system/departments"), controller.NewCRUDController(service.NewCRUDService[model.Department](db, []string{"code", "name"}, "")))
	protected.GET("/system/operation-logs", controller.NewCRUDController(service.NewCRUDService[model.OperationLog](db, []string{"module", "action", "biz_type"}, "")).List)

	registerCRUD(protected.Group("/master/customers"), controller.NewCRUDController(service.NewCRUDService[model.Customer](db, []string{"code", "name", "phone"}, "")))
	registerCRUD(protected.Group("/master/suppliers"), controller.NewCRUDController(service.NewCRUDService[model.Supplier](db, []string{"code", "name", "phone"}, "")))
	registerCRUD(protected.Group("/master/products"), controller.NewCRUDController(service.NewCRUDService[model.Product](db, []string{"sku", "name", "barcode"}, "")))
	registerCRUD(protected.Group("/master/product-categories"), controller.NewCRUDController(service.NewCRUDService[model.ProductCategory](db, []string{"code", "name"}, "")))
	registerCRUD(protected.Group("/master/warehouses"), controller.NewCRUDController(service.NewCRUDService[model.Warehouse](db, []string{"code", "name"}, "")))
	registerCRUD(protected.Group("/master/units"), controller.NewCRUDController(service.NewCRUDService[model.Unit](db, []string{"code", "name"}, "")))

	registerReadCRUD(protected.Group("/purchase/orders"), controller.NewCRUDController(service.NewCRUDService[model.PurchaseOrder](db.Preload("Lines"), []string{"code"}, "")))
	protected.POST("/purchase/orders", purchaseCtl.CreateOrder)
	protected.PUT("/purchase/orders/:id", purchaseCtl.UpdateOrder)
	protected.POST("/purchase/orders/:id/approve", purchaseCtl.ApproveOrder)
	protected.POST("/purchase/orders/:id/close", purchaseCtl.CloseOrder)
	protected.POST("/purchase/orders/:id/void", purchaseCtl.VoidOrder)
	registerReadCRUD(protected.Group("/purchase/inbounds"), controller.NewCRUDController(service.NewCRUDService[model.PurchaseInbound](db.Preload("Lines"), []string{"code"}, "")))
	protected.POST("/purchase/inbounds", purchaseCtl.CreateInbound)
	protected.POST("/purchase/inbounds/:id/confirm", purchaseCtl.ConfirmInbound)
	protected.POST("/purchase/inbounds/:id/void", purchaseCtl.VoidInbound)

	registerReadCRUD(protected.Group("/sales/orders"), controller.NewCRUDController(service.NewCRUDService[model.SalesOrder](db.Preload("Lines"), []string{"code"}, "")))
	protected.POST("/sales/orders", salesCtl.CreateOrder)
	protected.POST("/sales/orders/:id/approve", salesCtl.ApproveOrder)
	protected.POST("/sales/orders/:id/close", salesCtl.CloseOrder)
	protected.POST("/sales/orders/:id/void", salesCtl.VoidOrder)
	registerReadCRUD(protected.Group("/sales/outbounds"), controller.NewCRUDController(service.NewCRUDService[model.SalesOutbound](db.Preload("Lines"), []string{"code"}, "")))
	protected.POST("/sales/outbounds", salesCtl.CreateOutbound)
	protected.POST("/sales/outbounds/:id/confirm", salesCtl.ConfirmOutbound)
	protected.POST("/sales/outbounds/:id/void", salesCtl.VoidOutbound)

	protected.GET("/inventory/stocks", controller.NewCRUDController(service.NewCRUDService[model.Stock](db, []string{}, "")).List)
	protected.GET("/inventory/ledgers", controller.NewCRUDController(service.NewCRUDService[model.StockLedger](db, []string{"biz_code", "biz_type"}, "")).List)
	protected.GET("/finance/receivables", controller.NewCRUDController(service.NewCRUDService[model.Receivable](db, []string{"source_code"}, "")).List)
	protected.GET("/finance/receivables/:id", controller.NewCRUDController(service.NewCRUDService[model.Receivable](db, []string{"source_code"}, "")).Get)
	protected.GET("/finance/payables", controller.NewCRUDController(service.NewCRUDService[model.Payable](db, []string{"source_code"}, "")).List)
	protected.GET("/finance/payables/:id", controller.NewCRUDController(service.NewCRUDService[model.Payable](db, []string{"source_code"}, "")).Get)
	protected.GET("/dashboard/summary", dashboardCtl.Summary)
	protected.GET("/dashboard/recent-documents", dashboardCtl.RecentDocuments)

	return r
}

func registerCRUD[T any](group *gin.RouterGroup, ctl *controller.CRUDController[T]) {
	group.GET("", ctl.List)
	group.POST("", ctl.Create)
	group.GET("/:id", ctl.Get)
	group.PUT("/:id", ctl.Update)
	group.DELETE("/:id", ctl.Delete)
}

func registerReadCRUD[T any](group *gin.RouterGroup, ctl *controller.CRUDController[T]) {
	group.GET("", ctl.List)
	group.GET("/:id", ctl.Get)
	group.DELETE("/:id", ctl.Delete)
}
