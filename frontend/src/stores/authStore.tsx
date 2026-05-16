import { createContext, useContext, useMemo, useState, type ReactNode } from 'react'
import { TOKEN_KEY } from '../utils/constants'
import type { MenuNode, UserProfile } from '../types/common'

type AuthState = {
  token: string
  user?: UserProfile
  menus: MenuNode[]
  permissions: string[]
  setSession: (next: { token: string; user: UserProfile; menus?: MenuNode[]; permissions?: string[] }) => void
  setAccess: (next: { menus: MenuNode[]; permissions: string[] }) => void
  clearSession: () => void
  hasPermission: (code?: string) => boolean
}

const AuthContext = createContext<AuthState | null>(null)

export function AuthProvider({ children }: { children: ReactNode }) {
  const [token, setToken] = useState(() => localStorage.getItem(TOKEN_KEY) ?? '')
  const [user, setUser] = useState<UserProfile>()
  const [menus, setMenus] = useState<MenuNode[]>([])
  const [permissions, setPermissions] = useState<string[]>([])

  const value = useMemo<AuthState>(() => ({
    token,
    user,
    menus,
    permissions,
    setSession(next) {
      localStorage.setItem(TOKEN_KEY, next.token)
      setToken(next.token)
      setUser(next.user)
      setMenus(next.menus ?? [])
      setPermissions(next.permissions ?? [])
    },
    setAccess(next) {
      setMenus(next.menus)
      setPermissions(next.permissions)
    },
    clearSession() {
      localStorage.removeItem(TOKEN_KEY)
      setToken('')
      setUser(undefined)
      setMenus([])
      setPermissions([])
    },
    hasPermission(code) {
      return !code || permissions.includes(code)
    },
  }), [menus, permissions, token, user])

  return <AuthContext value={value}>{children}</AuthContext>
}

export function useAuth() {
  const ctx = useContext(AuthContext)
  if (!ctx) {
    throw new Error('useAuth must be used within AuthProvider')
  }
  return ctx
}
