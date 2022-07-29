import { AxiosResponse } from 'axios'

import request from './request'
import { MainResponseType } from './type'

export interface DataWallet {
  amount: string
  date_created: string
  asset_id: number
  symbol: string
}

export interface WalletResponse extends MainResponseType {
  data: DataWallet[]
}

const getWallets = async () => {
  try {
    const res: AxiosResponse<WalletResponse> = await request({
      method: 'GET',
      url: '/wallet/1',
    })
    return res.data
  } catch (e) {
    return null
  }
}

export default getWallets
