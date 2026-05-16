import { createBrowserRouter, Navigate, RouterProvider } from 'react-router-dom'
import { BasicLayout } from '../layouts/BasicLayout'
import { DashboardPage } from '../pages/dashboard/DashboardPage'
import { PayablesPage } from '../pages/finance/payables/PayablesPage'
import { ReceivablesPage } from '../pages/finance/receivables/ReceivablesPage'
import { InventoryLedgersPage } from '../pages/inventory/ledgers/InventoryLedgersPage'
import { InventoryStocksPage } from '../pages/inventory/stocks/InventoryStocksPage'
import { LoginPage } from '../pages/login/LoginPage'
import { CustomersPage } from '../pages/master/customers/CustomersPage'
import { ProductsPage } from '../pages/master/products/ProductsPage'
import { SuppliersPage } from '../pages/master/suppliers/SuppliersPage'
import { WarehousesPage } from '../pages/master/warehouses/WarehousesPage'
import { PurchaseInboundsPage } from '../pages/purchase/inbounds/PurchaseInboundsPage'
import { PurchaseOrdersPage } from '../pages/purchase/orders/PurchaseOrdersPage'
import { SalesOrdersPage } from '../pages/sales/orders/SalesOrdersPage'
import { SalesOutboundsPage } from '../pages/sales/outbounds/SalesOutboundsPage'
import { DepartmentsPage } from '../pages/system/departments/DepartmentsPage'
import { MenusPage } from '../pages/system/menus/MenusPage'
import { RolesPage } from '../pages/system/roles/RolesPage'
import { UsersPage } from '../pages/system/users/UsersPage'
import { useAuth } from '../stores/authStore'

function Protected() {
  const { token } = useAuth()
  return token ? <BasicLayout /> : <Navigate to="/login" replace />
}

const router = createBrowserRouter([
  { path: '/login', element: <LoginPage /> },
  {
    path: '/',
    element: <Protected />,
    children: [
      { index: true, element: <Navigate to="/dashboard" replace /> },
      { path: 'dashboard', element: <DashboardPage /> },
      { path: 'system/users', element: <UsersPage /> },
      { path: 'system/roles', element: <RolesPage /> },
      { path: 'system/menus', element: <MenusPage /> },
      { path: 'system/departments', element: <DepartmentsPage /> },
      { path: 'master/customers', element: <CustomersPage /> },
      { path: 'master/suppliers', element: <SuppliersPage /> },
      { path: 'master/products', element: <ProductsPage /> },
      { path: 'master/warehouses', element: <WarehousesPage /> },
      { path: 'purchase/orders', element: <PurchaseOrdersPage /> },
      { path: 'purchase/inbounds', element: <PurchaseInboundsPage /> },
      { path: 'sales/orders', element: <SalesOrdersPage /> },
      { path: 'sales/outbounds', element: <SalesOutboundsPage /> },
      { path: 'inventory/stocks', element: <InventoryStocksPage /> },
      { path: 'inventory/ledgers', element: <InventoryLedgersPage /> },
      { path: 'finance/receivables', element: <ReceivablesPage /> },
      { path: 'finance/payables', element: <PayablesPage /> },
    ],
  },
])

export function AppRouter() {
  return <RouterProvider router={router} />
}
