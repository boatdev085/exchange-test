import { AxiosResponse } from 'axios'

import request from './request'
import { MainResponseType } from './type'

interface CreateOrderParams {
  order_type: string
  price_action: number
  asset_id: number
  amount: number
  user_id: number
}

export interface CreateOrderResponse extends MainResponseType {
  data: number
}

const createOrder = async (data: CreateOrderParams) => {
  try {
    const res: AxiosResponse<CreateOrderResponse> = await request({
      method: 'POST',
      url: '/order',
      data,
    })
    return res.data
  } catch (e) {
    return null
  }
}

export default createOrder
