import { ResourcePage } from '../../master/ResourcePage'

export function InventoryStocksPage() {
  return (
    <ResourcePage
      title="库存查询"
      description="按商品和仓库查看实时现存量。"
      path="/inventory/stocks"
      fields={[
        { name: 'warehouseId', label: '仓库 ID', kind: 'number' },
        { name: 'productId', label: '商品 ID', kind: 'number' },
        { name: 'quantity', label: '现存量', kind: 'number' },
        { name: 'lockedQuantity', label: '锁定量', kind: 'number' },
        { name: 'averageCost', label: '平均成本', kind: 'number' },
      ]}
    />
  )
}
