# AKATM API 接口规范

## 一、概述

本文档定义了 AKATM 项目的 API 接口规范，包括 URL 格式、认证方式、数据格式、错误处理等。

**设计原则**：
- ✅ RESTful 风格
- ✅ 全局 Cookie 认证（非 JWT）
- ✅ 统一错误响应格式
- ✅ 时间统一使用毫秒级时间戳 int64
- ✅ 前后端同域部署，无需手动处理认证

---

## 二、URL 格式规范

### 2.1 标准格式

```
/api/{rpc}/{category}/{resource}[/:id][/{action}]
```

**参数说明**：

| 参数 | 说明 | 示例 |
|------|------|------|
| `{rpc}` | RPC 服务名（表前缀） | `admin`, `iam`, `fams` |
| `{category}` | 对象分类（对应 table/{category}.go 文件名） | `sys`, `data`, `user`, `bank`, `email` |
| `{resource}` | 资源名称（Struct 名去掉分类前缀后小写） | `staff`, `role`, `wallet`, `customer` |
| `:id` | 资源 ID（可选） | `123456789` |
| `{action}` | 操作动作（可选） | `role`, `password`, `status` |

### 2.2 表名与 URL 映射规则

#### 规则说明

**数据库表名格式**：`{rpc}_{category}_{resource}s`（复数）

**URL 路径格式**：`/api/{rpc}/{category}/{resource}`（单数）

#### 映射示例

| 数据库表名 | Go Struct | 文件名 | URL 路径 |
|-----------|----------|--------|---------|
| `admin_sys_staffs` | `SysStaff` | `sys.go` | `/api/admin/sys/staff` |
| `admin_sys_roles` | `SysRole` | `sys.go` | `/api/admin/sys/role` |
| `admin_sys_menus` | `SysMenu` | `sys.go` | `/api/admin/sys/menu` |
| `admin_data_countries` | `DataCountry` | `data.go` | `/api/admin/data/country` |
| `iam_users` | `User` | `user.go` | `/api/iam/user` |
| `iam_user_emails` | `UserEmail` | `user.go` | `/api/iam/user/email` |
| `iam_user_invites` | `UserInvite` | `user.go` | `/api/iam/user/invite` |
| `iam_email_verifies` | `EmailVerify` | `email.go` | `/api/iam/email/verify` |
| `fams_user_wallets` | `UserWallet` | `user.go` | `/api/fams/user/wallet` |
| `fams_user_wallet_addresses` | `UserWalletAddress` | `user.go` | `/api/fams/user/wallet-address` |
| `fams_bank_customers` | `BankCustomer` | `bank.go` | `/api/fams/bank/customer` |
| `fams_bank_accounts` | `BankAccount` | `bank.go` | `/api/fams/bank/account` |
| `fams_bank_deposits` | `BankDeposit` | `bank.go` | `/api/fams/bank/deposit` |

**特殊情况处理**：
- 对于 `iam_users` 表（Struct 为 `User`），因为文件名就是 `user.go`，资源名也是 `user`，所以 URL 为 `/api/iam/user`（不重复）
- 对于多词资源名使用中划线连接，如 `UserWalletAddress` → `wallet-address`

### 2.3 HTTP 方法与操作映射

| HTTP 方法 | 操作 | URL 示例 | 说明 |
|-----------|------|---------|------|
| `GET` | 列表查询 | `GET /api/admin/sys/staff` | 分页获取员工列表 |
| `GET` | 详情查询 | `GET /api/admin/sys/staff/:id` | 获取指定员工详情 |
| `POST` | 创建资源 | `POST /api/admin/sys/staff` | 创建新员工 |
| `PUT` | 完整更新 | `PUT /api/admin/sys/staff/:id` | 更新员工完整信息 |
| `PATCH` | 部分更新/关系操作 | `PATCH /api/admin/sys/staff/:id/role` | 设置员工角色 |
| `PATCH` | 部分更新/关系操作 | `PATCH /api/admin/sys/staff/:id/status` | 更新员工状态 |
| `DELETE` | 删除资源 | `DELETE /api/admin/sys/staff/:id` | 删除员工（软删除） |

