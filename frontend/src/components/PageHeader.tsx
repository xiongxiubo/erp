import { Button, Flex, Typography } from 'antd'
import type { ReactNode } from 'react'

export function PageHeader({ title, description, extra }: { title: string; description?: string; extra?: ReactNode }) {
  return (
    <Flex justify="space-between" align="center" className="erp-page-header">
      <div>
        <Typography.Title level={2}>{title}</Typography.Title>
        {description ? <Typography.Text type="secondary">{description}</Typography.Text> : null}
      </div>
      {extra ? <div>{extra}</div> : <Button type="primary">新增</Button>}
    </Flex>
  )
}
