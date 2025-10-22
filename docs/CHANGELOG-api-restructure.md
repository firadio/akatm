# API 结构重构变更日志

## 变更日期
2025-10-22

## 变更原因

根据新的API规范要求，URL格式需要增加**对象分类**层级，以更好地组织和管理API接口。

**旧格式**：`/api/{rpc}/{resource}[/:id][/{action}]`

**新格式**：`/api/{rpc}/{category}/{resource}[/:id][/{action}]`

其中 `{category}` 对应 `table/{category}.go` 文件名，例如：
- `sys.go` → `/api/admin/sys/`
- `data.go` → `/api/admin/data/`
- `user.go` → `/api/iam/user/`
- `bank.go` → `/api/fams/bank/`

## 核心变更

### 1. URL 格式调整

#### Admin RPC 变更

| 资源 | 旧 URL | 新 URL | 分类 |
|------|--------|--------|------|
| 员工管理 | `/api/admin/staff` | `/api/admin/sys/staff` | sys.go |
| 角色管理 | `/api/admin/role` | `/api/admin/sys/role` | sys.go |
| 菜单管理 | `/api/admin/menu` | `/api/admin/sys/menu` | sys.go |
| 国家管理 | `/api/admin/country` | `/api/admin/data/country` | data.go |

#### Admin 认证接口变更

| 接口类型 | 旧 URL | 新 URL | 说明 |
|---------|--------|--------|------|
| 获取验证码 | `/api/admin/auth/captcha` | `/api/public/admin/captcha` | 公开接口 |
| 管理员登录 | `/api/admin/auth/login` | `/api/public/admin/login` | 公开接口 |
| 管理员登出 | `/api/admin/auth/logout` | `/api/admin/auth/logout` | 需要认证 |
| 获取当前用户信息 | `/api/admin/auth/info` | `/api/admin/auth/info` | 需要认证 |
| 修改密码 | `/api/admin/auth/password` | `/api/admin/auth/password` | 需要认证 |

#### IAM RPC 变更

| 资源 | 旧 URL | 新 URL | 分类 |
|------|--------|--------|------|
| 用户管理 | `/api/iam/user` | `/api/iam/user/user` | user.go |
| 邀请管理 | `/api/iam/user-invite` | `/api/iam/user/invite` | user.go |

#### FAMS RPC 变更

| 资源 | 旧 URL | 新 URL | 分类 |
|------|--------|--------|------|
| 客户管理 | `/api/fams/customer` | `/api/fams/bank/customer` | bank.go |
| 账户管理 | `/api/fams/account` | `/api/fams/bank/account` | bank.go |
| 存款审核 | `/api/fams/deposit` | `/api/fams/bank/deposit` | bank.go |
| 提现审核 | `/api/fams/withdrawal` | `/api/fams/bank/withdrawal` | bank.go |

### 2. 公开接口统一

所有公开接口（无需认证）统一使用 `/api/public/` 前缀：

```
/api/public/admin/captcha       # 管理员验证码
/api/public/admin/login         # 管理员登录
/api/public/iam/login           # 用户登录
/api/public/iam/register        # 用户注册
/api/public/iam/email/send      # 发送邮箱验证码
/api/public/iam/email/verify    # 验证邮箱
```

### 3. 移除请求头说明

所有接口描述中移除 `"请求头: Authorization, X-Timestamp, X-Sign"` 说明：

**变更前**：
```go
@doc(
    summary: "创建员工"
    description: "创建后台人员账号。请求头: Authorization, X-Timestamp, X-Sign"
    id: "admin.staff.create"
)
```

**变更后**：
```go
@doc(
    summary: "创建员工"
    description: "创建后台人员账号"
    id: "admin.sys.staff.create"
)
```

### 4. 时间格式统一

所有时间字段统一使用**毫秒级时间戳 int64**：

**变更前**：
```go
CreatedAt string `json:"createdAt"` // 创建时间
UpdatedAt string `json:"updatedAt"` // 更新时间
ExpiresAt string `json:"expiresAt"` // 过期时间（ISO 8601格式）
```

