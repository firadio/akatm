# API 文件清单

## 概述

本文档列出了 AKATM 项目中所有的 API 定义文件，按照新的命名规范组织。

**文件命名规则**: `{rpc}_{category}_{resource}.api`

**特殊文件**: `public.api` - 集中管理所有公开接口（无需认证）

**总文件数**: 24个

---

## 特殊文件

### public.api - 公开接口集合

| 文件名 | URL前缀 | 说明 | 详细文档 |
|-------|---------|------|---------|
| `public.api` | `/api/public/*` | 所有无需认证的公开接口 | [Public API 使用指南](./public-api-guide.md) |

**包含的接口**:
- `/api/public/admin/captcha` - 管理员验证码
- `/api/public/admin/login` - 管理员登录
- `/api/public/iam/captcha` - 用户验证码
- `/api/public/iam/login` - 用户登录
- `/api/public/iam/register` - 用户注册
- `/api/public/iam/email/send` - 发送邮箱验证码
- `/api/public/iam/email/verify` - 验证邮箱验证码

---

## 文件清单

### 1. Admin RPC（7个文件）

#### 1.1 会话管理（auth）
| 文件名 | 表名 | URL前缀 | 说明 |
|-------|------|---------|------|
| `admin_auth.api` | `admin_system_staff_sessions` | `/api/admin/auth/*` | 管理员登出、获取信息、修改密码（需要认证） |

#### 1.2 系统管理（system）
| 文件名 | 表名 | URL前缀 | 说明 |
|-------|------|---------|------|
| `admin_system_staff.api` | `admin_system_staffs` | `/api/admin/system/staff` | 员工管理 |
| `admin_system_role.api` | `admin_system_roles` | `/api/admin/system/role` | 角色管理 |
| `admin_system_menu.api` | `admin_system_menus` | `/api/admin/system/menu` | 菜单管理 |
| `admin_system_config.api` | `admin_system_configs` | `/api/admin/system/config` | 系统配置管理 |

#### 1.3 数据管理（data）
| 文件名 | 表名 | URL前缀 | 说明 |
|-------|------|---------|------|
| `admin_data_country.api` | `admin_data_countries` | `/api/admin/data/country` | 开户国家管理 |

#### 1.4 仪表板（dashboard）
| 文件名 | 表名 | URL前缀 | 说明 |
|-------|------|---------|------|
| `admin_dashboard.api` | - | `/api/admin/dashboard/*` | 管理后台仪表板数据 |

---

### 2. IAM RPC（7个文件）

#### 2.1 用户管理（user）
| 文件名 | 表名 | URL前缀 | 说明 |
|-------|------|---------|------|
| `iam_user.api` | `iam_users` | `/api/iam/user` | 用户管理（总代/代理/经理） |
| `iam_user_profile.api` | `iam_user_profiles` | `/api/iam/user/{userId}/profile` | 用户资料 |
| `iam_user_email.api` | `iam_user_emails` | `/api/iam/user/{userId}/email` | 用户邮箱绑定 |
| `iam_user_credential.api` | `iam_user_credentials` | `/api/iam/user/{userId}/credential` | 用户凭证（密码/PIN） |
| `iam_user_session.api` | `iam_user_sessions` | `/api/iam/user/{userId}/session` | 用户会话管理 |
| `iam_user_invite.api` | `iam_user_invites` | `/api/iam/user/invite` | 邀请链接管理 |
| `iam_user_country.api` | `iam_user_country_auths` | `/api/iam/user/{userId}/country` | 用户国家授权 |

---

### 3. FAMS RPC（7个文件）

#### 3.1 用户钱包（user）
| 文件名 | 表名 | URL前缀 | 说明 |
|-------|------|---------|------|
| `fams_user_wallet.api` | `fams_user_wallets` | `/api/fams/user/wallet` | 用户USDT钱包 |

#### 3.2 银行管理（bank）
| 文件名 | 表名 | URL前缀 | 说明 |
|-------|------|---------|------|
| `fams_bank_customer.api` | `fams_bank_customers` | `/api/fams/bank/customer` | 银行客户管理 |
| `fams_bank_account.api` | `fams_bank_accounts` | `/api/fams/bank/account` | 银行账户管理 |
| `fams_bank_account_application.api` | `fams_bank_account_applications` | `/api/fams/bank/account-application` | 开户申请 |
| `fams_bank_deposit.api` | `fams_bank_deposits` | `/api/fams/bank/deposit` | 银行存款审核 |
| `fams_bank_withdrawal.api` | `fams_bank_withdrawals` | `/api/fams/bank/withdrawal` | 银行提现审核 |

#### 3.3 资金管理（fund）
| 文件名 | 表名 | URL前缀 | 说明 |
|-------|------|---------|------|
| `fams_fund_detail.api` | `fams_fund_details` | `/api/fams/fund/detail` | 资金明细 |

---

## 文件路径

所有 API 文件位于: `api/admin/docs/`

```
api/admin/docs/
├── public.api                            # 【特殊】公开接口集合
├── admin_auth.api
├── admin_system_staff.api
├── admin_system_role.api
├── admin_system_menu.api
├── admin_system_config.api
├── admin_data_country.api
├── admin_dashboard.api
├── iam_user.api
├── iam_user_profile.api
├── iam_user_email.api
├── iam_user_credential.api
├── iam_user_session.api
├── iam_user_invite.api
├── iam_user_country.api
├── fams_user_wallet.api
├── fams_bank_customer.api
├── fams_bank_account.api
├── fams_bank_account_application.api
├── fams_bank_deposit.api
├── fams_bank_withdrawal.api
└── fams_fund_detail.api
```

