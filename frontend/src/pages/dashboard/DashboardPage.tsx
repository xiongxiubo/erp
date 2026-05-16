import { Card, Col, Row, Statistic, Table, Typography } from 'antd'
import { useEffect, useState } from 'react'
import { getDashboardSummary, getRecentDocuments } from '../../services/dashboard'
import type { DashboardSummary, RecentDocument } from '../../types/common'
import { money } from '../../utils/format'
import { StatusTag } from '../../components/StatusTag'

export function DashboardPage() {
  const [summary, setSummary] = useState<DashboardSummary>()
  const [docs, setDocs] = useState<RecentDocument[]>([])

  useEffect(() => {
    getDashboardSummary().then(setSummary).catch(() => setSummary(fallbackSummary))
    getRecentDocuments().then(setDocs).catch(() => setDocs(fallbackDocs))
  }, [])

  return (
    <div className="erp-page dashboard-page">
      <section className="dashboard-hero">
        <div>
          <Typography.Text className="erp-kicker">TODAY'S OPERATING SIGNAL</Typography.Text>
          <Typography.Title>从库存脉搏看经营现金流。</Typography.Title>
          <p>采购入库、销售出库、库存流水、应收应付在同一张作战图里联动。</p>
        </div>
        <div className="hero-number">{money(summary?.monthSales)}</div>
      </section>
      <Row gutter={[18, 18]}>
        <Metric title="今日销售" value={summary?.todaySales} />
        <Metric title="本月销售" value={summary?.monthSales} />
        <Metric title="本月采购" value={summary?.monthPurchase} />
        <Metric title="库存金额" value={summary?.inventoryValue} />
        <Metric title="应收余额" value={summary?.receivableAmount} />
        <Metric title="应付余额" value={summary?.payableAmount} />
      </Row>
      <Card className="erp-card" title="最近业务单据">
        <Table
          rowKey={(row) => `${row.type}-${row.code}`}
          dataSource={docs}
          pagination={false}
          columns={[
            { title: '类型', dataIndex: 'type' },
            { title: '单号', dataIndex: 'code' },
            { title: '状态', dataIndex: 'status', render: (value) => <StatusTag value={value} /> },
            { title: '金额', dataIndex: 'amount', align: 'right', render: money },
          ]}
        />
      </Card>
    </div>
  )
}

function Metric({ title, value }: { title: string; value?: number }) {
  return (
    <Col xs={24} sm={12} xl={8}>
      <Card className="metric-card">
        <Statistic title={title} value={value ?? 0} precision={2} prefix="¥" />
      </Card>
    </Col>
  )
}

const fallbackSummary: DashboardSummary = {
  todaySales: 0,
  monthSales: 0,
  monthPurchase: 0,
  receivableAmount: 0,
  payableAmount: 0,
  inventoryValue: 0,
  lowStockCount: 0,
}

const fallbackDocs: RecentDocument[] = []
