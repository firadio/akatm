# AKATM 认证机制指南

## 概述

AKATM 系统采用基于 Cookie 的会话认证机制，而非传统的 JWT Token 方式。这种设计在安全性、易用性和性能上都有显著优势。

## 核心设计

### 1. Cookie 认证 vs JWT

**为什么选择 Cookie 认证？**

| 对比项 | Cookie 认证 | JWT Token |
|--------|------------|-----------|
| **前端处理** | ✅ 浏览器自动携带 | ❌ 需手动设置 Header |
| **XSS 防护** | ✅ HttpOnly 完全隔离 | ❌ 存储在 localStorage 易被窃取 |
| **CSRF 防护** | ✅ SameSite=Strict | ⚠️ 需额外 CSRF Token |
| **会话管理** | ✅ 服务端可主动失效 | ❌ 无法撤销已签发的 Token |
| **暴库安全** | ✅ SHA-256 哈希无法反推 | ⚠️ 密钥泄露风险 |
| **性能** | ✅ 简单高效 | ⚠️ 每次验证签名开销 |

### 2. 雪花算法 (Snowflake) ID

**用户ID生成策略**

所有用户ID（包括后台员工和前台用户）都使用雪花算法生成：

```
64位 ID 结构:
┌─────────┬──────────────────────────────┬────────┬────────┬──────────────┐
│ 符号位  │      时间戳 (41位)            │ 数据中心│ 机器ID │  序列号(12位) │
│  (1位)  │  (精确到毫秒)                 │ (5位)  │ (5位)  │              │
└─────────┴──────────────────────────────┴────────┴────────┴──────────────┘
     0           时间戳                     0-31    0-31     0-4095
```

**优势**：
- ✅ 全局唯一：分布式环境下无冲突
- ✅ 趋势递增：有利于数据库索引性能
- ✅ 高性能：本地生成，每毫秒最多 4096 个 ID
- ✅ 包含时间信息：可从ID中解析出创建时间

**配置示例**：
```go
// 推荐配置
epoch := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli()
dataCenterID := 1  // 0-31
machineID := 1     // 0-31
```

### 3. Session Token 安全设计

**生成流程**

```
登录时:
1. 生成随机Token (UUID v4 或 crypto/rand 32字节)
   原始Token: "a1b2c3d4-e5f6-7890-abcd-ef1234567890"

2. 计算 SHA-256 哈希
   HashedToken: SHA256(原始Token)
   = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

3. 存储到数据库
   INSERT INTO admin_staff_session (staff_id, token, expires_at)
   VALUES (123456, 'e3b0c44...', NOW() + INTERVAL 24 HOUR)

4. 设置到Cookie (仅发送一次)
   Set-Cookie: user_id=123456; HttpOnly; Secure; SameSite=Strict
   Set-Cookie: session_token=e3b0c44...; HttpOnly; Secure; SameSite=Strict; Max-Age=86400
```

**验证流程**

```
每次请求:
1. 从 Cookie 读取 user_id 和 session_token

2. 查询数据库
   SELECT * FROM admin_staff_session
   WHERE staff_id = ? AND token = ? AND expires_at > NOW()

3. 验证成功 → 更新最后活跃时间
   UPDATE admin_staff_session SET last_active_at = NOW() WHERE id = ?

4. 验证失败 → 返回 401 Unauthorized
```

**防暴库攻击**

即使数据库被攻破，攻击者也无法利用：
- ❌ 数据库中只存储 SHA-256 哈希值
- ❌ 无法从哈希值反推原始 Token
- ❌ 无法使用哈希值模拟会话（Cookie 中的值也是哈希）
- ✅ 必须拥有原始 Token 才能通过验证

## 中间件架构

### 全局中间件 vs 路由级中间件

**本项目采用全局中间件处理认证**，而非在 .api 文件中声明中间件：

