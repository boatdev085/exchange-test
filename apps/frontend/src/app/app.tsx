import styled from 'styled-components'
import 'antd/dist/antd.css'
import ExchangePage from './containers/exchange'
import Header from './components/Header'
import AuthProvider from './providers/profileAndAssets'

const Container = styled.div`
  max-width: 1248px;
  margin: 0 auto;
`

export function App() {
  return (
    <AuthProvider>
      <Header />
      <Container>
        <ExchangePage />
      </Container>
    </AuthProvider>
  )
}

export default App
