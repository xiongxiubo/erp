import { ResourcePage } from '../../master/ResourcePage'

export function PurchaseInboundsPage() {
  return (
    <ResourcePage
      title="采购入库"
      description="确认后增加库存并生成应付。"
      path="/purchase/inbounds"
      fields={[
        { name: 'code', label: '入库单号' },
        { name: 'orderId', label: '采购订单 ID', kind: 'number' },
        { name: 'supplierId', label: '供应商 ID', kind: 'number' },
        { name: 'status', label: '状态' },
        { name: 'totalAmount', label: '入库金额', kind: 'number' },
        { name: 'remark', label: '备注' },
      ]}
    />
  )
}
