package database

import (
	"erp/internal/model"
	"erp/internal/utils"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{}, &model.Role{}, &model.Permission{}, &model.Menu{}, &model.Department{}, &model.OperationLog{}, &model.LoginLog{},
		&model.Customer{}, &model.Supplier{}, &model.ProductCategory{}, &model.Unit{}, &model.Product{}, &model.Warehouse{},
		&model.PurchaseOrder{}, &model.PurchaseOrderLine{}, &model.PurchaseInbound{}, &model.PurchaseInboundLine{},
		&model.SalesOrder{}, &model.SalesOrderLine{}, &model.SalesOutbound{}, &model.SalesOutboundLine{},
		&model.Stock{}, &model.StockLedger{}, &model.Receivable{}, &model.Payable{},
	)
}

func Seed(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		adminRole := model.Role{Code: "admin", Name: "系统管理员", Status: model.StatusEnabled}
		if err := upsertBy(tx, &adminRole, "code = ?", adminRole.Code); err != nil {
			return err
		}
		dept := model.Department{Code: "HQ", Name: "总部", Status: model.StatusEnabled}
		if err := upsertBy(tx, &dept, "code = ?", dept.Code); err != nil {
			return err
		}
		hash, err := utils.HashPassword("admin123")
		if err != nil {
			return err
		}
		admin := model.User{Username: "admin", PasswordHash: hash, Name: "系统管理员", DepartmentID: dept.ID, Status: model.StatusEnabled}
		if err := upsertBy(tx, &admin, "username = ?", admin.Username); err != nil {
			return err
		}
		if err := tx.Model(&admin).Association("Roles").Replace(&adminRole); err != nil {
			return err
		}
		permissions := seedPermissions()
		for i := range permissions {
			if err := upsertBy(tx, &permissions[i], "code = ?", permissions[i].Code); err != nil {
				return err
			}
		}
		menus := seedMenus()
		for i := range menus {
			if err := upsertBy(tx, &menus[i], "path = ?", menus[i].Path); err != nil {
				return err
			}
		}
		if err := tx.Model(&adminRole).Association("Permissions").Replace(&permissions); err != nil {
			return err
		}
		if err := tx.Model(&adminRole).Association("Menus").Replace(&menus); err != nil {
			return err
		}
		units := []model.Unit{{Code: "PCS", Name: "件", Status: model.StatusEnabled}, {Code: "KG", Name: "千克", Status: model.StatusEnabled}, {Code: "BOX", Name: "箱", Status: model.StatusEnabled}}
		for i := range units {
			if err := upsertBy(tx, &units[i], "code = ?", units[i].Code); err != nil {
				return err
			}
		}
		warehouse := model.Warehouse{Code: "MAIN", Name: "主仓库", Status: model.StatusEnabled}
		return upsertBy(tx, &warehouse, "code = ?", warehouse.Code)
	})
}

func upsertBy[T any](tx *gorm.DB, value *T, query string, args ...any) error {
	result := tx.Where(query, args...).Assign(value).FirstOrCreate(value)
	return result.Error
}

func seedPermissions() []model.Permission {
	codes := []string{
		"system:user:list", "system:user:create", "system:role:list", "system:menu:list", "system:department:list",
		"master:customer:list", "master:supplier:list", "master:product:list", "master:warehouse:list", "master:unit:list", "master:category:list",
		"purchase:order:list", "purchase:order:create", "purchase:order:approve", "purchase:inbound:list", "purchase:inbound:confirm",
		"sales:order:list", "sales:order:create", "sales:order:approve", "sales:outbound:list", "sales:outbound:confirm",
		"inventory:stock:list", "inventory:ledger:list", "finance:receivable:list", "finance:payable:list", "dashboard:view",
	}
	items := make([]model.Permission, 0, len(codes))
	for _, code := range codes {
		items = append(items, model.Permission{Code: code, Name: code, ResourceType: "api", Status: model.StatusEnabled})
	}
	return items
}

