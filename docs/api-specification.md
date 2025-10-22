# AKATM API 接口规范

## 概述

本文档定义了 AKATM 项目的 RESTful API 接口规范，用于统一所有 API 接口的命名和结构。

## 核心原则

### 1. 认证与会话管理

#### 1.1 用户ID生成策略

**雪花算法 (Snowflake)**

系统使用雪花算法生成全局唯一的用户ID，具有以下优势：
- **全局唯一**: 分布式环境下保证ID唯一性
- **趋势递增**: 有利于数据库索引性能
- **高性能**: 本地生成，无需网络请求
- **信息量大**: 64位整数包含时间戳、机器ID、序列号

**ID结构**（64位）：
```
0 - 00000000 00000000 00000000 00000000 00000000 0 - 00000 - 00000 - 000000000000
|   |                                                |       |       |
|   时间戳(41位)                                      数据中心 机器ID  序列号(12位)
|                                                    ID(5位)  (5位)
符号位(1位，始终为0)
```

**配置建议**：
- 时间戳起始点：`2024-01-01 00:00:00`
- 数据中心ID：0-31（5位）
- 机器ID：0-31（5位）
- 序列号：0-4095（12位，每毫秒最多4096个ID）

#### 1.2 会话管理方案

**Cookie存储策略**

采用 Cookie 存储会话信息，而非 Authorization Header，具有以下优势：
- **自动携带**: 浏览器自动处理，前端无需手动设置
- **HttpOnly保护**: 防止XSS攻击窃取会话
- **同站保护**: 配合 SameSite 防止CSRF攻击
- **安全性高**: Session Token 使用 SHA-256 哈希，即使数据库泄露也无法模拟会话

**Cookie字段**：

| 字段 | 说明 | 示例值 |
|------|------|--------|
| `user_id` | 雪花算法生成的用户ID | `1234567890123456789` |
| `session_token` | SHA-256哈希的会话令牌 | `a3f5...d8c2` (64字符) |

**Cookie属性配置**：

```http
Set-Cookie: user_id=1234567890123456789; Path=/; HttpOnly; Secure; SameSite=Strict
Set-Cookie: session_token=a3f5d8c2...; Path=/; HttpOnly; Secure; SameSite=Strict; Max-Age=86400
```

| 属性 | 值 | 说明 |
|------|------|------|
| `Path` | `/` | 全站可用 |
| `HttpOnly` | `true` | **防XSS**: JavaScript无法读取 |
| `Secure` | `true` | **仅HTTPS**: 生产环境必须启用 |
| `SameSite` | `Strict` | **防CSRF**: 仅同站请求携带 |
| `Max-Age` | `86400` | 会话有效期（秒），24小时 |
| `Domain` | 不设置 | 默认当前域名 |

**安全机制**：

1. **Session Token 生成**
   ```
   原始Token = UUID v4 或 随机字符串(32字节)
   存储到Cookie = SHA-256(原始Token)
   存储到数据库 = SHA-256(原始Token)
   ```

2. **会话验证流程**
   ```
   1. 从Cookie读取 user_id 和 session_token
   2. 查询数据库: SELECT * FROM iam_user_session
      WHERE user_id = ? AND token = ? AND expires_at > NOW()
   3. 验证通过 → 更新 last_active_at
   4. 验证失败 → 返回 401 Unauthorized
   ```

3. **防暴库攻击**
   - 数据库仅存储 SHA-256 哈希值
   - 即使数据库泄露，攻击者也无法反推原始Token
   - 无法使用泄露的哈希值模拟会话

4. **会话刷新机制**
   ```
   - 每次请求验证会话有效性
   - 距离过期时间不足1小时时，自动延长会话
   - 用户主动登出时，立即删除会话记录
   ```

#### 1.3 认证流程

**后台员工登录 (Admin)**