**变更后**：
```go
CreatedAt int64 `json:"createdAt"` // 创建时间（毫秒级时间戳）
UpdatedAt int64 `json:"updatedAt"` // 更新时间（毫秒级时间戳）
ExpiresAt int64 `json:"expiresAt"` // 过期时间（毫秒级时间戳）
```

### 5. 中间件移除

所有 .api 文件中移除：
- ❌ `jwt: Auth`
- ❌ `middleware: SignCheck, JwtAuth`
- ✅ 改用全局 Cookie 认证中间件

### 6. Group 命名调整

Group 命名加入分类前缀，更清晰地标识资源归属：

| 资源 | 旧 Group | 新 Group |
|------|---------|----------|
| 员工管理 | `staff` | `sysStaff` |
| 角色管理 | `role` | `sysRole` |
| 菜单管理 | `menu` | `sysMenu` |
| 国家管理 | `country` | `dataCountry` |
| 用户管理 | `user` | `iamUser` |
| 邀请管理 | `invite` | `iamUserInvite` |
| 客户管理 | `customer` | `famsBankCustomer` |
| 账户管理 | `account` | `famsBankAccount` |
| 存款审核 | `auditDeposit` | `famsBankDeposit` |
| 提现审核 | `auditWithdrawal` | `famsBankWithdrawal` |

## 修改文件清单

### 新增文件
1. `docs/akatm-api.md` - 完整的API规范文档

### 修改文件（共11个.api文件）

#### Admin RPC（5个文件）
1. ✅ `api/admin/docs/auth.api` - 认证接口
   - 公开接口 prefix: `/api/public/admin`
   - 认证接口 prefix: `/api/admin/auth`
   - 时间字段改为 int64

2. ✅ `api/admin/docs/staff.api` - 员工管理
   - prefix: `/api/admin` → `/api/admin/sys`
   - group: `staff` → `sysStaff`
   - 移除请求头说明
   - 时间字段改为 int64

3. ✅ `api/admin/docs/role.api` - 角色管理
   - prefix: `/api/admin` → `/api/admin/sys`
   - group: `role` → `sysRole`
   - 移除请求头说明
   - 时间字段改为 int64

4. ✅ `api/admin/docs/menu.api` - 菜单管理
   - prefix: `/api/admin` → `/api/admin/sys`
   - group: `menu` → `sysMenu`
   - 移除请求头说明

5. ✅ `api/admin/docs/country.api` - 国家管理
   - prefix: `/api/admin` → `/api/admin/data`
   - group: `country` → `dataCountry`
   - 移除请求头说明

#### IAM RPC（2个文件）
6. ✅ `api/admin/docs/user.api` - 用户管理
   - prefix: `/api/iam` → `/api/iam/user`
   - group: `user` → `iamUser`
   - 移除请求头说明

7. ✅ `api/admin/docs/invite.api` - 邀请管理
   - prefix: `/api/iam` → `/api/iam/user`
   - group: `invite` → `iamUserInvite`
   - 移除请求头说明

#### FAMS RPC（4个文件）
8. ✅ `api/admin/docs/customer.api` - 客户管理
   - prefix: `/api/fams` → `/api/fams/bank`
   - group: `customer` → `famsBankCustomer`
   - 移除请求头说明

9. ✅ `api/admin/docs/account.api` - 账户管理
   - prefix: `/api/fams` → `/api/fams/bank`
   - group: `account` → `famsBankAccount`
   - 移除请求头说明

10. ✅ `api/admin/docs/deposit.api` - 存款审核
    - prefix: `/api/fams` → `/api/fams/bank`
    - group: `auditDeposit` → `famsBankDeposit`
    - 移除请求头说明

11. ✅ `api/admin/docs/withdrawal.api` - 提现审核
    - prefix: `/api/fams` → `/api/fams/bank`
    - group: `auditWithdrawal` → `famsBankWithdrawal`
    - 移除请求头说明

