import { deleteData, getData, postData, putData } from './http'
import type { BaseEntity, PageResult } from '../types/common'

export function listResource<T = BaseEntity>(path: string, params: object) {
  return getData<PageResult<T>>(path, params)
}

export function createResource<T = BaseEntity>(path: string, data: object) {
  return postData<T>(path, data)
}

export function updateResource(path: string, id: number, data: object) {
  return putData(`${path}/${id}`, data)
}

export function deleteResource(path: string, id: number) {
  return deleteData(`${path}/${id}`)
}

export function actionResource(path: string, id: number, action: string) {
  return postData(`${path}/${id}/${action}`)
}