```
POST /api/admin/auth/login
Content-Type: application/json

{
  "email": "admin@example.com",
  "password": "password123"
}

Response:
Set-Cookie: user_id=123456789; HttpOnly; Secure; SameSite=Strict
Set-Cookie: session_token=a3f5d8c2...; HttpOnly; Secure; SameSite=Strict; Max-Age=86400

{
  "code": 0,
  "message": "登录成功",
  "data": {
    "userId": 123456789,
    "name": "管理员",
    "email": "admin@example.com",
    "roles": ["admin", "auditor"]
  }
}
```

**前台用户登录 (IAM)**

```
POST /api/iam/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "userType": "manager"
}

Response:
Set-Cookie: user_id=987654321; HttpOnly; Secure; SameSite=Strict
Set-Cookie: session_token=b7e9f3a1...; HttpOnly; Secure; SameSite=Strict; Max-Age=86400

{
  "code": 0,
  "message": "登录成功",
  "data": {
    "userId": 987654321,
    "userType": "manager",
    "email": "user@example.com",
    "parentId": 123456,
    "exchangeFeeRate": "0.0500"
  }
}
```

**登出**

```
POST /api/admin/auth/logout
Cookie: user_id=123456789; session_token=a3f5d8c2...

Response:
Set-Cookie: user_id=; Max-Age=0
Set-Cookie: session_token=; Max-Age=0

{
  "code": 0,
  "message": "登出成功"
}
```

#### 1.4 请求认证

**同域名架构**

前端页面和API部署在同一域名下，浏览器自动携带Cookie：

```
前端页面: https://app.example.com
API接口:  https://app.example.com/api/*

或者使用子域名:
前端页面: https://www.example.com
API接口:  https://api.example.com  (需要设置Cookie Domain=.example.com)
```

**请求示例**

```http
GET /api/admin/staff HTTP/1.1
Host: app.example.com
Cookie: user_id=123456789; session_token=a3f5d8c2...
X-Timestamp: 1640000000
X-Sign: e7b8c9d0...
```

**无需前端处理**：
- ✅ 浏览器自动携带Cookie
- ✅ 无需手动设置 Authorization Header
- ✅ 无需前端存储Token
- ✅ 自动处理会话过期（401跳转登录）

#### 1.5 安全最佳实践

**生产环境配置清单**：

- [x] **HTTPS强制**: 所有接口必须使用HTTPS
- [x] **HttpOnly Cookie**: 防止XSS窃取Session
- [x] **Secure Cookie**: 仅通过HTTPS传输
- [x] **SameSite=Strict**: 防止CSRF攻击
- [x] **SHA-256哈希**: 数据库存储哈希值，防暴库
- [x] **会话过期**: 24小时自动过期
- [x] **并发登录限制**: 同一用户最多N个活跃会话
- [x] **异常IP检测**: 会话绑定IP，异地登录需二次验证
- [x] **登录失败锁定**: 5次失败锁定账户15分钟
- [x] **操作日志**: 记录所有敏感操作

**数据库表设计**：

```sql
-- Admin Staff Session
CREATE TABLE admin_staff_session (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    staff_id BIGINT NOT NULL,
    token CHAR(64) NOT NULL COMMENT 'SHA-256哈希值',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    last_active_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ip_address VARCHAR(45),
    user_agent VARCHAR(255),
    INDEX idx_staff_token (staff_id, token),
    INDEX idx_expires (expires_at)
);

-- IAM User Session
CREATE TABLE iam_user_session (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL COMMENT '雪花算法生成',
    token CHAR(64) NOT NULL COMMENT 'SHA-256哈希值',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    last_active_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ip_address VARCHAR(45),
    user_agent VARCHAR(255),
    INDEX idx_user_token (user_id, token),
    INDEX idx_expires (expires_at)
);
```

### 2. URL 结构规范

所有接口遵循以下 URL 结构：

```
/api/{rpc}/{resource}[/:id][/{action}]
```

