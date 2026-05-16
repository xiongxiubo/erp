import { ResourcePage } from './master/ResourcePage'

export function SimplePage({ title, description, path, variant = 'entity' }: { title: string; description: string; path: string; variant?: 'entity' | 'product' | 'warehouse' | 'document' }) {
  return <ResourcePage title={title} description={description} path={path} fields={fields[variant]} />
}

const fields = {
  entity: [
    { name: 'code', label: '编码' },
    { name: 'name', label: '名称', required: true },
    { name: 'contactName', label: '联系人' },
    { name: 'phone', label: '电话' },
    { name: 'address', label: '地址' },
    { name: 'status', label: '状态' },
    { name: 'remark', label: '备注' },
  ],
  product: [
    { name: 'sku', label: 'SKU' },
    { name: 'name', label: '商品名称', required: true },
    { name: 'spec', label: '规格' },
    { name: 'barcode', label: '条码' },
    { name: 'purchasePrice', label: '采购价', kind: 'number' as const },
    { name: 'salePrice', label: '销售价', kind: 'number' as const },
    { name: 'stockWarningQty', label: '预警库存', kind: 'number' as const },
    { name: 'status', label: '状态' },
  ],
  warehouse: [
    { name: 'code', label: '仓库编码' },
    { name: 'name', label: '仓库名称', required: true },
    { name: 'address', label: '地址' },
    { name: 'status', label: '状态' },
    { name: 'remark', label: '备注' },
  ],
  document: [
    { name: 'code', label: '单号' },
    { name: 'status', label: '状态' },
    { name: 'totalAmount', label: '金额', kind: 'number' as const },
    { name: 'remark', label: '备注' },
  ],
}
