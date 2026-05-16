import { getData } from './http'
import type { DashboardSummary, RecentDocument } from '../types/common'

export function getDashboardSummary() {
  return getData<DashboardSummary>('/dashboard/summary')
}

export function getRecentDocuments() {
  return getData<RecentDocument[]>('/dashboard/recent-documents')
}