**参数说明**：
- `{rpc}`: RPC 服务名，对应数据库表前缀（admin、iam、fams）
- `{resource}`: 资源名称，对应表名去掉 RPC 前缀后的部分
- `:id`: 资源 ID（可选）
- `{action}`: 子资源或特定操作（可选）

**示例**：
- `/api/admin/staff` - 员工列表
- `/api/admin/staff/:id` - 员工详情
- `/api/admin/staff/:id/role` - 员工角色设置
- `/api/iam/user/:id/emails` - 用户邮箱列表
- `/api/fams/customer/:id/accounts` - 客户账户列表

### 3. HTTP 方法规范

| 方法 | 用途 | 示例 |
|------|------|------|
| GET | 查询资源（列表/详情） | `GET /api/admin/staff` |
| POST | 创建资源 | `POST /api/admin/staff` |
| PUT | 完整更新资源 | `PUT /api/admin/staff/:id` |
| PATCH | 部分更新资源/特定操作 | `PATCH /api/admin/staff/:id/role` |
| DELETE | 删除资源 | `DELETE /api/admin/staff/:id` |

**PATCH 使用场景**：
- 更新资源的特定字段：`PATCH /api/admin/staff/:id/status`
- 设置关联关系：`PATCH /api/admin/staff/:id/role`
- 执行特定操作：`PATCH /api/iam/user/:id/fees`

### 4. RPC 服务分类

#### admin - 后台管理服务
管理后台员工、角色、菜单、权限等

**表前缀**: `admin_`

**资源列表**：
- `staff` - 员工管理
- `role` - 角色管理
- `menu` - 菜单管理
- `permission` - 权限管理
- `country` - 国家管理
- `staff-log` - 员工日志

#### iam - 用户身份认证服务
管理前台用户（总代、代理、客户经理、客户）的身份认证

**表前缀**: `iam_`

**资源列表**：
- `user` - 前台用户
- `user-credential` - 用户凭证
- `user-profile` - 用户资料
- `user-email` - 用户邮箱
- `user-session` - 用户会话
- `user-invite` - 用户邀请
- `user-country-auth` - 用户国家授权
- `country` - 国家信息
- `email-verify` - 邮箱验证

#### fams - 财务账户管理服务
管理用户钱包、银行账户、存取款、审核等

**表前缀**: `fams_`

**资源列表**：
- `user-setting` - 用户设置
- `wallet` - 用户钱包
- `wallet-ledger` - 钱包账变
- `wallet-withdrawal` - 钱包提现
- `wallet-address` - 钱包地址
- `wallet-address-audit` - 钱包地址审核
- `wallet-deposit` - 钱包存款
- `customer` - 银行客户
- `account` - 银行账户
- `deposit` - 银行存款
- `deposit-audit` - 存款审核
- `withdrawal` - 银行提现
- `account-application` - 开户申请
- `webhook-record` - 回调记录
- `fund-detail` - 资金明细

### 5. 资源命名规则

#### 表名到资源名的转换

1. **去掉 RPC 前缀**
   - `admin_staff` → `staff`
   - `iam_user` → `user`
   - `fams_bank_customer` → `customer`

2. **使用 kebab-case**
   - `user_wallet` → `user-wallet`
   - `staff_log` → `staff-log`
   - `email_verify` → `email-verify`

3. **简化常见名称**
   - `fams_bank_customer` → `customer` (而非 `bank-customer`)
   - `fams_bank_account` → `account` (而非 `bank-account`)
   - `fams_user_wallet` → `wallet` (而非 `user-wallet`)

4. **保留语义明确的复合名称**
   - `wallet-ledger` - 钱包账变
   - `wallet-address` - 钱包地址
   - `account-application` - 开户申请
   - `fund-detail` - 资金明细

## 标准接口模式

### 模式 1: CRUD 操作

```
# 列表查询
GET /api/{rpc}/{resource}
Query: page, pageSize, keyword, filters...

# 创建
POST /api/{rpc}/{resource}
Body: {...}

# 详情查询
GET /api/{rpc}/{resource}/:id

# 完整更新
PUT /api/{rpc}/{resource}/:id
Body: {...}

# 删除（软删除）
DELETE /api/{rpc}/{resource}/:id
```

