import type { ReactNode } from 'react'
import { useAuth } from '../stores/authStore'

export function PermissionButton({ code, children }: { code?: string; children: ReactNode }) {
  const { hasPermission } = useAuth()
  if (!hasPermission(code)) {
    return null
  }
  return children
}
