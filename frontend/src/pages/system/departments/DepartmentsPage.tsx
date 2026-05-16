import { ResourcePage } from '../../master/ResourcePage'

export function DepartmentsPage() {
  return (
    <ResourcePage
      title="部门管理"
      description="搭建组织结构和后续数据权限基础。"
      path="/system/departments"
      fields={[
        { name: 'code', label: '部门编码', required: true },
        { name: 'name', label: '部门名称', required: true },
        { name: 'parentId', label: '上级部门 ID', kind: 'number' },
        { name: 'sort', label: '排序', kind: 'number' },
        { name: 'status', label: '状态' },
      ]}
    />
  )
}