**示例**：
```
GET    /api/admin/staff              # 员工列表
POST   /api/admin/staff              # 创建员工
GET    /api/admin/staff/:id          # 员工详情
PUT    /api/admin/staff/:id          # 更新员工
DELETE /api/admin/staff/:id          # 删除员工
```

### 模式 2: 状态更新

```
PATCH /api/{rpc}/{resource}/:id/status
Body: { "status": 1 }
```

**示例**：
```
PATCH /api/admin/staff/:id/status    # 更新员工状态
PATCH /api/iam/user/:id/status       # 更新用户状态
PATCH /api/fams/account/:id/status   # 更新账户状态
```

### 模式 3: 关联关系管理

```
# 设置关联
PATCH /api/{rpc}/{resource}/:id/{relation}
Body: { "relationIds": [1, 2, 3] }

# 查询关联
GET /api/{rpc}/{resource}/:id/{relation}

# 添加关联项
POST /api/{rpc}/{resource}/:id/{relation}
Body: {...}

# 删除关联项
DELETE /api/{rpc}/{resource}/:id/{relation}/:relationId
```

**示例**：
```
# 员工角色管理
PATCH  /api/admin/staff/:id/role         # 设置员工角色
GET    /api/admin/staff/:id/menus        # 获取员工菜单

# 角色菜单管理
PATCH  /api/admin/role/:id/menus         # 分配角色菜单
GET    /api/admin/role/:id/menus         # 获取角色菜单

# 用户邮箱管理
GET    /api/iam/user/:id/emails          # 获取用户邮箱列表
POST   /api/iam/user/:id/emails          # 添加用户邮箱
DELETE /api/iam/user/:id/emails/:emailId # 删除用户邮箱

# 用户国家授权
PATCH  /api/iam/user/:id/countries       # 设置用户授权国家
```

### 模式 4: 特定字段更新

```
PATCH /api/{rpc}/{resource}/:id/{field}
Body: { "field": "value", ... }
```

**示例**：
```
PATCH /api/admin/staff/:id/password      # 重设员工密码
PATCH /api/iam/user/:id/fees             # 更新用户手续费
PATCH /api/iam/user/:id/profile          # 更新用户资料
```

### 模式 5: 子资源查询

```
GET /api/{rpc}/{resource}/:id/{sub-resource}
Query: page, pageSize, filters...
```

**示例**：
```
GET /api/fams/customer/:id/accounts           # 客户账户列表
GET /api/fams/account/:id/transactions        # 账户流水
GET /api/iam/user/:id/hierarchy               # 用户层级关系
GET /api/iam/user/:id/sub-users               # 下级用户列表
```

### 模式 6: 审核操作

```
POST /api/{rpc}/{resource}/:id/approve        # 审核通过
POST /api/{rpc}/{resource}/:id/reject         # 审核拒绝
POST /api/{rpc}/{resource}/batch-approve      # 批量审核通过
POST /api/{rpc}/{resource}/batch-reject       # 批量审核拒绝
```

**示例**：
```
POST /api/fams/deposit/:id/approve            # 存款审核通过
POST /api/fams/deposit/:id/reject             # 存款审核拒绝
POST /api/fams/withdrawal/:id/approve         # 提现审核通过
POST /api/fams/withdrawal/:id/reject          # 提现审核拒绝
POST /api/fams/withdrawal/batch-approve       # 批量审核提现
```

### 模式 7: 统计查询

```
GET /api/{rpc}/{resource}/stats
Query: startTime, endTime, filters...
```

**示例**：
```
GET /api/fams/account/stats                   # 账户统计
GET /api/fams/withdrawal/stats                # 提现统计
```

## 完整接口示例

### Admin RPC 接口

#### 员工管理 (staff)