## URL 映射对照表

### Admin RPC 完整映射

| 操作 | HTTP 方法 | 旧 URL | 新 URL |
|------|----------|--------|--------|
| **员工管理** | | | |
| 创建员工 | POST | `/api/admin/staff` | `/api/admin/sys/staff` |
| 员工列表 | GET | `/api/admin/staff` | `/api/admin/sys/staff` |
| 员工详情 | GET | `/api/admin/staff/:id` | `/api/admin/sys/staff/:id` |
| 更新员工 | PUT | `/api/admin/staff/:id` | `/api/admin/sys/staff/:id` |
| 删除员工 | DELETE | `/api/admin/staff/:id` | `/api/admin/sys/staff/:id` |
| 重设密码 | PATCH | `/api/admin/staff/:id/password` | `/api/admin/sys/staff/:id/password` |
| 设置角色 | PATCH | `/api/admin/staff/:id/role` | `/api/admin/sys/staff/:id/role` |
| 更新状态 | PATCH | `/api/admin/staff/:id/status` | `/api/admin/sys/staff/:id/status` |
| 获取员工菜单 | GET | `/api/admin/staff/:id/menus` | `/api/admin/sys/staff/:id/menu` |
| **角色管理** | | | |
| 创建角色 | POST | `/api/admin/role` | `/api/admin/sys/role` |
| 角色列表 | GET | `/api/admin/role` | `/api/admin/sys/role` |
| 角色详情 | GET | `/api/admin/role/:id` | `/api/admin/sys/role/:id` |
| 更新角色 | PUT | `/api/admin/role/:id` | `/api/admin/sys/role/:id` |
| 删除角色 | DELETE | `/api/admin/role/:id` | `/api/admin/sys/role/:id` |
| 分配菜单 | PATCH | `/api/admin/role/:id/menus` | `/api/admin/sys/role/:id/menu` |
| 获取角色菜单 | GET | `/api/admin/role/:id/menus` | `/api/admin/sys/role/:id/menu` |
| **菜单管理** | | | |
| 创建菜单 | POST | `/api/admin/menu` | `/api/admin/sys/menu` |
| 菜单列表 | GET | `/api/admin/menu` | `/api/admin/sys/menu` |
| 菜单详情 | GET | `/api/admin/menu/:id` | `/api/admin/sys/menu/:id` |
| 更新菜单 | PUT | `/api/admin/menu/:id` | `/api/admin/sys/menu/:id` |
| 删除菜单 | DELETE | `/api/admin/menu/:id` | `/api/admin/sys/menu/:id` |
| 更新菜单状态 | PATCH | `/api/admin/menu/:id/status` | `/api/admin/sys/menu/:id/status` |
| **国家管理** | | | |
| 创建国家 | POST | `/api/admin/country` | `/api/admin/data/country` |
| 国家列表 | GET | `/api/admin/country` | `/api/admin/data/country` |
| 国家详情 | GET | `/api/admin/country/:id` | `/api/admin/data/country/:id` |
| 更新国家 | PUT | `/api/admin/country/:id` | `/api/admin/data/country/:id` |
| 删除国家 | DELETE | `/api/admin/country/:id` | `/api/admin/data/country/:id` |
| 更新国家状态 | PATCH | `/api/admin/country/:id/status` | `/api/admin/data/country/:id/status` |

### IAM RPC 完整映射