**PATCH 方法使用场景**：
- ✅ 更新资源的单个字段（如状态、密码）
- ✅ 设置资源的关联关系（如角色、权限、菜单）
- ✅ 执行特定的业务操作（如审核、冻结、解冻）

---

## 三、完整 API 接口清单

### 3.1 Admin RPC（管理后台服务）

#### 3.1.1 系统管理（sys）

##### **员工管理** - `/api/admin/sys/staff`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/admin/sys/staff` | 创建员工 |
| `GET` | `/api/admin/sys/staff` | 员工列表（分页、筛选） |
| `GET` | `/api/admin/sys/staff/:id` | 员工详情 |
| `PUT` | `/api/admin/sys/staff/:id` | 更新员工信息 |
| `DELETE` | `/api/admin/sys/staff/:id` | 删除员工 |
| `PATCH` | `/api/admin/sys/staff/:id/password` | 重设密码 |
| `PATCH` | `/api/admin/sys/staff/:id/role` | 设置角色 |
| `PATCH` | `/api/admin/sys/staff/:id/status` | 更新状态 |
| `GET` | `/api/admin/sys/staff/:id/menu` | 获取员工菜单 |

##### **角色管理** - `/api/admin/sys/role`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/admin/sys/role` | 创建角色 |
| `GET` | `/api/admin/sys/role` | 角色列表 |
| `GET` | `/api/admin/sys/role/:id` | 角色详情 |
| `PUT` | `/api/admin/sys/role/:id` | 更新角色 |
| `DELETE` | `/api/admin/sys/role/:id` | 删除角色 |
| `PATCH` | `/api/admin/sys/role/:id/menu` | 分配菜单 |
| `GET` | `/api/admin/sys/role/:id/menu` | 获取角色菜单 |

##### **菜单管理** - `/api/admin/sys/menu`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/admin/sys/menu` | 创建菜单 |
| `GET` | `/api/admin/sys/menu` | 菜单列表（树形结构） |
| `GET` | `/api/admin/sys/menu/:id` | 菜单详情 |
| `PUT` | `/api/admin/sys/menu/:id` | 更新菜单 |
| `DELETE` | `/api/admin/sys/menu/:id` | 删除菜单 |
| `PATCH` | `/api/admin/sys/menu/:id/status` | 更新菜单状态 |

##### **权限管理** - `/api/admin/sys/permission`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/admin/sys/permission` | 创建权限 |
| `GET` | `/api/admin/sys/permission` | 权限列表 |
| `GET` | `/api/admin/sys/permission/:id` | 权限详情 |
| `PUT` | `/api/admin/sys/permission/:id` | 更新权限 |
| `DELETE` | `/api/admin/sys/permission/:id` | 删除权限 |

##### **操作日志** - `/api/admin/sys/staff-log`

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/api/admin/sys/staff-log` | 查询员工操作日志 |
| `GET` | `/api/admin/sys/staff-log/:id` | 日志详情 |

#### 3.1.2 数据管理（data）

##### **国家管理** - `/api/admin/data/country`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/admin/data/country` | 创建国家 |
| `GET` | `/api/admin/data/country` | 国家列表 |
| `GET` | `/api/admin/data/country/:id` | 国家详情 |
| `PUT` | `/api/admin/data/country/:id` | 更新国家 |
| `DELETE` | `/api/admin/data/country/:id` | 删除国家 |
| `PATCH` | `/api/admin/data/country/:id/status` | 更新国家状态 |

#### 3.1.3 认证接口

##### **管理员认证** - `/api/public/admin`（公开）

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/api/public/admin/captcha` | 获取验证码 |
| `POST` | `/api/public/admin/login` | 管理员登录 |

##### **管理员会话** - `/api/admin/auth`（需要认证）

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/admin/auth/logout` | 登出 |
| `GET` | `/api/admin/auth/info` | 获取当前管理员信息 |
| `PATCH` | `/api/admin/auth/password` | 修改密码 |

---

### 3.2 IAM RPC（身份认证服务）

#### 3.2.1 用户管理（user）

