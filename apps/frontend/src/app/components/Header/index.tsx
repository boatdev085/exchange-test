import styled from 'styled-components'
import { UserOutlined } from '@ant-design/icons'
import { Avatar } from 'antd'
import { useProfileAndAssets } from '../../providers/profileAndAssets'
import numeral from 'numeral'

const Header = () => {
  const { userProfile, wallets } = useProfileAndAssets()
  const findTHB = wallets?.find((f) => f.symbol === 'THB')
  return (
    <Container>
      <div>balance: {numeral(findTHB?.amount || 0).format('0,0.00')} thb</div>
      <div>
        {userProfile?.first_name} {userProfile?.last_name}
      </div>

      <Avatar icon={<UserOutlined />} />
    </Container>
  )
}

export default Header

const Container = styled.div`
  position: relative;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  padding: 12px;
  gap: 12px;
`
