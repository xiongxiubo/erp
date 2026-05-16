export type ApiResponse<T> = {
  code: number
  message: string
  data: T
}

export type PageResult<T> = {
  items: T[]
  total: number
  page: number
  pageSize: number
}

export type Status = 'enabled' | 'disabled' | 'draft' | 'approved' | 'confirmed' | 'completed' | 'closed' | 'voided' | 'unpaid' | 'partial' | 'paid'

export type MenuNode = {
  id: number
  parentId: number
  name: string
  path: string
  icon: string
  sort: number
  permissionCode: string
  children?: MenuNode[]
}

export type UserProfile = {
  id: number
  username: string
  name: string
  email: string
  phone: string
  departmentId: number
  status: string
  roles: Array<{ id: number; code: string; name: string }>
}

export type LoginResponse = {
  token: string
  expiresAt: string
  user: UserProfile
}

export type BaseEntity = {
  id: number
  code?: string
  name?: string
  status?: string
  remark?: string
  createdAt?: string
  updatedAt?: string
}

export type DashboardSummary = {
  todaySales: number
  monthSales: number
  monthPurchase: number
  receivableAmount: number
  payableAmount: number
  inventoryValue: number
  lowStockCount: number
}

export type RecentDocument = {
  type: string
  code: string
  status: string
  amount: number
  createdAt: string
}