##### **用户** - `/api/iam/user`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/iam/user/super-agent` | 创建总代 |
| `POST` | `/api/iam/user/agent` | 创建代理 |
| `POST` | `/api/iam/user/manager` | 创建客户经理 |
| `GET` | `/api/iam/user` | 用户列表 |
| `GET` | `/api/iam/user/super-agent` | 总代列表 |
| `GET` | `/api/iam/user/agent` | 代理列表 |
| `GET` | `/api/iam/user/manager` | 客户经理列表 |
| `GET` | `/api/iam/user/:id` | 用户详情 |
| `PUT` | `/api/iam/user/:id` | 更新用户 |
| `DELETE` | `/api/iam/user/:id` | 删除用户 |
| `PATCH` | `/api/iam/user/:id/status` | 更新状态 |
| `PATCH` | `/api/iam/user/:id/fee` | 设置手续费 |
| `PATCH` | `/api/iam/user/:id/country` | 授权国家 |

##### **用户资料** - `/api/iam/user/profile`

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/api/iam/user/:id/profile` | 获取用户资料 |
| `PUT` | `/api/iam/user/:id/profile` | 更新用户资料 |

##### **用户邮箱** - `/api/iam/user/email`

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/api/iam/user/:id/email` | 获取用户邮箱列表 |
| `POST` | `/api/iam/user/:id/email` | 绑定邮箱 |
| `DELETE` | `/api/iam/user/:id/email/:emailId` | 删除邮箱 |

##### **用户凭证** - `/api/iam/user/credential`

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/api/iam/user/:id/credential` | 获取用户凭证 |
| `POST` | `/api/iam/user/:id/credential` | 创建凭证 |
| `PATCH` | `/api/iam/user/:id/credential/:credentialId` | 更新凭证 |
| `DELETE` | `/api/iam/user/:id/credential/:credentialId` | 删除凭证 |

##### **邀请管理** - `/api/iam/user/invite`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/iam/user/invite` | 生成邀请链接 |
| `GET` | `/api/iam/user/invite` | 邀请记录列表 |
| `GET` | `/api/iam/user/invite/:id` | 邀请详情 |
| `PATCH` | `/api/iam/user/invite/:id/disable` | 禁用邀请链接 |

##### **国家授权** - `/api/iam/user/country-auth`

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/api/iam/user/:id/country-auth` | 获取用户授权国家 |
| `POST` | `/api/iam/user/:id/country-auth` | 授权国家 |
| `DELETE` | `/api/iam/user/:id/country-auth/:countryId` | 取消授权 |

#### 3.2.2 邮箱验证（email）

##### **邮箱验证** - `/api/public/iam/email`（公开）

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/public/iam/email/send` | 发送验证码 |
| `POST` | `/api/public/iam/email/verify` | 验证邮箱 |

#### 3.2.3 认证接口

##### **用户认证** - `/api/public/iam`（公开）

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/public/iam/login` | 用户登录 |
| `POST` | `/api/public/iam/register` | 用户注册（通过邀请码） |

##### **用户会话** - `/api/iam/auth`（需要认证）

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/iam/auth/logout` | 登出 |
| `GET` | `/api/iam/auth/info` | 获取当前用户信息 |
| `PATCH` | `/api/iam/auth/password` | 修改密码 |

---

### 3.3 FAMS RPC（金融账户服务）

#### 3.3.1 用户钱包（user）

##### **钱包管理** - `/api/fams/user/wallet`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/fams/user/wallet` | 创建钱包 |
| `GET` | `/api/fams/user/wallet` | 钱包列表 |
| `GET` | `/api/fams/user/wallet/:id` | 钱包详情 |
| `GET` | `/api/fams/user/wallet/:id/balance` | 钱包余额 |
| `GET` | `/api/fams/user/wallet/:id/ledger` | 账变记录 |

##### **提现地址** - `/api/fams/user/wallet-address`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/fams/user/wallet-address` | 添加提现地址 |
| `GET` | `/api/fams/user/wallet-address` | 提现地址列表 |
| `GET` | `/api/fams/user/wallet-address/:id` | 地址详情 |
| `PATCH` | `/api/fams/user/wallet-address/:id/audit` | 审核地址 |
| `DELETE` | `/api/fams/user/wallet-address/:id` | 删除地址 |

