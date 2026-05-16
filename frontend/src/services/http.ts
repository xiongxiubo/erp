import axios from 'axios'
import { message } from 'antd'
import { TOKEN_KEY } from '../utils/constants'
import type { ApiResponse } from '../types/common'

export const http = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL ?? '/api/v1',
  timeout: 15000,
})

http.interceptors.request.use((config) => {
  const token = localStorage.getItem(TOKEN_KEY)
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

http.interceptors.response.use(
  (response) => {
    const body = response.data as ApiResponse<unknown>
    if (body?.code && body.code !== 0) {
      return Promise.reject(new Error(body.message || '请求失败'))
    }
    return response
  },
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem(TOKEN_KEY)
      if (location.pathname !== '/login') {
        location.href = '/login'
      }
    }
    message.error(error.response?.data?.message ?? error.message ?? '请求失败')
    return Promise.reject(error)
  },
)

export async function getData<T>(url: string, params?: object) {
  const res = await http.get<ApiResponse<T>>(url, { params })
  return res.data.data
}

export async function postData<T>(url: string, data?: object) {
  const res = await http.post<ApiResponse<T>>(url, data)
  return res.data.data
}

export async function putData<T>(url: string, data?: object) {
  const res = await http.put<ApiResponse<T>>(url, data)
  return res.data.data
}

export async function deleteData<T>(url: string) {
  const res = await http.delete<ApiResponse<T>>(url)
  return res.data.data
}
