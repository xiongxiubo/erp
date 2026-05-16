import { ResourcePage } from '../../master/ResourcePage'

export function UsersPage() {
  return (
    <ResourcePage
      title="用户管理"
      description="维护登录账号、角色和部门归属。"
      path="/system/users"
      fields={[
        { name: 'username', label: '账号', required: true },
        { name: 'name', label: '姓名', required: true },
        { name: 'phone', label: '电话' },
        { name: 'email', label: '邮箱' },
        { name: 'status', label: '状态' },
      ]}
    />
  )
}
