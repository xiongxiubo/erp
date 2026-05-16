import { ResourcePage } from '../ResourcePage'

export function SuppliersPage() {
  return (
    <ResourcePage
      title="供应商资料"
      description="采购订单、入库和应付的供应商主数据。"
      path="/master/suppliers"
      fields={[
        { name: 'code', label: '供应商编码' },
        { name: 'name', label: '供应商名称', required: true },
        { name: 'contactName', label: '联系人' },
        { name: 'phone', label: '电话' },
        { name: 'address', label: '地址' },
        { name: 'status', label: '状态' },
        { name: 'remark', label: '备注' },
      ]}
    />
  )
}