| 操作 | HTTP 方法 | 旧 URL | 新 URL |
|------|----------|--------|--------|
| **用户管理** | | | |
| 创建总代 | POST | `/api/iam/user/super-agent` | `/api/iam/user/super-agent` |
| 创建代理 | POST | `/api/iam/user/agent` | `/api/iam/user/agent` |
| 创建客户经理 | POST | `/api/iam/user/manager` | `/api/iam/user/manager` |
| 用户列表 | GET | `/api/iam/user` | `/api/iam/user` |
| 总代列表 | GET | `/api/iam/user/super-agent` | `/api/iam/user/super-agent` |
| 代理列表 | GET | `/api/iam/user/agent` | `/api/iam/user/agent` |
| 客户经理列表 | GET | `/api/iam/user/manager` | `/api/iam/user/manager` |
| 用户详情 | GET | `/api/iam/user/:id` | `/api/iam/user/:id` |
| 更新用户 | PUT | `/api/iam/user/:id` | `/api/iam/user/:id` |
| 删除用户 | DELETE | `/api/iam/user/:id` | `/api/iam/user/:id` |
| 更新状态 | PATCH | `/api/iam/user/:id/status` | `/api/iam/user/:id/status` |
| 设置手续费 | PATCH | `/api/iam/user/:id/fees` | `/api/iam/user/:id/fee` |
| 授权国家 | PATCH | `/api/iam/user/:id/countries` | `/api/iam/user/:id/country` |
| **邀请管理** | | | |
| 生成邀请链接 | POST | `/api/iam/user-invite` | `/api/iam/user/invite` |
| 邀请记录列表 | GET | `/api/iam/user-invite` | `/api/iam/user/invite` |
| 邀请详情 | GET | `/api/iam/user-invite/:id` | `/api/iam/user/invite/:id` |
| 禁用邀请链接 | PATCH | `/api/iam/user-invite/:id/disable` | `/api/iam/user/invite/:id/disable` |

### FAMS RPC 完整映射

| 操作 | HTTP 方法 | 旧 URL | 新 URL |
|------|----------|--------|--------|
| **客户管理** | | | |
| 客户列表 | GET | `/api/fams/customer` | `/api/fams/bank/customer` |
| 客户详情 | GET | `/api/fams/customer/:id` | `/api/fams/bank/customer/:id` |
| 创建客户 | POST | `/api/fams/customer` | `/api/fams/bank/customer` |
| 更新客户 | PUT | `/api/fams/customer/:id` | `/api/fams/bank/customer/:id` |
| 删除客户 | DELETE | `/api/fams/customer/:id` | `/api/fams/bank/customer/:id` |
| 更新KYC状态 | PATCH | `/api/fams/customer/:id/kyc-status` | `/api/fams/bank/customer/:id/kyc-status` |
| **账户管理** | | | |
| 账户列表 | GET | `/api/fams/account` | `/api/fams/bank/account` |
| 账户详情 | GET | `/api/fams/account/:id` | `/api/fams/bank/account/:id` |
| 创建账户 | POST | `/api/fams/account` | `/api/fams/bank/account` |
| 更新账户 | PUT | `/api/fams/account/:id` | `/api/fams/bank/account/:id` |
| 删除账户 | DELETE | `/api/fams/account/:id` | `/api/fams/bank/account/:id` |
| 更新账户状态 | PATCH | `/api/fams/account/:id/status` | `/api/fams/bank/account/:id/status` |
| 账户余额 | GET | `/api/fams/account/:id/balance` | `/api/fams/bank/account/:id/balance` |
| 账户流水 | GET | `/api/fams/account/:id/transactions` | `/api/fams/bank/account/:id/transaction` |
| 账户统计 | GET | `/api/fams/account/stats` | `/api/fams/bank/account/stats` |
| **存款审核** | | | |
| 存款列表 | GET | `/api/fams/deposit` | `/api/fams/bank/deposit` |
| 存款详情 | GET | `/api/fams/deposit/:id` | `/api/fams/bank/deposit/:id` |
| 审核通过 | POST | `/api/fams/deposit/:id/approve` | `/api/fams/bank/deposit/:id/approve` |
| 审核拒绝 | POST | `/api/fams/deposit/:id/reject` | `/api/fams/bank/deposit/:id/reject` |
| **提现审核** | | | |
| 提现列表 | GET | `/api/fams/withdrawal` | `/api/fams/bank/withdrawal` |
| 提现详情 | GET | `/api/fams/withdrawal/:id` | `/api/fams/bank/withdrawal/:id` |
| 审核通过 | POST | `/api/fams/withdrawal/:id/approve` | `/api/fams/bank/withdrawal/:id/approve` |
| 审核拒绝 | POST | `/api/fams/withdrawal/:id/reject` | `/api/fams/bank/withdrawal/:id/reject` |
| 批量审核通过 | POST | `/api/fams/withdrawal/batch-approve` | `/api/fams/bank/withdrawal/batch-approve` |
| 批量审核拒绝 | POST | `/api/fams/withdrawal/batch-reject` | `/api/fams/bank/withdrawal/batch-reject` |
| 提现统计 | GET | `/api/fams/withdrawal/stats` | `/api/fams/bank/withdrawal/stats` |

