import axios, { AxiosRequestConfig, AxiosError } from 'axios'
interface RequestProps extends AxiosRequestConfig {
  json?: boolean
  baseURL: string
}

const initialRequest = (props: RequestProps) => {
  const { method, data, baseURL, responseType } = props
  const dataOrParams = ['GET', 'DELETE'].includes(method || 'GET')
    ? 'params'
    : 'data'
  return axios.create({
    baseURL: baseURL,
    withCredentials: false,
    responseType,
    method: method || 'GET',
    [dataOrParams]: data,
  })
}

export const onRejectedFunc = async (res: AxiosError) => {
  return Promise.reject(res)
}

const request = initialRequest({
  baseURL: `http://localhost:8081`,
})

request.interceptors.response.use(undefined, onRejectedFunc)

export default request
