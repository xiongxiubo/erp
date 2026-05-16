import { Tag } from 'antd'
import { statusColor, statusText } from '../utils/constants'

export function StatusTag({ value }: { value?: string }) {
  const status = value ?? 'enabled'
  return <Tag color={statusColor[status] ?? 'default'}>{statusText[status] ?? status}</Tag>
}
