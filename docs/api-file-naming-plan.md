# API 文件命名重构计划

## 命名规则

**文件名格式**: `{rpc}_{category}_{resource}.api`

**对应关系**:
- 文件名 = 表名（去掉复数的s）
- 例如: 表 `admin_sys_staffs` → 文件 `admin_sys_staff.api`

---

## 文件重命名计划

### Admin RPC - System Management (sys)

| 当前文件名 | 新文件名 | 表名 | 说明 |
|-----------|---------|------|------|
| `auth.api` | `admin_auth.api` | `admin_sys_staff_sessions` | 管理员认证（放在admin级别） |
| `staff.api` | `admin_sys_staff.api` | `admin_sys_staffs` | 员工管理 |
| `role.api` | `admin_sys_role.api` | `admin_sys_roles` | 角色管理 |
| `menu.api` | `admin_sys_menu.api` | `admin_sys_menus` | 菜单管理 |
| - | `admin_sys_permission.api` | `admin_sys_permissions` | 权限管理（新建） |
| - | `admin_sys_staff_log.api` | `admin_sys_staff_logs` | 员工操作日志（新建） |

### Admin RPC - Data Management (data)

| 当前文件名 | 新文件名 | 表名 | 说明 |
|-----------|---------|------|------|
| `country.api` | `admin_data_country.api` | `admin_data_countries` | 国家管理 |

### IAM RPC - Authentication & Email (auth/email)

| 当前文件名 | 新文件名 | 表名 | 说明 |
|-----------|---------|------|------|
| - | `iam_auth_email.api` | - | 邮箱验证码（公开接口，新建） |

### IAM RPC - User Management (user)

| 当前文件名 | 新文件名 | 表名 | 说明 |
|-----------|---------|------|------|
| `user.api` | `iam_user.api` | `iam_users` | 用户管理（总代/代理/经理） |
| - | `iam_user_profile.api` | `iam_user_profiles` | 用户资料（新建） |
| - | `iam_user_email.api` | `iam_user_emails` | 用户邮箱（新建） |
| - | `iam_user_credential.api` | `iam_user_credentials` | 用户凭证（新建） |
| - | `iam_user_session.api` | `iam_user_sessions` | 用户会话（新建） |
| `invite.api` | `iam_user_invite.api` | `iam_user_invites` | 邀请管理 |
| - | `iam_user_country.api` | `iam_user_country_auths` | 国家授权（新建） |

### FAMS RPC - User Wallet Management (user)

| 当前文件名 | 新文件名 | 表名 | 说明 |
|-----------|---------|------|------|
| - | `fams_user_setting.api` | `fams_user_settings` | 用户设置（新建） |
| - | `fams_user_wallet.api` | `fams_user_wallets` | 用户钱包（新建） |
| - | `fams_user_wallet_ledger.api` | `fams_user_wallet_ledgers` | 钱包账变记录（新建） |
| - | `fams_user_wallet_withdrawal.api` | `fams_user_wallet_withdrawals` | 钱包提现（新建） |
| - | `fams_user_wallet_address.api` | `fams_user_wallet_addresses` | 提现地址（新建） |
| - | `fams_user_wallet_deposit.api` | `fams_user_wallet_deposits` | 钱包存款（新建） |

### FAMS RPC - Bank Management (bank)

| 当前文件名 | 新文件名 | 表名 | 说明 |
|-----------|---------|------|------|
| `customer.api` | `fams_bank_customer.api` | `fams_bank_customers` | 银行客户 |
| `account.api` | `fams_bank_account.api` | `fams_bank_accounts` | 银行账户 |
| `accountApplication.api` | `fams_bank_account_application.api` | `fams_bank_account_applications` | 开户申请 |
| `deposit.api` | `fams_bank_deposit.api` | `fams_bank_deposits` | 银行存款 |
| `withdrawal.api` | `fams_bank_withdrawal.api` | `fams_bank_withdrawals` | 银行提现 |
| - | `fams_bank_webhook_record.api` | `fams_bank_webhook_records` | Webhook记录（新建） |

