import { ConfigProvider, theme } from 'antd'
import zhCN from 'antd/locale/zh_CN'
import type { ReactNode } from 'react'
import { AuthProvider } from '../stores/authStore'

export function Providers({ children }: { children: ReactNode }) {
  return (
    <ConfigProvider
      locale={zhCN}
      theme={{
        algorithm: theme.defaultAlgorithm,
        token: {
          colorPrimary: '#155eef',
          colorInfo: '#0f766e',
          borderRadius: 12,
          fontFamily: 'Aptos, "Microsoft YaHei", "Noto Sans SC", sans-serif',
        },
        components: {
          Layout: { bodyBg: 'transparent', headerBg: 'rgba(255,255,255,0.8)', siderBg: '#0d1b2a' },
          Menu: { darkItemBg: '#0d1b2a', darkSubMenuItemBg: '#0b1724', darkItemSelectedBg: '#155eef' },
          Card: { borderRadiusLG: 18 },
        },
      }}
    >
      <AuthProvider>{children}</AuthProvider>
    </ConfigProvider>
  )
}