```
# 基础 CRUD
GET    /api/admin/staff                       # 员工列表
POST   /api/admin/staff                       # 创建员工
GET    /api/admin/staff/:id                   # 员工详情
PUT    /api/admin/staff/:id                   # 更新员工
DELETE /api/admin/staff/:id                   # 删除员工

# 特定操作
PATCH  /api/admin/staff/:id/password          # 重设密码
PATCH  /api/admin/staff/:id/role              # 设置角色
PATCH  /api/admin/staff/:id/status            # 更新状态

# 关联查询
GET    /api/admin/staff/:id/menus             # 员工菜单
```

#### 角色管理 (role)

```
# 基础 CRUD
GET    /api/admin/role                        # 角色列表
POST   /api/admin/role                        # 创建角色
GET    /api/admin/role/:id                    # 角色详情
PUT    /api/admin/role/:id                    # 更新角色
DELETE /api/admin/role/:id                    # 删除角色

# 关联管理
PATCH  /api/admin/role/:id/menus              # 分配菜单
GET    /api/admin/role/:id/menus              # 获取角色菜单
```

#### 菜单管理 (menu)

```
GET    /api/admin/menu                        # 菜单列表
POST   /api/admin/menu                        # 创建菜单
GET    /api/admin/menu/:id                    # 菜单详情
PUT    /api/admin/menu/:id                    # 更新菜单
DELETE /api/admin/menu/:id                    # 删除菜单
PATCH  /api/admin/menu/:id/status             # 更新状态
```

#### 国家管理 (country)

```
GET    /api/admin/country                     # 国家列表
POST   /api/admin/country                     # 创建国家
GET    /api/admin/country/:id                 # 国家详情
PUT    /api/admin/country/:id                 # 更新国家
DELETE /api/admin/country/:id                 # 删除国家
PATCH  /api/admin/country/:id/status          # 更新状态
```

### IAM RPC 接口

#### 用户管理 (user)

```
# 基础 CRUD
GET    /api/iam/user                          # 用户列表
POST   /api/iam/user                          # 创建用户（普通用户）
POST   /api/iam/user/super-agent              # 创建总代
GET    /api/iam/user/:id                      # 用户详情
PUT    /api/iam/user/:id                      # 更新用户
DELETE /api/iam/user/:id                      # 删除用户

# 特定字段更新
PATCH  /api/iam/user/:id/status               # 更新状态
PATCH  /api/iam/user/:id/fees                 # 更新手续费
PATCH  /api/iam/user/:id/countries            # 国家授权
PATCH  /api/iam/user/:id/profile              # 更新资料

# 关联管理
GET    /api/iam/user/:id/emails               # 获取邮箱列表
POST   /api/iam/user/:id/emails               # 添加邮箱
DELETE /api/iam/user/:id/emails/:emailId      # 删除邮箱

# 层级关系
GET    /api/iam/user/:id/hierarchy            # 层级关系
GET    /api/iam/user/:id/sub-users            # 下级用户列表
```

#### 邀请管理 (user-invite)

```
GET    /api/iam/user-invite                   # 邀请列表
POST   /api/iam/user-invite                   # 生成邀请码
GET    /api/iam/user-invite/:id               # 邀请详情
DELETE /api/iam/user-invite/:id               # 删除邀请
```

### FAMS RPC 接口

#### 客户管理 (customer)

```
# 基础 CRUD
GET    /api/fams/customer                     # 客户列表
POST   /api/fams/customer                     # 创建客户
GET    /api/fams/customer/:id                 # 客户详情
PUT    /api/fams/customer/:id                 # 更新客户
DELETE /api/fams/customer/:id                 # 删除客户

# KYC 管理
PATCH  /api/fams/customer/:id/kyc-status      # 更新KYC状态
POST   /api/fams/customer/:id/kyc-audit       # KYC审核

# 关联查询
GET    /api/fams/customer/:id/accounts        # 客户账户列表

# 标签管理
GET    /api/fams/customer/:id/tags            # 获取客户标签
POST   /api/fams/customer/:id/tags            # 添加客户标签
DELETE /api/fams/customer/:id/tags            # 移除客户标签
GET    /api/fams/customer/tags                # 标签列表
POST   /api/fams/customer/tags                # 创建标签
PUT    /api/fams/customer/tags/:id            # 更新标签
DELETE /api/fams/customer/tags/:id            # 删除标签
```