##### **钱包提现** - `/api/fams/user/wallet-withdrawal`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/fams/user/wallet-withdrawal` | 申请提现 |
| `GET` | `/api/fams/user/wallet-withdrawal` | 提现记录列表 |
| `GET` | `/api/fams/user/wallet-withdrawal/:id` | 提现详情 |
| `PATCH` | `/api/fams/user/wallet-withdrawal/:id/approve` | 审核通过 |
| `PATCH` | `/api/fams/user/wallet-withdrawal/:id/reject` | 审核拒绝 |
| `POST` | `/api/fams/user/wallet-withdrawal/batch-approve` | 批量审核通过 |
| `POST` | `/api/fams/user/wallet-withdrawal/batch-reject` | 批量审核拒绝 |
| `GET` | `/api/fams/user/wallet-withdrawal/stats` | 提现统计 |

##### **钱包存款** - `/api/fams/user/wallet-deposit`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/fams/user/wallet-deposit` | 申请存款 |
| `GET` | `/api/fams/user/wallet-deposit` | 存款记录列表 |
| `GET` | `/api/fams/user/wallet-deposit/:id` | 存款详情 |

##### **用户设置** - `/api/fams/user/setting`

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/api/fams/user/:id/setting` | 获取用户设置 |
| `PATCH` | `/api/fams/user/:id/setting/fee` | 更新手续费设置 |

#### 3.3.2 银行账户（bank）

##### **银行客户** - `/api/fams/bank/customer`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/fams/bank/customer` | 创建银行客户 |
| `GET` | `/api/fams/bank/customer` | 客户列表 |
| `GET` | `/api/fams/bank/customer/:id` | 客户详情 |
| `PUT` | `/api/fams/bank/customer/:id` | 更新客户信息 |
| `DELETE` | `/api/fams/bank/customer/:id` | 删除客户 |
| `PATCH` | `/api/fams/bank/customer/:id/kyc-status` | 更新KYC状态 |

##### **银行账户** - `/api/fams/bank/account`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/fams/bank/account` | 创建账户 |
| `GET` | `/api/fams/bank/account` | 账户列表 |
| `GET` | `/api/fams/bank/account/:id` | 账户详情 |
| `PUT` | `/api/fams/bank/account/:id` | 更新账户 |
| `DELETE` | `/api/fams/bank/account/:id` | 删除账户 |
| `PATCH` | `/api/fams/bank/account/:id/status` | 更新账户状态 |
| `GET` | `/api/fams/bank/account/:id/balance` | 账户余额 |
| `GET` | `/api/fams/bank/account/:id/transaction` | 账户流水 |
| `GET` | `/api/fams/bank/account/stats` | 账户统计 |

##### **开户申请** - `/api/fams/bank/account-application`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/fams/bank/account-application` | 提交开户申请 |
| `GET` | `/api/fams/bank/account-application` | 申请列表 |
| `GET` | `/api/fams/bank/account-application/:id` | 申请详情 |
| `PATCH` | `/api/fams/bank/account-application/:id/approve` | 审核通过 |
| `PATCH` | `/api/fams/bank/account-application/:id/reject` | 审核拒绝 |

##### **银行存款** - `/api/fams/bank/deposit`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/fams/bank/deposit` | 记录存款 |
| `GET` | `/api/fams/bank/deposit` | 存款列表 |
| `GET` | `/api/fams/bank/deposit/:id` | 存款详情 |
| `PATCH` | `/api/fams/bank/deposit/:id/approve` | 审核通过 |
| `PATCH` | `/api/fams/bank/deposit/:id/reject` | 审核拒绝 |

##### **银行提现** - `/api/fams/bank/withdrawal`

| 方法 | 路径 | 说明 |
|------|------|------|
| `POST` | `/api/fams/bank/withdrawal` | 申请提现 |
| `GET` | `/api/fams/bank/withdrawal` | 提现列表 |
| `GET` | `/api/fams/bank/withdrawal/:id` | 提现详情 |
| `PATCH` | `/api/fams/bank/withdrawal/:id/approve` | 审核通过 |
| `PATCH` | `/api/fams/bank/withdrawal/:id/reject` | 审核拒绝 |

