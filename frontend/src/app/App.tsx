import { AppRouter } from './router'
import { Providers } from './providers'
import '../styles/global.scss'

export default function App() {
  return (
    <Providers>
      <AppRouter />
    </Providers>
  )
}