| 特性 | 全局中间件（本项目） | 路由级中间件 |
|------|---------------------|-------------|
| **配置位置** | 服务启动代码 | .api 定义文件 |
| **优先级** | ✅ 高于路由处理器 | 路由级别 |
| **维护性** | ✅ 集中管理 | 分散在多个文件 |
| **灵活性** | ✅ 动态判断公开/私有路径 | 需要分离 @server 块 |

**重要说明**：
- ✅ .api 文件中**不声明** `middleware: CookieAuth`
- ✅ 认证逻辑在服务启动时通过 `server.Use()` 全局注册
- ✅ 中间件内部通过白名单判断哪些路径是公开的

详细配置请参考：[全局中间件配置指南](./global-middleware-guide.md)

### 中间件注册示例

```go
// api/admin/admin.go
func main() {
    server := rest.MustNewServer(c.RestConf)

    // 注册全局认证中间件（优先级高于路由）
    server.Use(middleware.NewCookieAuthMiddleware([]string{
        "/api/admin/auth/captcha",  // 公开路径白名单
        "/api/admin/auth/login",
    }).Handle)

    // 注册路由
    handler.RegisterHandlers(server, ctx)
    server.Start()
}
```

## 接口规范

### 公开接口（无需认证）

只有以下接口是公开的，无需 Cookie（由全局中间件白名单控制）：

```
GET  /api/admin/auth/captcha    # 获取验证码
POST /api/admin/auth/login      # 管理员登录
POST /api/iam/auth/login        # 前台用户登录
POST /api/iam/auth/register     # 用户注册（通过邀请码）
```

### 认证接口（需要 Cookie）

所有其他接口都需要 Cookie 认证（由全局中间件自动验证）：

**管理后台 (Admin RPC)**
```
/api/admin/staff/*       # 员工管理
/api/admin/role/*        # 角色管理
/api/admin/menu/*        # 菜单管理
/api/admin/country/*     # 国家管理
/api/admin/auth/logout   # 登出
/api/admin/auth/info     # 获取当前用户信息
```

**用户身份 (IAM RPC)**
```
/api/iam/user/*          # 用户管理
/api/iam/user-invite/*   # 邀请管理
```

**财务账户 (FAMS RPC)**
```
/api/fams/customer/*     # 客户管理
/api/fams/account/*      # 账户管理
/api/fams/deposit/*      # 存款审核
/api/fams/withdrawal/*   # 提现审核
```

## Cookie 配置

### 开发环境

```go
// 开发环境配置（HTTP）
http.SetCookie(w, &http.Cookie{
    Name:     "user_id",
    Value:    strconv.FormatInt(userID, 10),
    Path:     "/",
    HttpOnly: true,
    SameSite: http.SameSiteStrictMode,
    // Secure: false (开发环境可以不启用)
})

http.SetCookie(w, &http.Cookie{
    Name:     "session_token",
    Value:    hashedToken,
    Path:     "/",
    MaxAge:   86400, // 24小时
    HttpOnly: true,
    SameSite: http.SameSiteStrictMode,
    // Secure: false (开发环境可以不启用)
})
```

### 生产环境

```go
// 生产环境配置（HTTPS）
http.SetCookie(w, &http.Cookie{
    Name:     "user_id",
    Value:    strconv.FormatInt(userID, 10),
    Path:     "/",
    HttpOnly: true,
    Secure:   true, // 生产环境必须启用
    SameSite: http.SameSiteStrictMode,
})

http.SetCookie(w, &http.Cookie{
    Name:     "session_token",
    Value:    hashedToken,
    Path:     "/",
    MaxAge:   86400, // 24小时
    HttpOnly: true,
    Secure:   true, // 生产环境必须启用
    SameSite: http.SameSiteStrictMode,
})
```

### Cookie 属性说明

| 属性 | 值 | 作用 |
|------|-----|------|
| `HttpOnly` | `true` | **防XSS**: JavaScript 无法读取，防止被盗取 |
| `Secure` | `true` | **HTTPS专用**: 仅在 HTTPS 连接中传输 |
| `SameSite` | `Strict` | **防CSRF**: 仅同站请求携带，跨站请求不携带 |
| `Path` | `/` | **作用域**: 整个站点都可访问 |
| `MaxAge` | `86400` | **有效期**: 24小时后自动过期 |