#### 账户管理 (account)

```
# 基础 CRUD
GET    /api/fams/account                      # 账户列表
POST   /api/fams/account                      # 创建账户
GET    /api/fams/account/:id                  # 账户详情
PUT    /api/fams/account/:id                  # 更新账户
DELETE /api/fams/account/:id                  # 删除账户

# 状态和余额
PATCH  /api/fams/account/:id/status           # 更新状态
GET    /api/fams/account/:id/balance          # 账户余额

# 交易流水
GET    /api/fams/account/:id/transactions     # 账户流水

# 统计
GET    /api/fams/account/stats                # 账户统计
```

#### 存款审核 (deposit)

```
# 审核列表
GET    /api/fams/deposit                      # 存款列表
GET    /api/fams/deposit/:id                  # 存款详情

# 审核操作
POST   /api/fams/deposit/:id/approve          # 审核通过
POST   /api/fams/deposit/:id/reject           # 审核拒绝
```

#### 提现审核 (withdrawal)

```
# 审核列表
GET    /api/fams/withdrawal                   # 提现列表
GET    /api/fams/withdrawal/:id               # 提现详情

# 审核操作
POST   /api/fams/withdrawal/:id/approve       # 审核通过
POST   /api/fams/withdrawal/:id/reject        # 审核拒绝
POST   /api/fams/withdrawal/batch-approve     # 批量审核通过
POST   /api/fams/withdrawal/batch-reject      # 批量审核拒绝

# 统计
GET    /api/fams/withdrawal/stats             # 提现统计
```

#### 钱包管理 (wallet)

```
GET    /api/fams/wallet                       # 钱包列表
GET    /api/fams/wallet/:id                   # 钱包详情
GET    /api/fams/wallet/:id/ledger            # 钱包账变

# 钱包地址管理
GET    /api/fams/wallet/:id/addresses         # 地址列表
POST   /api/fams/wallet/:id/addresses         # 添加地址
DELETE /api/fams/wallet/:id/addresses/:addressId # 删除地址
PATCH  /api/fams/wallet/:id/addresses/:addressId/audit # 地址审核
```

#### 开户申请 (account-application)

```
GET    /api/fams/account-application          # 申请列表
POST   /api/fams/account-application          # 提交申请
GET    /api/fams/account-application/:id      # 申请详情
PUT    /api/fams/account-application/:id      # 更新申请
POST   /api/fams/account-application/:id/approve # 审核通过
POST   /api/fams/account-application/:id/reject  # 审核拒绝
```

## 迁移对照表

### 现有接口 → 新规范接口

#### Admin RPC

| 现有接口 | 新规范接口 | 说明 |
|---------|-----------|------|
| `POST /api/staff` | `POST /api/admin/staff` | 添加 RPC 前缀 |
| `GET /api/staff` | `GET /api/admin/staff` | 添加 RPC 前缀 |
| `GET /api/staff/:id` | `GET /api/admin/staff/:id` | 添加 RPC 前缀 |
| `PUT /api/staff/:id` | `PUT /api/admin/staff/:id` | 添加 RPC 前缀 |
| `DELETE /api/staff/:id` | `DELETE /api/admin/staff/:id` | 添加 RPC 前缀 |
| `PATCH /api/staff/:id/password` | `PATCH /api/admin/staff/:id/password` | 添加 RPC 前缀 |
| `PATCH /api/staff/:id/role` | `PATCH /api/admin/staff/:id/role` | ✅ 已符合规范 |
| `PATCH /api/staff/:id/status` | `PATCH /api/admin/staff/:id/status` | 添加 RPC 前缀 |
| `GET /api/staff/:id/menus` | `GET /api/admin/staff/:id/menus` | 添加 RPC 前缀 |
| `POST /api/role/roles` | `POST /api/admin/role` | 修改路径 |
| `GET /api/role/roles` | `GET /api/admin/role` | 修改路径 |
| `GET /api/role/roles/:id` | `GET /api/admin/role/:id` | 修改路径 |
| `POST /api/role/roles/:id/menus` | `PATCH /api/admin/role/:id/menus` | 改用 PATCH |