##### **Webhook 记录** - `/api/fams/bank/webhook-record`

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/api/fams/bank/webhook-record` | Webhook 记录列表 |
| `GET` | `/api/fams/bank/webhook-record/:id` | 记录详情 |

#### 3.3.3 资金明细

##### **资金明细** - `/api/fams/fund/detail`

| 方法 | 路径 | 说明 |
|------|------|------|
| `GET` | `/api/fams/fund/detail` | 资金明细列表 |
| `GET` | `/api/fams/fund/detail/:id` | 明细详情 |
| `GET` | `/api/fams/fund/detail/stats` | 资金统计 |

---

## 四、认证机制

### 4.1 全局 Cookie 认证

**认证方式**：所有接口（除公开接口外）均通过全局中间件验证 Cookie。

**Cookie 字段**：

| 字段 | 类型 | 说明 | 示例 |
|------|------|------|------|
| `user_id` | `int64` | 用户ID（雪花算法生成） | `1234567890123456789` |
| `session_token` | `string` | 会话令牌（SHA-256哈希值） | `e3b0c44298fc1c...` (64字符) |

**Cookie 属性**：
- `HttpOnly`: true （防止 XSS 攻击）
- `Secure`: true （生产环境使用 HTTPS）
- `SameSite`: Strict （防止 CSRF 攻击）
- `Max-Age`: 86400 （24小时）

**前端处理**：
- ✅ 浏览器自动携带 Cookie，无需手动处理
- ✅ 无需存储和管理 access_token / refresh_token
- ✅ 登录成功后，后端自动设置 Cookie
- ✅ 登出时，后端清除 Cookie

### 4.2 公开接口清单

**只有以下路径无需认证**（其他所有接口都需要 Cookie）：

```
/api/public/admin/captcha       # 管理员验证码
/api/public/admin/login         # 管理员登录
/api/public/iam/login           # 用户登录
/api/public/iam/register        # 用户注册
/api/public/iam/email/send      # 发送邮箱验证码
/api/public/iam/email/verify    # 验证邮箱
```

### 4.3 用户 ID 生成（Snowflake 雪花算法）

**ID 结构**：64位整数

```
┌─────────┬──────────────────────────────┬────────┬────────┬──────────────┐
│ 符号位  │      时间戳 (41位)            │ 数据中心│ 机器ID │  序列号(12位) │
│  (1位)  │  (毫秒级时间戳)                │ (5位)  │ (5位)  │              │
└─────────┴──────────────────────────────┴────────┴────────┴──────────────┘
     0           1704067200000              0-31    0-31     0-4095
```

**优势**：
- ✅ 全局唯一，分布式环境无冲突
- ✅ 趋势递增，有利于数据库索引
- ✅ 高性能，本地生成，每毫秒最多 4096 个 ID
- ✅ 可从 ID 中解析出创建时间

**配置示例**：
```go
epoch := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli()
dataCenterID := 1  // 0-31
machineID := 1     // 0-31
```

### 4.4 Session 安全机制

**Token 生成与存储**：

1. **生成随机 Token**（登录时）
   ```go
   token := uuid.New().String() // 或 crypto/rand 生成32字节
   ```

2. **计算 SHA-256 哈希**
   ```go
   hash := sha256.Sum256([]byte(token))
   hashedToken := hex.EncodeToString(hash[:])
   ```

3. **存储到数据库**
   ```sql
   INSERT INTO admin_sys_staff_sessions (staff_id, session_token, expires_at)
   VALUES (?, ?, NOW() + INTERVAL 24 HOUR)
   ```

4. **设置到 Cookie**
   ```go
   http.SetCookie(w, &http.Cookie{
       Name:     "user_id",
       Value:    fmt.Sprintf("%d", staffId),
       HttpOnly: true,
       Secure:   true,
       SameSite: http.SameSiteStrictMode,
   })
   http.SetCookie(w, &http.Cookie{
       Name:     "session_token",
       Value:    hashedToken, // 注意：Cookie中也存哈希值
       MaxAge:   86400,
       HttpOnly: true,
       Secure:   true,
       SameSite: http.SameSiteStrictMode,
   })
   ```

**验证流程**：

```go
// 1. 从 Cookie 读取
userIdStr := r.Cookie("user_id")
sessionToken := r.Cookie("session_token")