---

## RPC 服务映射

### Public（公开接口）
**文件数**: 1
**URL前缀**: `/api/public/*`
**认证要求**: ❌ 无需认证

### Admin RPC
**服务名**: `admin`
**表前缀**: `admin_`
**文件数**: 7
**URL前缀**: `/api/admin/*`
**认证要求**: ✅ 需要认证

### IAM RPC
**服务名**: `iam`
**表前缀**: `iam_`
**文件数**: 7
**URL前缀**: `/api/iam/*`
**认证要求**: ✅ 需要认证

### FAMS RPC
**服务名**: `fams`
**表前缀**: `fams_`
**文件数**: 7
**URL前缀**: `/api/fams/*`
**认证要求**: ✅ 需要认证

---

## 分类统计

### 按分类（category）统计

| 分类 | 文件数 | 说明 |
|------|--------|------|
| `auth` | 1 | 认证相关 |
| `system` | 4 | 系统管理 |
| `data` | 1 | 基础数据 |
| `dashboard` | 1 | 仪表板 |
| `user` | 7 | 用户相关 |
| `bank` | 5 | 银行业务 |
| `fund` | 1 | 资金管理 |

### 按资源（resource）统计

| 资源 | 文件数 | 涉及RPC |
|------|--------|---------|
| user | 2 | iam, iam_user |
| user_* | 6 | iam (profile/email/credential/session/invite/country) |
| staff | 1 | admin |
| role | 1 | admin |
| menu | 1 | admin |
| country | 2 | admin_data, iam_user |
| bank_* | 5 | fams (customer/account/application/deposit/withdrawal) |
| wallet | 1 | fams_user |
| fund | 1 | fams |

---

## 公开接口文件

**所有公开接口集中在 `public.api` 文件中**：

| 文件名 | 公开接口数量 | 接口列表 |
|-------|------------|---------|
| `public.api` | 7个 | 详见 [Public API 使用指南](./public-api-guide.md) |

**具体接口**：
- `/api/public/admin/captcha` - 管理员验证码
- `/api/public/admin/login` - 管理员登录
- `/api/public/iam/captcha` - 用户验证码
- `/api/public/iam/login` - 用户登录
- `/api/public/iam/register` - 用户注册
- `/api/public/iam/email/send` - 发送邮箱验证码
- `/api/public/iam/email/verify` - 验证邮箱验证码

---

## 已删除的文件

以下旧文件已被重命名或合并：

| 旧文件名 | 新文件名/状态 | 说明 |
|---------|-------------|------|
| `auth.api` | `admin_auth.api` | 重命名 |
| `staff.api` | `admin_system_staff.api` | 重命名（sys→system） |
| `role.api` | `admin_system_role.api` | 重命名（sys→system） |
| `menu.api` | `admin_system_menu.api` | 重命名（sys→system） |
| `system.api` | `admin_system_config.api` | 重命名 |
| `country.api` | `admin_data_country.api` | 重命名 |
| `user.api` | `iam_user.api` | 重命名 |
| `invite.api` | `iam_user_invite.api` | 重命名 |
| `customer.api` | `fams_bank_customer.api` | 重命名 |
| `account.api` | `fams_bank_account.api` | 重命名 |
| `accountApplication.api` | `fams_bank_account_application.api` | 重命名 |
| `deposit.api` | `fams_bank_deposit.api` | 重命名 |
| `withdrawal.api` | `fams_bank_withdrawal.api` | 重命名 |
| `fund.api` | `fams_fund_detail.api` | 重命名 |
| - | `iam_auth_email.api` | ❌ 删除（合并到 public.api） |
| `audit.api` | `admin_audit.api` | 保留（日志审计） |
| `dashboard.api` | `admin_dashboard.api` | 保留（仪表板） |
| `report.api` | `admin_report.api` | 保留（报表） |

---

## 待创建的文件

根据表结构，以下文件可能需要创建（根据业务需求）：

### Admin RPC
- `admin_system_permission.api` - 权限管理
- `admin_system_staff_log.api` - 员工操作日志

### IAM RPC
- ✅ `iam_user_profile.api` - 已创建
- ✅ `iam_user_email.api` - 已创建
- ✅ `iam_user_credential.api` - 已创建
- ✅ `iam_user_session.api` - 已创建
- ✅ `iam_user_country.api` - 已创建

### FAMS RPC
- ✅ `fams_user_wallet.api` - 已创建
- `fams_user_wallet_ledger.api` - 钱包账变记录（可选，已包含在wallet中）
- `fams_user_wallet_withdrawal.api` - 钱包提现（可选，已包含在wallet中）
- `fams_user_wallet_address.api` - 提现地址管理
- `fams_user_wallet_deposit.api` - 钱包存款（可选）
- `fams_user_setting.api` - 用户手续费设置
- `fams_bank_webhook_record.api` - Webhook记录

---

## 命名规范总结

### 文件命名
```
{rpc}_{category}_{resource}.api
```

### URL路径
```
/api/{rpc}/{category}/{resource}[/:id][/{action}]
```

### Group命名
```
{rpc}{Category}{Resource}
```
示例: `adminSystemStaff`, `iamUserProfile`, `famsBankCustomer`

### Handler命名
```
{action}{Resource}
```
示例: `createStaff`, `getUserProfile`, `listBankCustomers`

---

**更新日期**: 2025-10-22
**版本**: v2.0
