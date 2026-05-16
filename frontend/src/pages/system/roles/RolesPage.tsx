import { ResourcePage } from '../../master/ResourcePage'

export function RolesPage() {
  return (
    <ResourcePage
      title="角色管理"
      description="组织菜单、按钮和 API 权限的授权边界。"
      path="/system/roles"
      fields={[
        { name: 'code', label: '角色编码', required: true },
        { name: 'name', label: '角色名称', required: true },
        { name: 'description', label: '描述' },
        { name: 'status', label: '状态' },
      ]}
    />
  )
}
