import { ResourcePage } from '../../master/ResourcePage'

export function ReceivablesPage() {
  return (
    <ResourcePage
      title="应收管理"
      description="销售出库确认后自动生成的客户应收。"
      path="/finance/receivables"
      fields={[
        { name: 'sourceType', label: '来源类型' },
        { name: 'sourceCode', label: '来源单号' },
        { name: 'customerId', label: '客户 ID', kind: 'number' },
        { name: 'amount', label: '应收金额', kind: 'number' },
        { name: 'unpaidAmount', label: '未收金额', kind: 'number' },
        { name: 'status', label: '状态' },
      ]}
    />
  )
}
