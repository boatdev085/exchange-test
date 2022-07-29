import { AxiosResponse } from 'axios'

import request from './request'
import { MainResponseType } from './type'

export interface DataAssets {
  id: number
  date_created: string
  date_updated?: string
  last_price: string
  logo: string
  name: string
  symbol: string
}

export interface ProfileResponse extends MainResponseType {
  data: DataAssets[]
}

const getAssets = async () => {
  try {
    const res: AxiosResponse<ProfileResponse> = await request({
      method: 'GET',
      url: '/assets',
    })
    return res.data
  } catch (e) {
    return null
  }
}

export default getAssets
