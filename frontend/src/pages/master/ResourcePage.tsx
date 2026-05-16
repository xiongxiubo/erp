import { Button, Card, Form, Input, InputNumber, Modal, Space, message } from 'antd'
import { useEffect, useMemo, useState } from 'react'
import { DataTable } from '../../components/DataTable'
import { PageHeader } from '../../components/PageHeader'
import { SearchForm } from '../../components/SearchForm'
import { StatusTag } from '../../components/StatusTag'
import { createResource, deleteResource, listResource, updateResource } from '../../services/resource'
import type { BaseEntity } from '../../types/common'

type FieldKind = 'text' | 'number'

type Field = {
  name: string
  label: string
  kind?: FieldKind
  required?: boolean
}

export function ResourcePage({ title, description, path, fields }: { title: string; description: string; path: string; fields: Field[] }) {
  const [data, setData] = useState<BaseEntity[]>([])
  const [total, setTotal] = useState(0)
  const [loading, setLoading] = useState(false)
  const [modalOpen, setModalOpen] = useState(false)
  const [editing, setEditing] = useState<BaseEntity>()
  const [query, setQuery] = useState({ page: 1, pageSize: 20, keyword: '', status: '' })
  const [form] = Form.useForm()

  const load = useMemo(() => async () => {
    setLoading(true)
    try {
      const result = await listResource(path, query)
      setData(result.items)
      setTotal(result.total)
    } finally {
      setLoading(false)
    }
  }, [path, query])

  useEffect(() => {
    load()
  }, [load])

  function openCreate() {
    setEditing(undefined)
    form.resetFields()
    form.setFieldsValue({ status: 'enabled' })
    setModalOpen(true)
  }

  function openEdit(row: BaseEntity) {
    setEditing(row)
    form.setFieldsValue(row)
    setModalOpen(true)
  }

  async function submit() {
    const values = await form.validateFields()
    if (editing?.id) {
      await updateResource(path, editing.id, values)
      message.success('更新成功')
    } else {
      await createResource(path, values)
      message.success('创建成功')
    }
    setModalOpen(false)
    load()
  }

  return (
    <div className="erp-page">
      <PageHeader title={title} description={description} extra={<Button type="primary" onClick={openCreate}>新增</Button>} />
      <Card className="erp-card">
        <SearchForm
          onSearch={(values) => setQuery((prev) => ({ ...prev, page: 1, keyword: values.keyword ?? '', status: values.status ?? '' }))}
          onReset={() => setQuery((prev) => ({ ...prev, page: 1, keyword: '', status: '' }))}
        />
        <DataTable
          columns={[
            { title: 'ID', dataIndex: 'id', width: 80 },
            ...fields.slice(0, 4).map((field) => ({ title: field.label, dataIndex: field.name })),
            { title: '状态', dataIndex: 'status', render: (value) => <StatusTag value={value} /> },
            {
              title: '操作',
              width: 160,
              render: (_, row) => (
                <Space>
                  <Button type="link" onClick={() => openEdit(row)}>编辑</Button>
                  <Button type="link" danger onClick={async () => { await deleteResource(path, row.id); message.success('删除成功'); load() }}>删除</Button>
                </Space>
              ),
            },
          ]}
          data={data}
          loading={loading}
          total={total}
          page={query.page}
          pageSize={query.pageSize}
          onChange={(page, pageSize) => setQuery((prev) => ({ ...prev, page, pageSize }))}
        />
      </Card>
      <Modal open={modalOpen} title={editing ? `编辑${title}` : `新增${title}`} onOk={submit} onCancel={() => setModalOpen(false)} destroyOnHidden>
        <Form form={form} layout="vertical">
          {fields.map((field) => (
            <Form.Item key={field.name} name={field.name} label={field.label} rules={field.required ? [{ required: true, message: `请输入${field.label}` }] : undefined}>
              {field.kind === 'number' ? <InputNumber style={{ width: '100%' }} /> : <Input />}
            </Form.Item>
          ))}
        </Form>
      </Modal>
    </div>
  )
}