// 2. 查询数据库
session := db.Where("staff_id = ? AND session_token = ? AND expires_at > ?",
    userId, sessionToken, time.Now()).First(&session)

// 3. 验证成功 → 更新最后活跃时间
db.Model(&session).Update("last_active_at", time.Now())

// 4. 验证失败 → 返回 401
```

**防暴库攻击**：
- ❌ 数据库被攻破，攻击者只能看到哈希值
- ❌ 无法从哈希值反推原始 Token
- ❌ 无法使用哈希值模拟会话（因为 Cookie 中存的也是哈希）
- ✅ 只有持有原始 Token 才能通过验证（但原始Token只在登录时返回一次，不存储）

实际上，按照上面的设计，原始Token也没有返回给客户端，Cookie中直接存的就是哈希值。这样即使中间人截获了Cookie，也无法用这个哈希值去生成新的会话。

---

## 五、数据格式规范

### 5.1 请求格式

#### 5.1.1 请求头

**标准请求头**：

```http
Content-Type: application/json
```

**说明**：
- ❌ 不使用 `Authorization` 头（改用 Cookie）
- ❌ 不使用 `X-Timestamp`、`X-Sign`（取消签名验证）
- ✅ Cookie 由浏览器自动携带

#### 5.1.2 请求体

**JSON 格式**：

```json
{
  "field1": "value1",
  "field2": 123,
  "field3": true
}
```

**列表查询参数**（Query String）：

```
GET /api/admin/sys/staff?page=1&pageSize=10&keyword=admin&status=1
```

### 5.2 响应格式

#### 5.2.1 统一响应结构

**成功响应**：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    // 业务数据
  }
}
```

**失败响应**：

```json
{
  "code": 400,
  "message": "参数错误",
  "data": null
}
```

#### 5.2.2 分页响应

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 100,
    "list": [
      { "id": 1, "name": "..." },
      { "id": 2, "name": "..." }
    ]
  }
}
```

#### 5.2.3 响应码规范

| HTTP 状态码 | 业务码 (code) | 说明 | message 示例 |
|------------|--------------|------|-------------|
| `200` | `0` | 成功 | `"success"` |
| `400` | `400` | 请求参数错误 | `"参数错误"`, `"邮箱格式不正确"` |
| `401` | `401` | 未登录或会话过期 | `"未登录或会话已过期"` |
| `403` | `403` | 无权限访问 | `"权限不足"` |
| `404` | `404` | 资源不存在 | `"数据不存在"` |
| `409` | `409` | 资源冲突 | `"邮箱已存在"`, `"角色代码重复"` |
| `500` | `500` | 服务器错误 | `"服务器内部错误"` |

### 5.3 时间格式规范

**统一使用毫秒级时间戳 int64**：

```json
{
  "createdAt": 1704067200000,     // 创建时间（毫秒时间戳）
  "updatedAt": 1704153600000,     // 更新时间（毫秒时间戳）
  "expiresAt": 1704240000000      // 过期时间（毫秒时间戳）
}
```

**注意事项**：
- ✅ 数据库中保留 `gorm.Model`（包含 `CreatedAt`, `UpdatedAt`, `DeletedAt`）
- ✅ API 响应时，将 `time.Time` 转换为毫秒时间戳 `int64`
- ✅ 前端统一使用 `new Date(timestamp)` 处理时间

**转换示例**：

```go
// Go: time.Time → 毫秒时间戳
timestamp := time.Now().UnixMilli()

// Go: 毫秒时间戳 → time.Time
t := time.UnixMilli(timestamp)

// JavaScript: 时间戳 → Date
const date = new Date(timestamp)

