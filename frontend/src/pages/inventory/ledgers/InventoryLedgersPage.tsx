import { ResourcePage } from '../../master/ResourcePage'

export function InventoryLedgersPage() {
  return (
    <ResourcePage
      title="库存流水"
      description="追踪每一次库存变化的来源和前后数量。"
      path="/inventory/ledgers"
      fields={[
        { name: 'bizType', label: '业务类型' },
        { name: 'bizCode', label: '业务单号' },
        { name: 'direction', label: '方向' },
        { name: 'warehouseId', label: '仓库 ID', kind: 'number' },
        { name: 'productId', label: '商品 ID', kind: 'number' },
        { name: 'quantity', label: '数量', kind: 'number' },
      ]}
    />
  )
}
