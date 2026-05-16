export function money(value?: number) {
  return new Intl.NumberFormat('zh-CN', { style: 'currency', currency: 'CNY' }).format(value ?? 0)
}

export function number(value?: number) {
  return new Intl.NumberFormat('zh-CN', { maximumFractionDigits: 2 }).format(value ?? 0)
}