## 数据格式变更

### 时间字段示例

**登录响应（变更前）**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "staffId": 1234567890123456789,
    "name": "张三",
    "email": "zhangsan@example.com",
    "expiresAt": "2025-10-23T12:00:00Z",
    "createdAt": "2025-10-22T10:30:00Z"
  }
}
```

**登录响应（变更后）**：
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "staffId": 1234567890123456789,
    "name": "张三",
    "email": "zhangsan@example.com",
    "expiresAt": 1729688400000,
    "createdAt": 1729594200000
  }
}
```

## 兼容性说明

### ⚠️ 破坏性变更

本次重构包含以下破坏性变更：

1. **URL 路径变更** - 所有API路径都发生了变化
2. **时间格式变更** - 从 ISO 8601 字符串改为毫秒级时间戳
3. **认证方式变更** - 从 JWT Token 改为 Cookie 认证
4. **请求头变更** - 不再需要 Authorization, X-Timestamp, X-Sign 请求头

### ✅ 前端适配建议

1. **更新所有API调用URL**
   ```javascript
   // 旧代码
   const response = await fetch('/api/admin/staff')

   // 新代码
   const response = await fetch('/api/admin/sys/staff')
   ```

2. **时间处理统一使用时间戳**
   ```javascript
   // 旧代码
   const date = new Date(data.createdAt) // ISO 8601 字符串

   // 新代码
   const date = new Date(data.createdAt) // 毫秒级时间戳（同样能工作）
   ```

3. **移除手动设置的认证头**
   ```javascript
   // 旧代码
   const response = await fetch('/api/admin/staff', {
     headers: {
       'Authorization': `Bearer ${token}`,
       'X-Timestamp': timestamp,
       'X-Sign': sign
     }
   })

   // 新代码
   const response = await fetch('/api/admin/sys/staff', {
     credentials: 'include' // 自动携带 Cookie
   })
   ```

## 迁移检查清单

### 后端检查
- [ ] 更新所有 API 路由配置
- [ ] 实现全局 Cookie 认证中间件
- [ ] 配置公开路径白名单
- [ ] 更新时间字段序列化逻辑（time.Time → int64）
- [ ] 创建 Snowflake ID 生成器
- [ ] 创建 Session 管理表
- [ ] 更新登录逻辑（生成 Cookie）
- [ ] 更新登出逻辑（清除 Cookie）

### 前端检查
- [ ] 更新所有 API 调用 URL
- [ ] 移除手动设置的认证头
- [ ] 添加 `credentials: 'include'` 到所有fetch请求
- [ ] 更新时间格式化逻辑
- [ ] 更新登录流程（自动保存 Cookie）
- [ ] 更新登出流程
- [ ] 更新401错误处理（跳转登录页）

### 测试检查
- [ ] 公开接口无需认证可访问
- [ ] 私有接口无 Cookie 返回 401
- [ ] 私有接口有效 Cookie 可访问
- [ ] Session 过期返回 401
- [ ] 登出后 Cookie 失效
- [ ] 时间字段正确序列化/反序列化
- [ ] 所有URL路径正确映射

## 参考文档

- [AKATM API 接口规范](./akatm-api.md)
- [全局中间件配置指南](./global-middleware-guide.md)
- [认证机制指南](./authentication-guide.md)

---

**变更完成日期**：2025-10-22
**文档版本**：v1.0
