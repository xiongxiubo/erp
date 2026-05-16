import { ResourcePage } from '../ResourcePage'

export function CustomersPage() {
  return (
    <ResourcePage
      title="客户资料"
      description="销售订单、应收和项目的客户主数据。"
      path="/master/customers"
      fields={[
        { name: 'code', label: '客户编码' },
        { name: 'name', label: '客户名称', required: true },
        { name: 'contactName', label: '联系人' },
        { name: 'phone', label: '电话' },
        { name: 'address', label: '地址' },
        { name: 'status', label: '状态' },
        { name: 'remark', label: '备注' },
      ]}
    />
  )
}
