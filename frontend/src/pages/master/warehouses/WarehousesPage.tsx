import { ResourcePage } from '../ResourcePage'

export function WarehousesPage() {
  return (
    <ResourcePage
      title="仓库资料"
      description="多仓库存和业务单据的仓库基础。"
      path="/master/warehouses"
      fields={[
        { name: 'code', label: '仓库编码' },
        { name: 'name', label: '仓库名称', required: true },
        { name: 'address', label: '地址' },
        { name: 'managerId', label: '负责人 ID', kind: 'number' },
        { name: 'status', label: '状态' },
        { name: 'remark', label: '备注' },
      ]}
    />
  )
}
