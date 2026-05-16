import { getData, postData } from './http'
import type { LoginResponse, MenuNode, UserProfile } from '../types/common'

export function login(data: { username: string; password: string }) {
  return postData<LoginResponse>('/auth/login', data)
}

export function logout() {
  return postData('/auth/logout')
}

export function getProfile() {
  return getData<UserProfile>('/auth/profile')
}

export function getMenus() {
  return getData<MenuNode[]>('/auth/menus')
}

export function getPermissions() {
  return getData<string[]>('/auth/permissions')
}