// JavaScript: Date → 时间戳
const timestamp = Date.now()
```

### 5.4 字段命名规范

**JSON 字段命名**：驼峰命名法（camelCase）

```json
{
  "userId": 123,
  "userName": "admin",
  "createdAt": 1704067200000,
  "isActive": true
}
```

**数据库字段命名**：蛇形命名法（snake_case）

```sql
user_id, user_name, created_at, is_active
```

---

## 六、错误处理

### 6.1 错误响应示例

#### 参数错误（400）

```json
{
  "code": 400,
  "message": "邮箱格式不正确",
  "data": null
}
```

#### 未认证（401）

```json
{
  "code": 401,
  "message": "未登录或会话已过期",
  "data": null
}
```

#### 权限不足（403）

```json
{
  "code": 403,
  "message": "权限不足，无法访问此资源",
  "data": null
}
```

#### 资源不存在（404）

```json
{
  "code": 404,
  "message": "员工不存在",
  "data": null
}
```

#### 资源冲突（409）

```json
{
  "code": 409,
  "message": "邮箱已被使用",
  "data": null
}
```

#### 服务器错误（500）

```json
{
  "code": 500,
  "message": "服务器内部错误",
  "data": null
}
```

### 6.2 字段验证错误

**多字段错误**：

```json
{
  "code": 400,
  "message": "参数验证失败",
  "data": {
    "errors": {
      "email": "邮箱格式不正确",
      "password": "密码长度必须在8-20位之间"
    }
  }
}
```

---

## 七、分页规范

### 7.1 请求参数

```
GET /api/admin/sys/staff?page=1&pageSize=10
```

| 参数 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `page` | `int` | `1` | 页码（从1开始） |
| `pageSize` | `int` | `10` | 每页数量（最大100） |

### 7.2 响应格式

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 156,              // 总记录数
    "list": [                  // 当前页数据
      { "id": 1, "name": "..." },
      { "id": 2, "name": "..." }
    ]
  }
}
```

---

## 八、筛选与排序

### 8.1 筛选参数

```
GET /api/admin/sys/staff?keyword=admin&status=1&roleId=2
```

**常用筛选参数**：

| 参数 | 说明 | 示例 |
|------|------|------|
| `keyword` | 关键词搜索 | `keyword=admin` |
| `status` | 状态筛选 | `status=1` |
| `startTime` | 开始时间（毫秒时间戳） | `startTime=1704067200000` |
| `endTime` | 结束时间（毫秒时间戳） | `endTime=1704153600000` |

### 8.2 排序参数

```
GET /api/admin/sys/staff?sortBy=createdAt&sortOrder=desc
```

| 参数 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| `sortBy` | `string` | `createdAt` | 排序字段 |
| `sortOrder` | `string` | `desc` | 排序方向：`asc`, `desc` |

---

## 九、API 文档示例

### 9.1 创建员工

**接口**：`POST /api/admin/sys/staff`

**请求头**：
```http
Content-Type: application/json
Cookie: user_id=1234567890123456789; session_token=e3b0c44...
```

**请求体**：
```json
{
  "name": "张三",
  "email": "zhangsan@example.com",
  "password": "password123",
  "roleIds": [1, 2]
}
```

**成功响应**（200）：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1234567890123456789,
    "name": "张三",
    "email": "zhangsan@example.com",
    "status": 1,
    "statusText": "启用",
    "roles": [
      { "id": 1, "name": "管理员", "code": "admin" },
      { "id": 2, "name": "审核员", "code": "auditor" }
    ],
    "createdAt": 1704067200000,
    "updatedAt": 1704067200000
  }
}
```

**失败响应**（409）：
```json
{
  "code": 409,
  "message": "邮箱已被使用",
  "data": null
}
```

### 9.2 设置员工角色

**接口**：`PATCH /api/admin/sys/staff/:id/role`

**请求头**：
```http
Content-Type: application/json
Cookie: user_id=1234567890123456789; session_token=e3b0c44...
```

**请求体**：
```json
{
  "roleIds": [1, 3, 5]
}
```

**成功响应**（200）：
```json
{
  "code": 0,
  "message": "success",
  "data": null
}
```

---

## 十、最佳实践

### 10.1 前端集成

**登录示例**：

```javascript
// 登录（浏览器自动保存 Cookie）
const response = await fetch('/api/public/admin/login', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  credentials: 'include', // 重要：允许跨域携带 Cookie
  body: JSON.stringify({
    email: 'admin@example.com',
    password: 'password123',
    captchaId: 'xxx',
    captcha: '1234'
  })
})

const result = await response.json()
if (result.code === 0) {
  // 登录成功，Cookie 已自动设置
  console.log('登录成功', result.data)
}
```

**调用认证接口**：

```javascript
// 获取员工列表（Cookie 自动携带）
const response = await fetch('/api/admin/sys/staff?page=1&pageSize=10', {
  credentials: 'include' // 重要：自动携带 Cookie
})

