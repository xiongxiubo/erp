import { ResourcePage } from '../../master/ResourcePage'

export function PayablesPage() {
  return (
    <ResourcePage
      title="应付管理"
      description="采购入库确认后自动生成的供应商应付。"
      path="/finance/payables"
      fields={[
        { name: 'sourceType', label: '来源类型' },
        { name: 'sourceCode', label: '来源单号' },
        { name: 'supplierId', label: '供应商 ID', kind: 'number' },
        { name: 'amount', label: '应付金额', kind: 'number' },
        { name: 'unpaidAmount', label: '未付金额', kind: 'number' },
        { name: 'status', label: '状态' },
      ]}
    />
  )
}
