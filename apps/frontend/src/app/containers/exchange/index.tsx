import { Button, Input, notification } from 'antd'
import { useState } from 'react'
import styled from 'styled-components'
import { useProfileAndAssets } from '../../providers/profileAndAssets'
import createOrder from '../../services/order'

enum BuyOrSell {
  BUY = 'BUY',
  SELL = 'SELL',
}
const ExchangePage = () => {
  const { assets, refresh } = useProfileAndAssets()
  const [selectedAsset, setSelectedAsset] = useState('')
  const [selectBuyOrSell, setSelectBuyOrSell] = useState<BuyOrSell>()
  const [amount, setAmount] = useState<number>()
  const [isLoading, setLoading] = useState(false)

  const handleCreateOrder = async () => {
    setLoading(true)
    const asset = assets?.find((a) => a.symbol === selectedAsset)
    if (!selectBuyOrSell || !asset?.last_price || !asset?.id || !amount) {
      return
    }
    const params = {
      order_type: selectBuyOrSell,
      price_action: Number(asset.last_price),
      asset_id: asset.id,
      amount,
      user_id: 1,
    }
    const res = await createOrder(params)
    if (res?.status_code === 200) {
      refresh()
      notification['success']({
        message: 'create order success',
      })
    } else {
      notification['error']({
        message: 'create order fail',
      })
    }
    setAmount(undefined)
    setSelectBuyOrSell(undefined)
    setSelectedAsset('')
    setLoading(false)
    return true
  }
  return (
    <Container>
      select assets
      <WrapAssets>
        {assets?.map((asset) => (
          <Button
            type={asset.symbol === selectedAsset ? 'primary' : 'default'}
            onClick={() => setSelectedAsset(asset.symbol)}
          >
            <LogoAsset src={asset.logo} alt='coin' /> {asset.name}
          </Button>
        ))}
      </WrapAssets>
      <WrapBuyOrSale>
        <Button
          onClick={() => setSelectBuyOrSell(BuyOrSell.BUY)}
          style={{ background: 'green', color: '#fff' }}
        >
          Buy
        </Button>
        <Button
          onClick={() => setSelectBuyOrSell(BuyOrSell.SELL)}
          style={{ background: 'red', color: '#fff' }}
        >
          Sale
        </Button>
      </WrapBuyOrSale>
      <WrapInputAmount>
        amount:{' '}
        <Input
          type='number'
          placeholder='please input amount'
          value={amount}
          onChange={(e) => setAmount(Number(e.target.value))}
        />
      </WrapInputAmount>
      <WrapSubmit onClick={handleCreateOrder}>
        <Button type='primary' disabled={isLoading} loading={isLoading}>
          Create order
        </Button>
      </WrapSubmit>
    </Container>
  )
}

export default ExchangePage

const Container = styled.div`
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
`

const WrapAssets = styled.div`
  display: flex;
  align-items: center;
  margin-top: 16px;
  gap: 12px;
`

const LogoAsset = styled.img`
  width: 20px;
  height: 20px;
  margin-right: 8px;
  border-radius: 50%;
`

const WrapBuyOrSale = styled(WrapAssets)`
  gap: 0;
`

const WrapInputAmount = styled(WrapAssets)``

const WrapSubmit = styled(WrapAssets)``