const result = await response.json()
if (result.code === 401) {
  // 未登录或会话过期，跳转登录页
  window.location.href = '/login'
}
```

**登出示例**：

```javascript
// 登出（后端清除 Cookie）
const response = await fetch('/api/admin/auth/logout', {
  method: 'POST',
  credentials: 'include'
})

const result = await response.json()
if (result.code === 0) {
  // 登出成功，跳转登录页
  window.location.href = '/login'
}
```

### 10.2 后端中间件

**全局 Cookie 认证中间件**：

```go
func CookieAuthMiddleware(publicPaths map[string]bool) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // 1. 检查是否为公开路径
            if publicPaths[r.URL.Path] || strings.HasPrefix(r.URL.Path, "/api/public/") {
                next.ServeHTTP(w, r)
                return
            }

            // 2. 获取 Cookie
            userIdCookie, err := r.Cookie("user_id")
            if err != nil {
                unauthorized(w, "未登录或会话已过期")
                return
            }

            sessionTokenCookie, err := r.Cookie("session_token")
            if err != nil {
                unauthorized(w, "未登录或会话已过期")
                return
            }

            userId, _ := strconv.ParseInt(userIdCookie.Value, 10, 64)
            sessionToken := sessionTokenCookie.Value

            // 3. 验证 session
            var session Session
            err = db.Where("user_id = ? AND session_token = ? AND expires_at > ?",
                userId, sessionToken, time.Now()).First(&session).Error
            if err != nil {
                unauthorized(w, "会话不存在或已失效")
                return
            }

            // 4. 更新最后活跃时间
            db.Model(&session).Update("last_active_at", time.Now())

            // 5. 将用户ID存入context
            ctx := context.WithValue(r.Context(), "userId", userId)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

func unauthorized(w http.ResponseWriter, message string) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusUnauthorized)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "code":    401,
        "message": message,
        "data":    nil,
    })
}
```

### 10.3 安全建议

**生产环境配置**：

- ✅ 使用 HTTPS（启用 `Secure` Cookie 属性）
- ✅ 设置 `HttpOnly` 防止 XSS
- ✅ 设置 `SameSite=Strict` 防止 CSRF
- ✅ 定期清理过期 session
- ✅ 限制登录失败次数（防暴力破解）
- ✅ 记录所有登录、登出、敏感操作日志
- ✅ 使用雪花算法生成不可预测的用户ID

---

## 十一、版本历史

| 版本 | 日期 | 变更内容 |
|------|------|---------|
| v1.0 | 2025-10-22 | 初始版本，定义 API 规范 |

---

## 十二、附录

### 12.1 完整表清单（35张表）

#### Admin RPC（8张表）
```
admin_sys_staffs
admin_sys_roles
admin_sys_menus
admin_sys_permissions
admin_sys_staff_roles
admin_sys_role_menus
admin_sys_staff_logs
admin_data_countries
```

#### IAM RPC（12张表）
```
iam_users
iam_user_credentials
iam_user_credential_logs
iam_user_profiles
iam_user_emails
iam_user_email_logs
iam_user_sessions
iam_user_session_logs
iam_user_invites
iam_user_invite_logs
iam_user_country_auths
iam_email_verifies
```

#### FAMS RPC（15张表）
```
fams_user_settings
fams_user_wallets
fams_user_wallet_ledgers
fams_user_wallet_withdrawals
fams_user_wallet_addresses
fams_user_wallet_address_audits
fams_user_wallet_address_logs
fams_user_wallet_deposits
fams_bank_customers
fams_bank_accounts
fams_bank_deposits
fams_bank_deposit_audits
fams_bank_account_applications
fams_bank_withdrawals
fams_bank_webhook_records
fams_fund_details
```

### 12.2 参考文档

- [全局中间件配置指南](./global-middleware-guide.md)
- [认证机制指南](./authentication-guide.md)
- [Go-Zero 官方文档](https://go-zero.dev/)
- [RESTful API 设计指南](https://restfulapi.net/)

---

**文档结束**
