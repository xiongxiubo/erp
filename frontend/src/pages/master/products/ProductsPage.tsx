import { ResourcePage } from '../ResourcePage'

export function ProductsPage() {
  return (
    <ResourcePage
      title="商品物料"
      description="库存、采购、销售共用的商品档案。"
      path="/master/products"
      fields={[
        { name: 'sku', label: 'SKU' },
        { name: 'name', label: '商品名称', required: true },
        { name: 'spec', label: '规格' },
        { name: 'barcode', label: '条码' },
        { name: 'purchasePrice', label: '采购价', kind: 'number' },
        { name: 'salePrice', label: '销售价', kind: 'number' },
        { name: 'stockWarningQty', label: '预警库存', kind: 'number' },
        { name: 'status', label: '状态' },
      ]}
    />
  )
}