## 中间件实现

### CookieAuth 中间件伪代码

```go
func CookieAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 1. 读取 Cookie
        userIDCookie, err := r.Cookie("user_id")
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        sessionTokenCookie, err := r.Cookie("session_token")
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        // 2. 解析 user_id
        userID, err := strconv.ParseInt(userIDCookie.Value, 10, 64)
        if err != nil {
            http.Error(w, "Invalid user_id", http.StatusUnauthorized)
            return
        }

        // 3. 查询会话
        var session Session
        err = db.QueryRow(`
            SELECT id, staff_id, token, expires_at
            FROM admin_staff_session
            WHERE staff_id = ? AND token = ? AND expires_at > NOW()
        `, userID, sessionTokenCookie.Value).Scan(
            &session.ID, &session.StaffID, &session.Token, &session.ExpiresAt,
        )

        if err != nil {
            http.Error(w, "Session expired", http.StatusUnauthorized)
            return
        }

        // 4. 更新最后活跃时间
        db.Exec(`
            UPDATE admin_staff_session
            SET last_active_at = NOW()
            WHERE id = ?
        `, session.ID)

        // 5. 会话刷新（距离过期不足1小时时延长）
        if time.Until(session.ExpiresAt) < time.Hour {
            newExpiry := time.Now().Add(24 * time.Hour)
            db.Exec(`
                UPDATE admin_staff_session
                SET expires_at = ?
                WHERE id = ?
            `, newExpiry, session.ID)

            // 更新 Cookie 过期时间
            http.SetCookie(w, &http.Cookie{
                Name:     "session_token",
                Value:    sessionTokenCookie.Value,
                Path:     "/",
                MaxAge:   86400,
                HttpOnly: true,
                Secure:   true,
                SameSite: http.SameSiteStrictMode,
            })
        }

        // 6. 将用户信息放入 Context
        ctx := context.WithValue(r.Context(), "user_id", userID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

## 登录登出流程

### 后台管理员登录

**请求**
```http
POST /api/admin/auth/login HTTP/1.1
Content-Type: application/json

{
  "email": "admin@example.com",
  "password": "P@ssw0rd123",
  "captchaId": "cap_123456",
  "captcha": "ABCD"
}
```

**响应**
```http
HTTP/1.1 200 OK
Set-Cookie: user_id=1748123456789012345; Path=/; HttpOnly; Secure; SameSite=Strict
Set-Cookie: session_token=e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855; Path=/; HttpOnly; Secure; SameSite=Strict; Max-Age=86400
Content-Type: application/json

{
  "code": 0,
  "message": "登录成功",
  "data": {
    "staffId": 1748123456789012345,
    "name": "管理员",
    "email": "admin@example.com",
    "status": 1,
    "roles": [
      {"id": 1, "name": "超级管理员", "code": "admin"}
    ],
    "menus": [...],
    "expiresAt": "2024-01-02T12:00:00Z"
  }
}
```

### 前台用户登录

**请求**
```http
POST /api/iam/auth/login HTTP/1.1
Content-Type: application/json

{
  "email": "manager@example.com",
  "password": "P@ssw0rd123",
  "userType": "manager"
}
```

**响应**
```http
HTTP/1.1 200 OK
Set-Cookie: user_id=1748987654321098765; Path=/; HttpOnly; Secure; SameSite=Strict
Set-Cookie: session_token=a7f2c91e85b3d4f8c1a6e9b2d5f8c3a7e9b1d4f6c8a2e5b9d1f4c7a3e6b8d2f5; Path=/; HttpOnly; Secure; SameSite=Strict; Max-Age=86400
Content-Type: application/json

