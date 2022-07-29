import { AxiosResponse } from 'axios'

import request from './request'
import { MainResponseType } from './type'

export interface DataProfile {
  id: number
  date_created: string
  date_updated?: string
  first_name: string
  last_name: string
}

export interface ProfileResponse extends MainResponseType {
  data: DataProfile
}

const getProfile = async () => {
  try {
    const res: AxiosResponse<ProfileResponse> = await request({
      method: 'GET',
      url: '/profile/1',
    })
    return res.data
  } catch (e) {
    return null
  }
}

export default getProfile
