import { ResourcePage } from '../../master/ResourcePage'

export function SalesOutboundsPage() {
  return (
    <ResourcePage
      title="销售出库"
      description="确认前校验库存，确认后扣减库存并生成应收。"
      path="/sales/outbounds"
      fields={[
        { name: 'code', label: '出库单号' },
        { name: 'orderId', label: '销售订单 ID', kind: 'number' },
        { name: 'customerId', label: '客户 ID', kind: 'number' },
        { name: 'status', label: '状态' },
        { name: 'totalAmount', label: '出库金额', kind: 'number' },
        { name: 'remark', label: '备注' },
      ]}
    />
  )
}
