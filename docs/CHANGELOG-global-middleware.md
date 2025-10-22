# 全局中间件架构变更日志

## 变更日期
2025-10-22

## 变更原因

用户反馈："我的意思是接口本身不用中间件，而是全局中间件，比接口的优先级更高"

之前的实现在每个 .api 文件的 `@server` 块中声明了 `middleware: CookieAuth`，这会导致：
1. 中间件在路由级别执行，优先级不够高
2. 需要在多个文件中重复声明
3. 需要将公开接口和私有接口分离到不同的 `@server` 块

## 变更内容

### 1. 移除所有 .api 文件中的中间件声明

**修改的文件（共11个）**：

#### Admin RPC (管理后台)
- `api/admin/docs/auth.api` - 移除 `middleware: CookieAuth`
- `api/admin/docs/staff.api` - 移除 `middleware: CookieAuth`
- `api/admin/docs/role.api` - 移除 `middleware: CookieAuth`
- `api/admin/docs/menu.api` - 移除 `middleware: CookieAuth`
- `api/admin/docs/country.api` - 移除 `middleware: CookieAuth`

#### IAM RPC (身份认证)
- `api/admin/docs/invite.api` - 移除 `middleware: CookieAuth`
- `api/admin/docs/user.api` - 移除 `middleware: CookieAuth`

#### FAMS RPC (财务账户)
- `api/admin/docs/customer.api` - 移除 `middleware: CookieAuth`
- `api/admin/docs/account.api` - 移除 `middleware: CookieAuth`
- `api/admin/docs/deposit.api` - 移除 `middleware: CookieAuth`
- `api/admin/docs/withdrawal.api` - 移除 `middleware: CookieAuth`

**变更示例**：

```diff
@server(
    prefix: /api/admin
    group: staff
-   middleware: CookieAuth
    tags: "员工管理"
)
service admin {
    @handler createStaff
    post /staff (CreateStaffReq) returns (CreateStaffResp)
    // ...
}
```

### 2. 新增全局中间件配置文档

**新增文件**：
- `docs/global-middleware-guide.md` - 全局中间件配置指南

**主要内容**：
1. 为什么使用全局中间件（优势对比）
2. Go-Zero 全局中间件配置方法
3. 中间件实现示例代码
4. 公开路径白名单管理（3种方式）
5. Session 验证逻辑完整示例
6. Handler 中获取用户信息的方法
7. 多服务架构下的中间件复用
8. 最佳实践和避免做法
9. 测试验证方法

### 3. 更新认证指南

**修改文件**：
- `docs/authentication-guide.md`

**新增章节**：
- "中间件架构" - 解释全局中间件 vs 路由级中间件
- "中间件注册示例" - 展示如何在服务启动代码中注册

**更新说明**：
- 明确说明 .api 文件中**不声明**中间件
- 公开接口由全局中间件白名单控制
- 认证接口由全局中间件自动验证

## 实现架构

### 旧架构（路由级中间件）

```
.api 文件定义
    ↓
@server(middleware: CookieAuth) ← 每个文件都要声明
    ↓
生成的路由代码 → 路由级中间件执行
    ↓
Handler 处理
```

**缺点**：
- ❌ 需要在多个 .api 文件中重复声明
- ❌ 公开/私有接口需要分离到不同的 @server 块
- ❌ 中间件优先级不够高

### 新架构（全局中间件）

```
服务启动
    ↓
server.Use(CookieAuthMiddleware) ← 一次注册，全局生效
    ↓
全局中间件执行（优先级最高）
    ├─ 检查白名单 → 公开接口 → 直接放行
    └─ 验证 Cookie → 私有接口 → 验证通过/401
        ↓
路由匹配
    ↓
Handler 处理
```

**优点**：
- ✅ 集中配置，一处维护
- ✅ 优先级高于路由
- ✅ .api 文件只关注业务定义
- ✅ 灵活的白名单管理

## 公开接口白名单

