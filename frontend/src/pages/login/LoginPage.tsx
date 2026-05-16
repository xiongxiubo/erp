import { Button, Card, Form, Input, Typography, message } from 'antd'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { getMenus, getPermissions, login } from '../../services/auth'
import { useAuth } from '../../stores/authStore'

export function LoginPage() {
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()
  const { setSession, setAccess } = useAuth()

  async function onFinish(values: { username: string; password: string }) {
    setLoading(true)
    try {
      const data = await login(values)
      setSession({ token: data.token, user: data.user })
      const [menus, permissions] = await Promise.all([getMenus(), getPermissions()])
      setAccess({ menus, permissions })
      message.success('登录成功')
      navigate('/dashboard')
    } finally {
      setLoading(false)
    }
  }

  return (
    <main className="login-page">
      <section className="login-hero">
        <div className="login-orbit" />
        <Typography.Text className="erp-kicker">DINGHENG ERP / PHASE ONE</Typography.Text>
        <Typography.Title>把采购、销售、库存和财务压进一条清晰的经营链路。</Typography.Title>
        <p>从单据到流水，从库存到应收应付，第一阶段先跑通企业最核心的进销存闭环。</p>
        <div className="login-metrics">
          <span>RBAC 权限</span>
          <span>库存事务</span>
          <span>应收应付</span>
        </div>
      </section>
      <Card className="login-card">
        <Typography.Title level={3}>登录控制台</Typography.Title>
        <Typography.Text type="secondary">默认账号：admin / admin123</Typography.Text>
        <Form layout="vertical" onFinish={onFinish} initialValues={{ username: 'admin', password: 'admin123' }}>
          <Form.Item label="账号" name="username" rules={[{ required: true, message: '请输入账号' }]}>
            <Input size="large" placeholder="admin" />
          </Form.Item>
          <Form.Item label="密码" name="password" rules={[{ required: true, message: '请输入密码' }]}>
            <Input.Password size="large" placeholder="admin123" />
          </Form.Item>
          <Button size="large" type="primary" block htmlType="submit" loading={loading}>进入 ERP</Button>
        </Form>
      </Card>
    </main>
  )
}
