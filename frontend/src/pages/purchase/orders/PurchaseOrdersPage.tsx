import { ResourcePage } from '../../master/ResourcePage'

export function PurchaseOrdersPage() {
  return (
    <ResourcePage
      title="采购订单"
      description="采购计划审批后的供应商采购单据。"
      path="/purchase/orders"
      fields={[
        { name: 'code', label: '采购单号' },
        { name: 'supplierId', label: '供应商 ID', kind: 'number' },
        { name: 'status', label: '状态' },
        { name: 'totalAmount', label: '订单金额', kind: 'number' },
        { name: 'inboundAmount', label: '入库金额', kind: 'number' },
        { name: 'remark', label: '备注' },
      ]}
    />
  )
}
