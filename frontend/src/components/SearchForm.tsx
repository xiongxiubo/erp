import { Button, Form, Input, Select, Space } from 'antd'

export function SearchForm({ onSearch, onReset }: { onSearch: (values: { keyword?: string; status?: string }) => void; onReset: () => void }) {
  const [form] = Form.useForm()
  return (
    <Form form={form} layout="inline" className="erp-search" onFinish={onSearch}>
      <Form.Item name="keyword">
        <Input allowClear placeholder="编码 / 名称 / 单号" />
      </Form.Item>
      <Form.Item name="status">
        <Select
          allowClear
          placeholder="状态"
          style={{ width: 128 }}
          options={[{ value: 'enabled', label: '启用' }, { value: 'disabled', label: '停用' }, { value: 'draft', label: '草稿' }, { value: 'approved', label: '已审批' }, { value: 'confirmed', label: '已确认' }]}
        />
      </Form.Item>
      <Space>
        <Button type="primary" htmlType="submit">查询</Button>
        <Button onClick={() => { form.resetFields(); onReset() }}>重置</Button>
      </Space>
    </Form>
  )
}
