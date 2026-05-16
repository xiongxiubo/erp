import { Table } from 'antd'
import type { ColumnsType, TablePaginationConfig } from 'antd/es/table'

export function DataTable<T extends object>({
  columns,
  data,
  loading,
  total,
  page,
  pageSize,
  onChange,
}: {
  columns: ColumnsType<T>
  data: T[]
  loading: boolean
  total: number
  page: number
  pageSize: number
  onChange: (page: number, pageSize: number) => void
}) {
  const pagination: TablePaginationConfig = {
    current: page,
    pageSize,
    total,
    showSizeChanger: true,
    showTotal: (value) => `共 ${value} 条`,
    onChange,
  }
  return <Table rowKey="id" columns={columns} dataSource={data} loading={loading} pagination={pagination} />
}
