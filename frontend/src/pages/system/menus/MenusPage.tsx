import { ResourcePage } from '../../master/ResourcePage'

export function MenusPage() {
  return (
    <ResourcePage
      title="菜单管理"
      description="配置后台导航和权限码。"
      path="/system/menus"
      fields={[
        { name: 'name', label: '菜单名称', required: true },
        { name: 'path', label: '路由路径' },
        { name: 'icon', label: '图标' },
        { name: 'permissionCode', label: '权限码' },
        { name: 'sort', label: '排序', kind: 'number' },
        { name: 'status', label: '状态' },
      ]}
    />
  )
}