### Admin RPC
- `GET /api/admin/auth/captcha` - 获取验证码
- `POST /api/admin/auth/login` - 管理员登录

### IAM RPC
- `GET /api/iam/auth/captcha` - 获取验证码
- `POST /api/iam/auth/login` - 用户登录
- `POST /api/iam/auth/register` - 用户注册

### FAMS RPC
- 无公开接口（所有接口都需要认证）

## 待实现工作

### 1. 创建中间件实现

需要在各个服务中创建中间件：

```
api/admin/middleware/cookieauth.go
api/iam/middleware/cookieauth.go
api/fams/middleware/cookieauth.go
```

或者创建公共中间件：

```
common/middleware/cookieauth.go
```

### 2. 修改服务启动文件

在以下文件中注册全局中间件：

```
api/admin/admin.go
api/iam/iam.go
api/fams/fams.go
```

示例代码：

```go
func main() {
    server := rest.MustNewServer(c.RestConf)

    // 注册全局中间件
    server.Use(middleware.NewCookieAuthMiddleware(c.PublicPaths).Handle)

    // 注册路由
    handler.RegisterHandlers(server, ctx)
    server.Start()
}
```

### 3. 配置文件更新

在 `etc/admin.yaml` 等配置文件中添加公开路径配置：

```yaml
PublicPaths:
  - /api/admin/auth/captcha
  - /api/admin/auth/login
```

### 4. Session 表创建

创建 session 表（如果尚未创建）：

```sql
-- Admin 服务
CREATE TABLE `admin_staff_session` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `staff_id` BIGINT UNSIGNED NOT NULL,
  `session_token` VARCHAR(64) NOT NULL,
  `expires_at` DATETIME NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_token` (`session_token`),
  KEY `idx_staff_id` (`staff_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- IAM 服务
CREATE TABLE `iam_user_session` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `session_token` VARCHAR(64) NOT NULL,
  `expires_at` DATETIME NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_token` (`session_token`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

## 测试清单

### 功能测试

- [ ] 公开接口无需 Cookie 可访问
  - [ ] GET /api/admin/auth/captcha
  - [ ] POST /api/admin/auth/login
- [ ] 私有接口无 Cookie 返回 401
  - [ ] GET /api/admin/staff
  - [ ] GET /api/fams/customer
- [ ] 私有接口有效 Cookie 可访问
  - [ ] 登录获取 Cookie
  - [ ] 使用 Cookie 访问私有接口
- [ ] Session 过期返回 401
- [ ] 登出后 Cookie 失效

### 安全测试

- [ ] Cookie 包含 HttpOnly 属性
- [ ] Cookie 包含 Secure 属性（生产环境）
- [ ] Cookie 包含 SameSite=Strict 属性
- [ ] 无法通过 JavaScript 读取 Cookie
- [ ] Session token 使用 SHA-256 哈希存储
- [ ] 暴库后无法模拟会话

### 性能测试

- [ ] 全局中间件对性能影响可接受
- [ ] Session 查询有索引优化
- [ ] 并发请求下中间件正常工作

## 兼容性说明

### 不影响

- ✅ 现有的 API 接口定义
- ✅ 请求/响应数据结构
- ✅ 业务逻辑代码

### 需要调整

- ⚠️ 服务启动代码（需要注册全局中间件）
- ⚠️ 配置文件（需要添加公开路径配置）
- ⚠️ 中间件实现（从路由级改为全局级）

## 文档更新

- ✅ `docs/global-middleware-guide.md` - 新增全局中间件配置指南
- ✅ `docs/authentication-guide.md` - 更新中间件架构说明
- ✅ `docs/CHANGELOG-global-middleware.md` - 本变更日志
- ✅ `api/admin/docs/*.api` - 移除所有中间件声明（11个文件）

## 参考资料

- [全局中间件配置指南](./global-middleware-guide.md)
- [认证机制指南](./authentication-guide.md)
- [Go-Zero 官方文档 - 中间件](https://go-zero.dev/docs/tutorials/http/middleware/middleware)