{
  "code": 0,
  "message": "登录成功",
  "data": {
    "userId": 1748987654321098765,
    "userType": "manager",
    "email": "manager@example.com",
    "parentId": 123456,
    "exchangeFeeRate": "0.0500",
    "expiresAt": "2024-01-02T12:00:00Z"
  }
}
```

### 登出

**请求**
```http
POST /api/admin/auth/logout HTTP/1.1
Cookie: user_id=1748123456789012345; session_token=e3b0c44...
```

**响应**
```http
HTTP/1.1 200 OK
Set-Cookie: user_id=; Path=/; MaxAge=0
Set-Cookie: session_token=; Path=/; MaxAge=0
Content-Type: application/json

{
  "code": 0,
  "message": "登出成功"
}
```

**服务端操作**
```sql
-- 删除会话记录
DELETE FROM admin_staff_session
WHERE staff_id = ? AND token = ?
```

## 前端集成

### 无需手动处理

前端开发时，**完全不需要处理认证逻辑**：

```javascript
// ✅ 正确：直接调用 API，浏览器自动携带 Cookie
async function getStaffList() {
  const response = await fetch('/api/admin/staff', {
    method: 'GET',
    credentials: 'same-origin' // 或 'include'
  });

  if (response.status === 401) {
    // 会话过期，跳转登录页
    window.location.href = '/login';
    return;
  }

  return await response.json();
}

// ❌ 错误：不需要手动设置 Authorization Header
// headers: {
//   'Authorization': 'Bearer xxx' // 不需要！
// }
```

### Axios 配置

```javascript
// axios 全局配置
import axios from 'axios';

const api = axios.create({
  baseURL: '/api',
  withCredentials: true, // 自动携带 Cookie
});

