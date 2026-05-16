import { ResourcePage } from '../../master/ResourcePage'

export function SalesOrdersPage() {
  return (
    <ResourcePage
      title="销售订单"
      description="客户销售履约的核心单据。"
      path="/sales/orders"
      fields={[
        { name: 'code', label: '销售单号' },
        { name: 'customerId', label: '客户 ID', kind: 'number' },
        { name: 'status', label: '状态' },
        { name: 'totalAmount', label: '订单金额', kind: 'number' },
        { name: 'outboundAmount', label: '出库金额', kind: 'number' },
        { name: 'remark', label: '备注' },
      ]}
    />
  )
}
