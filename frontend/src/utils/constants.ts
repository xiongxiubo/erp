export const TOKEN_KEY = 'erp_token'

export const statusText: Record<string, string> = {
  enabled: '启用',
  disabled: '停用',
  draft: '草稿',
  approved: '已审批',
  confirmed: '已确认',
  partially_inbound: '部分入库',
  partially_outbound: '部分出库',
  completed: '已完成',
  closed: '已关闭',
  voided: '已作废',
  unpaid: '未结清',
  partial: '部分结清',
  paid: '已结清',
}

export const statusColor: Record<string, string> = {
  enabled: 'green',
  disabled: 'default',
  draft: 'default',
  approved: 'blue',
  confirmed: 'green',
  partially_inbound: 'gold',
  partially_outbound: 'gold',
  completed: 'green',
  closed: 'purple',
  voided: 'red',
  unpaid: 'red',
  partial: 'gold',
  paid: 'green',
}