// 响应拦截器：处理 401
api.interceptors.response.use(
  response => response,
  error => {
    if (error.response?.status === 401) {
      // 会话过期，跳转登录
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export default api;
```

### 使用示例

```javascript
// 登录
await api.post('/admin/auth/login', {
  email: 'admin@example.com',
  password: 'P@ssw0rd123',
  captchaId: 'cap_123',
  captcha: 'ABCD'
});

// 登录后，所有请求自动携带 Cookie
const staff = await api.get('/admin/staff');
const roles = await api.get('/admin/role');

// 登出
await api.post('/admin/auth/logout');
```

## 数据库表设计

### 后台员工会话表

```sql
CREATE TABLE admin_staff_session (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    staff_id BIGINT NOT NULL COMMENT '员工ID（雪花算法）',
    token CHAR(64) NOT NULL COMMENT 'Session Token SHA-256哈希值',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    expires_at TIMESTAMP NOT NULL COMMENT '过期时间',
    last_active_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '最后活跃时间',
    ip_address VARCHAR(45) COMMENT '登录IP地址',
    user_agent VARCHAR(255) COMMENT '用户代理（浏览器）',

    INDEX idx_staff_token (staff_id, token),
    INDEX idx_expires (expires_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台员工会话表';
```

### 前台用户会话表

```sql
CREATE TABLE iam_user_session (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL COMMENT '用户ID（雪花算法）',
    token CHAR(64) NOT NULL COMMENT 'Session Token SHA-256哈希值',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    expires_at TIMESTAMP NOT NULL COMMENT '过期时间',
    last_active_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '最后活跃时间',
    ip_address VARCHAR(45) COMMENT '登录IP地址',
    user_agent VARCHAR(255) COMMENT '用户代理（浏览器）',

    INDEX idx_user_token (user_id, token),
    INDEX idx_expires (expires_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='前台用户会话表';
```

## 安全最佳实践

### 生产环境检查清单

- [x] **强制 HTTPS**: 所有接口必须使用 HTTPS
- [x] **HttpOnly Cookie**: 防止 XSS 攻击窃取 Session
- [x] **Secure Cookie**: 仅通过 HTTPS 传输
- [x] **SameSite=Strict**: 防止 CSRF 攻击
- [x] **SHA-256 哈希**: 数据库存储哈希值，防暴库
- [x] **会话过期**: 24小时自动过期
- [x] **自动刷新**: 距离过期不足1小时时自动延长
- [x] **并发登录限制**: 同一用户最多 N 个活跃会话
- [x] **异常IP检测**: 会话绑定IP，异地登录需二次验证
- [x] **登录失败锁定**: 5次失败锁定账户15分钟
- [x] **操作日志**: 记录所有登录、登出、敏感操作

### 会话管理策略

**并发登录限制**
```sql
-- 登录时检查会话数量
SELECT COUNT(*) FROM admin_staff_session
WHERE staff_id = ? AND expires_at > NOW();

-- 超过限制时，删除最旧的会话
DELETE FROM admin_staff_session
WHERE id IN (
    SELECT id FROM admin_staff_session
    WHERE staff_id = ?
    ORDER BY last_active_at ASC
    LIMIT 1
);
```

**定期清理过期会话**
```sql
-- 每小时执行一次
DELETE FROM admin_staff_session
WHERE expires_at < NOW();

DELETE FROM iam_user_session
WHERE expires_at < NOW();
```

**异常IP检测**
```go
// 登录时记录IP
session.IPAddress = r.RemoteAddr

// 验证时检查IP
if session.IPAddress != r.RemoteAddr {
    // 发送安全提醒邮件
    sendSecurityAlert(session.StaffID, "异地登录检测")
    // 可选：要求二次验证
}
```

## 部署架构

### 同域名部署（推荐）

```
┌─────────────────────────────────────┐
│   https://app.example.com           │
├─────────────────────────────────────┤
│                                     │
│  前端页面: /                        │
│  API接口:  /api/*                   │
│                                     │
│  Cookie Domain: app.example.com    │
│  浏览器自动携带Cookie              │
└─────────────────────────────────────┘
```

### 子域名部署

```
前端: https://www.example.com
API:  https://api.example.com

Cookie配置:
  Domain: .example.com  (注意前面的点)

允许跨子域名共享Cookie
```

## 常见问题

### Q1: 为什么不用 JWT？

**A**: Cookie 认证在以下方面优于 JWT：
1. **安全性**: HttpOnly Cookie 无法被 JavaScript 读取，完全防 XSS
2. **易用性**: 浏览器自动处理，前端无需手动设置 Header
3. **会话控制**: 服务端可主动失效会话，JWT 无法撤销
4. **防暴库**: 哈希存储，数据库泄露也无法模拟会话

### Q2: Cookie 的 CSRF 风险如何防范？

**A**: 使用 `SameSite=Strict` 属性：
- 只有同站请求才携带 Cookie
- 跨站请求（如钓鱼网站）无法携带 Cookie
- 配合 HTTPS，安全性极高

### Q3: 如何处理跨域请求？

**A**: 建议使用同域名部署，如果必须跨域：
```javascript
// 前端配置
axios.defaults.withCredentials = true;

// 后端配置 CORS
w.Header().Set("Access-Control-Allow-Origin", "https://www.example.com")
w.Header().Set("Access-Control-Allow-Credentials", "true")
```

### Q4: 会话过期后如何处理？

**A**:
1. 后端返回 401 状态码
2. 前端拦截器检测到 401
3. 自动跳转到登录页
4. 用户重新登录

### Q5: 如何实现"记住我"功能？

**A**: 调整 Cookie 的 MaxAge：
```go
// 普通登录：24小时
MaxAge: 86400

// 记住我：30天
MaxAge: 2592000
```

## 总结

AKATM 的 Cookie 认证方案具有以下优势：

1. ✅ **安全性高**: HttpOnly + Secure + SameSite + SHA-256 多重保护
2. ✅ **开发简单**: 前端无需处理认证逻辑，浏览器自动携带
3. ✅ **用户体验好**: 会话自动刷新，无需频繁登录
4. ✅ **可控性强**: 服务端可主动失效会话，支持强制登出
5. ✅ **防暴库攻击**: 即使数据库泄露，也无法模拟会话
6. ✅ **易于扩展**: 支持并发登录限制、异地登录检测等高级功能

结合雪花算法生成的用户ID，系统具备了高性能、高安全性、高可用性的认证基础。