### FAMS RPC - Fund Management (fund)

| 当前文件名 | 新文件名 | 表名 | 说明 |
|-----------|---------|------|------|
| `fund.api` | `fams_fund_detail.api` | `fams_fund_details` | 资金明细 |

---

## 重命名执行计划

### 第一步：重命名现有文件

```bash
# Admin RPC
mv api/admin/docs/auth.api api/admin/docs/admin_auth.api
mv api/admin/docs/staff.api api/admin/docs/admin_sys_staff.api
mv api/admin/docs/role.api api/admin/docs/admin_sys_role.api
mv api/admin/docs/menu.api api/admin/docs/admin_sys_menu.api
mv api/admin/docs/country.api api/admin/docs/admin_data_country.api

# IAM RPC
mv api/admin/docs/user.api api/admin/docs/iam_user.api
mv api/admin/docs/invite.api api/admin/docs/iam_user_invite.api

# FAMS RPC
mv api/admin/docs/customer.api api/admin/docs/fams_bank_customer.api
mv api/admin/docs/account.api api/admin/docs/fams_bank_account.api
mv api/admin/docs/accountApplication.api api/admin/docs/fams_bank_account_application.api
mv api/admin/docs/deposit.api api/admin/docs/fams_bank_deposit.api
mv api/admin/docs/withdrawal.api api/admin/docs/fams_bank_withdrawal.api
mv api/admin/docs/fund.api api/admin/docs/fams_fund_detail.api
```

### 第二步：创建新文件

需要创建的新API文件：

**Admin RPC**:
- admin_sys_permission.api
- admin_sys_staff_log.api

**IAM RPC**:
- iam_auth_email.api
- iam_user_profile.api
- iam_user_email.api
- iam_user_credential.api
- iam_user_session.api
- iam_user_country.api

**FAMS RPC**:
- fams_user_setting.api
- fams_user_wallet.api
- fams_user_wallet_ledger.api
- fams_user_wallet_withdrawal.api
- fams_user_wallet_address.api
- fams_user_wallet_deposit.api
- fams_bank_webhook_record.api

### 第三步：删除不需要的文件

```bash
rm api/admin/docs/audit.api
rm api/admin/docs/dashboard.api
rm api/admin/docs/report.api
rm api/admin/docs/system.api
```

---

## URL 路径映射

### 文件名与URL路径的关系

| 文件名 | URL路径 | 说明 |
|-------|---------|------|
| `admin_auth.api` | `/api/public/admin/*`, `/api/admin/auth/*` | 认证接口 |
| `admin_sys_staff.api` | `/api/admin/sys/staff` | 员工管理 |
| `admin_sys_role.api` | `/api/admin/sys/role` | 角色管理 |
| `admin_sys_menu.api` | `/api/admin/sys/menu` | 菜单管理 |
| `admin_data_country.api` | `/api/admin/data/country` | 国家管理 |
| `iam_auth_email.api` | `/api/public/iam/email/*` | 邮箱验证 |
| `iam_user.api` | `/api/iam/user` | 用户管理 |
| `iam_user_invite.api` | `/api/iam/user/invite` | 邀请管理 |
| `fams_bank_customer.api` | `/api/fams/bank/customer` | 银行客户 |
| `fams_bank_account.api` | `/api/fams/bank/account` | 银行账户 |
| `fams_user_wallet.api` | `/api/fams/user/wallet` | 用户钱包 |

---

## 优势

1. **文件名即表名** - 一目了然知道操作的是哪张表
2. **命名空间清晰** - 通过前缀区分 admin/iam/fams
3. **分类明确** - sys/data/user/bank 等分类清晰
4. **易于查找** - 按字母排序后，相关文件聚合在一起
5. **代码生成友好** - 可以根据文件名自动生成对应的代码

---

**更新日期**: 2025-10-22
