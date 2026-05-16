import {
  AccountBookOutlined,
  DashboardOutlined,
  DatabaseOutlined,
  InboxOutlined,
  LogoutOutlined,
  SafetyCertificateOutlined,
  ShopOutlined,
  ShoppingCartOutlined,
  UserOutlined,
} from "@ant-design/icons";
import { Avatar, Dropdown, Layout, Menu, Typography, type MenuProps } from "antd";
import { Outlet, useLocation, useNavigate } from "react-router-dom";
import { useAuth } from "../stores/authStore";
import type { MenuNode } from "../types/common";

const icons = {
  DashboardOutlined: <DashboardOutlined />,
  SafetyCertificateOutlined: <SafetyCertificateOutlined />,
  DatabaseOutlined: <DatabaseOutlined />,
  ShoppingCartOutlined: <ShoppingCartOutlined />,
  ShopOutlined: <ShopOutlined />,
  InboxOutlined: <InboxOutlined />,
  AccountBookOutlined: <AccountBookOutlined />,
};

function toMenuItems(nodes: MenuNode[]): MenuProps["items"] {
  return nodes.map((item) => ({
    key: item.path,
    icon: icons[item.icon as keyof typeof icons],
    label: item.name,
    children: item.children?.length ? toMenuItems(item.children) : undefined,
  }));
}

export function BasicLayout() {
  const { menus, user, clearSession } = useAuth();
  const navigate = useNavigate();
  const location = useLocation();
  const items = menus.length ? toMenuItems(menus) : toMenuItems(defaultMenus);

  return (
    <Layout className="erp-shell">
      <Layout.Sider width={248} className="erp-sider">
        <div className="erp-brand">
          <div className="erp-brand-mark">鼎</div>
          <div>
            <strong>鼎衡 ERP</strong>
            <span>Flow Control Console</span>
          </div>
        </div>
        <Menu theme="dark" mode="inline" selectedKeys={[location.pathname]} items={items} onClick={({ key }) => navigate(key)} />
      </Layout.Sider>
      <Layout>
        <Layout.Header className="erp-topbar">
          <div>
            <Typography.Text className="erp-kicker">核心进销存财务闭环</Typography.Text>
            <Typography.Title level={4}>经营驾驶舱</Typography.Title>
          </div>
          <Dropdown
            menu={{
              items: [{ key: "logout", icon: <LogoutOutlined />, label: "退出登录" }],
              onClick: () => {
                clearSession();
                navigate("/login");
              },
            }}>
            <div className="erp-user">
              <Avatar icon={<UserOutlined />} />
              <span>{user?.name ?? "管理员"}</span>
            </div>
          </Dropdown>
        </Layout.Header>
        <Layout.Content className="erp-content">
          <Outlet />
        </Layout.Content>
      </Layout>
    </Layout>
  );
}

const defaultMenus: MenuNode[] = [
  { id: 1, parentId: 0, name: "看板", path: "/dashboard", icon: "DashboardOutlined", sort: 1, permissionCode: "" },
  {
    id: 2,
    parentId: 0,
    name: "系统管理",
    path: "/system",
    icon: "SafetyCertificateOutlined",
    sort: 2,
    permissionCode: "",
    children: [
      { id: 21, parentId: 2, name: "用户管理", path: "/system/users", icon: "", sort: 1, permissionCode: "" },
      { id: 22, parentId: 2, name: "角色管理", path: "/system/roles", icon: "", sort: 2, permissionCode: "" },
      { id: 23, parentId: 2, name: "菜单管理", path: "/system/menus", icon: "", sort: 3, permissionCode: "" },
      { id: 24, parentId: 2, name: "部门管理", path: "/system/departments", icon: "", sort: 4, permissionCode: "" },
    ],
  },
  {
    id: 3,
    parentId: 0,
    name: "基础资料",
    path: "/master",
    icon: "DatabaseOutlined",
    sort: 3,
    permissionCode: "",
    children: [
      { id: 31, parentId: 3, name: "客户", path: "/master/customers", icon: "", sort: 1, permissionCode: "" },
      { id: 32, parentId: 3, name: "供应商", path: "/master/suppliers", icon: "", sort: 2, permissionCode: "" },
      { id: 33, parentId: 3, name: "商品", path: "/master/products", icon: "", sort: 3, permissionCode: "" },
      { id: 34, parentId: 3, name: "仓库", path: "/master/warehouses", icon: "", sort: 4, permissionCode: "" },
    ],
  },
  {
    id: 4,
    parentId: 0,
    name: "采购",
    path: "/purchase",
    icon: "ShoppingCartOutlined",
    sort: 4,
    permissionCode: "",
    children: [
      { id: 41, parentId: 4, name: "采购订单", path: "/purchase/orders", icon: "", sort: 1, permissionCode: "" },
      { id: 42, parentId: 4, name: "采购入库", path: "/purchase/inbounds", icon: "", sort: 2, permissionCode: "" },
    ],
  },
  {
    id: 5,
    parentId: 0,
    name: "销售",
    path: "/sales",
    icon: "ShopOutlined",
    sort: 5,
    permissionCode: "",
    children: [
      { id: 51, parentId: 5, name: "销售订单", path: "/sales/orders", icon: "", sort: 1, permissionCode: "" },
      { id: 52, parentId: 5, name: "销售出库", path: "/sales/outbounds", icon: "", sort: 2, permissionCode: "" },
    ],
  },
  {
    id: 6,
    parentId: 0,
    name: "库存",
    path: "/inventory",
    icon: "InboxOutlined",
    sort: 6,
    permissionCode: "",
    children: [
      { id: 61, parentId: 6, name: "库存查询", path: "/inventory/stocks", icon: "", sort: 1, permissionCode: "" },
      { id: 62, parentId: 6, name: "库存流水", path: "/inventory/ledgers", icon: "", sort: 2, permissionCode: "" },
    ],
  },
  {
    id: 7,
    parentId: 0,
    name: "财务",
    path: "/finance",
    icon: "AccountBookOutlined",
    sort: 7,
    permissionCode: "",
    children: [
      { id: 71, parentId: 7, name: "应收", path: "/finance/receivables", icon: "", sort: 1, permissionCode: "" },
      { id: 72, parentId: 7, name: "应付", path: "/finance/payables", icon: "", sort: 2, permissionCode: "" },
    ],
  },
];