func seedMenus() []model.Menu {
	return []model.Menu{
		{Name: "看板", Path: "/dashboard", Icon: "DashboardOutlined", Sort: 1, PermissionCode: "dashboard:view", Visible: true, Status: model.StatusEnabled},
		{Name: "系统管理", Path: "/system", Icon: "SafetyCertificateOutlined", Sort: 10, Visible: true, Status: model.StatusEnabled},
		{Name: "用户管理", ParentID: 2, Path: "/system/users", Sort: 11, PermissionCode: "system:user:list", Visible: true, Status: model.StatusEnabled},
		{Name: "角色管理", ParentID: 2, Path: "/system/roles", Sort: 12, PermissionCode: "system:role:list", Visible: true, Status: model.StatusEnabled},
		{Name: "菜单管理", ParentID: 2, Path: "/system/menus", Sort: 13, PermissionCode: "system:menu:list", Visible: true, Status: model.StatusEnabled},
		{Name: "部门管理", ParentID: 2, Path: "/system/departments", Sort: 14, PermissionCode: "system:department:list", Visible: true, Status: model.StatusEnabled},
		{Name: "基础资料", Path: "/master", Icon: "DatabaseOutlined", Sort: 20, Visible: true, Status: model.StatusEnabled},
		{Name: "客户", ParentID: 7, Path: "/master/customers", Sort: 21, PermissionCode: "master:customer:list", Visible: true, Status: model.StatusEnabled},
		{Name: "供应商", ParentID: 7, Path: "/master/suppliers", Sort: 22, PermissionCode: "master:supplier:list", Visible: true, Status: model.StatusEnabled},
		{Name: "商品", ParentID: 7, Path: "/master/products", Sort: 23, PermissionCode: "master:product:list", Visible: true, Status: model.StatusEnabled},
		{Name: "仓库", ParentID: 7, Path: "/master/warehouses", Sort: 24, PermissionCode: "master:warehouse:list", Visible: true, Status: model.StatusEnabled},
		{Name: "采购", Path: "/purchase", Icon: "ShoppingCartOutlined", Sort: 30, Visible: true, Status: model.StatusEnabled},
		{Name: "采购订单", ParentID: 12, Path: "/purchase/orders", Sort: 31, PermissionCode: "purchase:order:list", Visible: true, Status: model.StatusEnabled},
		{Name: "采购入库", ParentID: 12, Path: "/purchase/inbounds", Sort: 32, PermissionCode: "purchase:inbound:list", Visible: true, Status: model.StatusEnabled},
		{Name: "销售", Path: "/sales", Icon: "ShopOutlined", Sort: 40, Visible: true, Status: model.StatusEnabled},
		{Name: "销售订单", ParentID: 15, Path: "/sales/orders", Sort: 41, PermissionCode: "sales:order:list", Visible: true, Status: model.StatusEnabled},
		{Name: "销售出库", ParentID: 15, Path: "/sales/outbounds", Sort: 42, PermissionCode: "sales:outbound:list", Visible: true, Status: model.StatusEnabled},
		{Name: "库存", Path: "/inventory", Icon: "InboxOutlined", Sort: 50, Visible: true, Status: model.StatusEnabled},
		{Name: "库存查询", ParentID: 18, Path: "/inventory/stocks", Sort: 51, PermissionCode: "inventory:stock:list", Visible: true, Status: model.StatusEnabled},
		{Name: "库存流水", ParentID: 18, Path: "/inventory/ledgers", Sort: 52, PermissionCode: "inventory:ledger:list", Visible: true, Status: model.StatusEnabled},
		{Name: "财务", Path: "/finance", Icon: "AccountBookOutlined", Sort: 60, Visible: true, Status: model.StatusEnabled},
		{Name: "应收", ParentID: 21, Path: "/finance/receivables", Sort: 61, PermissionCode: "finance:receivable:list", Visible: true, Status: model.StatusEnabled},
		{Name: "应付", ParentID: 21, Path: "/finance/payables", Sort: 62, PermissionCode: "finance:payable:list", Visible: true, Status: model.StatusEnabled},
	}
}