#### IAM RPC

| 现有接口 | 新规范接口 | 说明 |
|---------|-----------|------|
| `POST /api/user/users/super-agents` | `POST /api/iam/user/super-agent` | 修改路径 |
| `GET /api/user/users` | `GET /api/iam/user` | 简化路径 |
| `GET /api/user/users/:id` | `GET /api/iam/user/:id` | 简化路径 |
| `PUT /api/user/users/:id` | `PUT /api/iam/user/:id` | 简化路径 |
| `PATCH /api/user/users/:id/status` | `PATCH /api/iam/user/:id/status` | 简化路径 |
| `PATCH /api/user/users/:id/fees` | `PATCH /api/iam/user/:id/fees` | 简化路径 |
| `PATCH /api/user/users/:id/countries` | `PATCH /api/iam/user/:id/countries` | 简化路径 |
| `GET /api/user/users/:id/emails` | `GET /api/iam/user/:id/emails` | 简化路径 |
| `POST /api/user/users/:id/emails` | `POST /api/iam/user/:id/emails` | 简化路径 |

#### FAMS RPC

| 现有接口 | 新规范接口 | 说明 |
|---------|-----------|------|
| `GET /api/customer/customers` | `GET /api/fams/customer` | 修改前缀和路径 |
| `GET /api/customer/customers/:id` | `GET /api/fams/customer/:id` | 修改前缀和路径 |
| `PUT /api/customer/customers/:id` | `PUT /api/fams/customer/:id` | 修改前缀和路径 |
| `GET /admin/accounts` | `GET /api/fams/account` | 修改前缀 |
| `GET /admin/accounts/:id` | `GET /api/fams/account/:id` | 修改前缀 |
| `PATCH /admin/accounts/:id/status` | `PATCH /api/fams/account/:id/status` | 修改前缀 |
| `GET /api/deposit/audits/deposits` | `GET /api/fams/deposit` | 简化路径 |
| `POST /api/deposit/audits/deposits/:id/approve` | `POST /api/fams/deposit/:id/approve` | 简化路径 |
| `GET /api/withdrawal/audits/withdrawals` | `GET /api/fams/withdrawal` | 简化路径 |
| `POST /api/withdrawal/audits/withdrawals/:id/approve` | `POST /api/fams/withdrawal/:id/approve` | 简化路径 |

## 实施建议

### 分阶段迁移

1. **第一阶段**：新增接口使用新规范
2. **第二阶段**：现有接口添加新路径支持（保留旧路径兼容）
3. **第三阶段**：前端迁移到新路径
4. **第四阶段**：移除旧路径支持

### 向后兼容方案

可以在路由层面同时支持新旧两种路径：

```go
// 旧路径（兼容）
router.GET("/api/staff", handler)
// 新路径（推荐）
router.GET("/api/admin/staff", handler)
```

### 文档更新

- 所有新接口文档必须使用新规范
- 现有接口文档标注"已废弃"并指向新路径
- API 文档工具自动生成时使用新路径

## 附录

### 常用缩写

- CRUD: Create, Read, Update, Delete
- RPC: Remote Procedure Call
- IAM: Identity and Access Management
- FAMS: Financial Account Management System
- KYC: Know Your Customer

### 参考资源

- RESTful API 设计指南
- HTTP 方法语义
- URL 设计最佳实践
