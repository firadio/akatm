# Public API 使用指南

## 概述

`public.api` 是一个特殊的 API 定义文件，**集中管理所有无需认证的公开接口**。

### 设计目的

1. **集中管理** - 所有公开接口集中在一个文件中，便于维护和查看
2. **清晰区分** - 明确区分公开接口和需要认证的接口
3. **白名单配置** - 方便配置全局中间件的公开路径白名单
4. **安全审计** - 便于安全审计，确保只有必要的接口是公开的

---

## 文件位置

```
api/admin/docs/public.api
```

---

## 包含的公开接口

### 1. Admin 管理员公开接口

**URL前缀**: `/api/public/admin`

| 方法 | 路径 | 说明 | Handler |
|------|------|------|---------|
| `GET` | `/api/public/admin/captcha` | 获取图片验证码 | `getAdminCaptcha` |
| `POST` | `/api/public/admin/login` | 管理员登录 | `adminLogin` |

**使用场景**:
- 管理后台登录页面
- 验证码刷新

### 2. IAM 用户公开接口

**URL前缀**: `/api/public/iam`

| 方法 | 路径 | 说明 | Handler |
|------|------|------|---------|
| `GET` | `/api/public/iam/captcha` | 获取图片验证码 | `getIamCaptcha` |
| `POST` | `/api/public/iam/login` | 用户登录 | `iamLogin` |
| `POST` | `/api/public/iam/register` | 用户注册 | `iamRegister` |

**使用场景**:
- 前台用户登录页面
- 用户注册页面（通过邀请码）
- 验证码刷新

### 3. IAM 邮箱验证公开接口

**URL前缀**: `/api/public/iam/email`

| 方法 | 路径 | 说明 | Handler |
|------|------|------|---------|
| `POST` | `/api/public/iam/email/send` | 发送邮箱验证码 | `sendEmailCode` |
| `POST` | `/api/public/iam/email/verify` | 验证邮箱验证码 | `verifyEmailCode` |

**使用场景**:
- 用户注册时验证邮箱
- 重置密码时验证邮箱
- 绑定新邮箱时验证

---

## 接口详细说明

### 1. 获取图片验证码

**Admin端**:
```
GET /api/public/admin/captcha
```

**IAM端**:
```
GET /api/public/iam/captcha
```

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "captchaId": "abc123",
    "captchaImage": "data:image/png;base64,iVBORw0KG...",
    "expiresAt": 1704067500000
  }
}
```

### 2. 管理员登录

```
POST /api/public/admin/login
```

**请求**:
```json
{
  "email": "admin@example.com",
  "password": "password123",
  "captchaId": "abc123",
  "captcha": "1234"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "staffId": 1234567890123456789,
    "name": "管理员",
    "email": "admin@example.com",
    "status": 1,
    "roles": [
      {
        "id": 1,
        "name": "超级管理员",
        "code": "super_admin"
      }
    ],
    "menus": [...],
    "expiresAt": 1704153600000
  }
}
```

**响应头（自动设置）**:
```
Set-Cookie: user_id=1234567890123456789; Path=/; HttpOnly; Secure; SameSite=Strict
Set-Cookie: session_token=e3b0c44298fc1c...; Path=/; HttpOnly; Secure; SameSite=Strict; Max-Age=86400
```

### 3. 用户登录

```
POST /api/public/iam/login
```

**请求**:
```json
{
  "email": "user@example.com",
  "password": "password123",
  "captchaId": "abc123",
  "captcha": "1234"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "userId": 9876543210987654321,
    "userType": "manager",
    "email": "user@example.com",
    "nickname": "张经理",
    "status": 1,
    "parentId": 1111111111111111111,
    "parentName": "李代理",
    "expiresAt": 1704153600000
  }
}
```

### 4. 用户注册

```
POST /api/public/iam/register
```

**请求**:
```json
{
  "inviteCode": "INVITE123456",
  "email": "newuser@example.com",
  "password": "password123",
  "confirmPassword": "password123",
  "emailCode": "123456",
  "nickname": "新用户",
  "phone": "13800138000"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "userId": 5555555555555555555,
    "userType": "manager",
    "email": "newuser@example.com",
    "nickname": "新用户",
    "parentId": 1111111111111111111,
    "createdAt": 1704067200000
  }
}
```

### 5. 发送邮箱验证码

```
POST /api/public/iam/email/send
```

**请求**:
```json
{
  "email": "user@example.com",
  "scene": "register"
}
```

**场景类型**:
- `register` - 注册
- `login` - 登录
- `reset_password` - 重置密码
- `bind_email` - 绑定邮箱

**响应**:
```json
{
  "code": 0,
  "message": "验证码已发送",
  "data": null
}
```

### 6. 验证邮箱验证码

```
POST /api/public/iam/email/verify
```

**请求**:
```json
{
  "email": "user@example.com",
  "code": "123456",
  "scene": "register"
}
```

**响应**:
```json
{
  "code": 0,
  "message": "验证成功",
  "data": null
}
```

---

## 中间件配置

### 全局中间件白名单

在全局 Cookie 认证中间件中，需要配置公开路径白名单：

```go
// 公开路径白名单（不需要认证）
var publicPaths = []string{
    "/api/public/admin/captcha",
    "/api/public/admin/login",
    "/api/public/iam/captcha",
    "/api/public/iam/login",
    "/api/public/iam/register",
    "/api/public/iam/email/send",
    "/api/public/iam/email/verify",
}

// 或者使用前缀匹配（推荐）
func isPublicPath(path string) bool {
    return strings.HasPrefix(path, "/api/public/")
}
```

### 推荐配置方式

**使用前缀匹配**（更简洁）:

```go
func CookieAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 公开路径直接放行
        if strings.HasPrefix(r.URL.Path, "/api/public/") {
            next.ServeHTTP(w, r)
            return
        }

        // 验证 Cookie...
        // ...
    })
}
```

---

## 前端集成示例

### 1. 管理员登录

```javascript
// 获取验证码
const getCaptcha = async () => {
  const res = await fetch('/api/public/admin/captcha', {
    credentials: 'include'
  });
  const data = await res.json();
  return data.data; // { captchaId, captchaImage, expiresAt }
};

// 登录
const adminLogin = async (email, password, captchaId, captcha) => {
  const res = await fetch('/api/public/admin/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include', // 重要：允许携带 Cookie
    body: JSON.stringify({ email, password, captchaId, captcha })
  });

  const data = await res.json();

  if (data.code === 0) {
    // 登录成功，Cookie 已自动保存
    console.log('登录成功', data.data);
    // 跳转到管理后台首页
    window.location.href = '/admin/dashboard';
  } else {
    // 登录失败
    console.error('登录失败', data.message);
  }
};
```

### 2. 用户注册

```javascript
// 发送邮箱验证码
const sendEmailCode = async (email, scene = 'register') => {
  const res = await fetch('/api/public/iam/email/send', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify({ email, scene })
  });

  const data = await res.json();
  return data.code === 0;
};

// 注册
const register = async (formData) => {
  const res = await fetch('/api/public/iam/register', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(formData)
  });

  const data = await res.json();

  if (data.code === 0) {
    console.log('注册成功', data.data);
    // 跳转到登录页
    window.location.href = '/login';
  } else {
    console.error('注册失败', data.message);
  }
};
```

---

## 安全建议

### 1. 验证码机制

- ✅ 验证码有效期不超过 5 分钟
- ✅ 验证码使用后立即失效
- ✅ 同一邮箱短时间内限制发送次数（防止滥用）

### 2. 登录限制

- ✅ 限制登录失败次数（如 5 次失败后锁定账户 15 分钟）
- ✅ 记录登录日志（IP、时间、设备信息）
- ✅ 异常登录告警（异地登录、频繁登录等）

### 3. 密码安全

- ✅ 密码长度要求 8-20 位
- ✅ 密码必须包含字母和数字
- ✅ 服务端使用 bcrypt 加密存储
- ✅ 传输时使用 HTTPS

### 4. 邮箱验证码

- ✅ 验证码 6 位数字
- ✅ 有效期 10 分钟
- ✅ 同一邮箱 1 分钟内只能发送一次
- ✅ 验证码使用后立即失效

### 5. 邀请码

- ✅ 邀请码唯一性校验
- ✅ 邀请码过期时间校验
- ✅ 邀请码使用次数限制
- ✅ 记录邀请关系链

---

## 与其他文件的关系

### public.api vs admin_auth.api

| 文件 | 用途 | URL前缀 | 认证要求 |
|------|------|---------|---------|
| `public.api` | 公开接口（登录、注册） | `/api/public/*` | ❌ 无需认证 |
| `admin_auth.api` | 会话管理（登出、修改密码） | `/api/admin/auth/*` | ✅ 需要认证 |

**职责划分**:
- `public.api` - 负责**获取**认证凭证（登录、注册）
- `admin_auth.api` - 负责**管理**认证凭证（登出、修改密码、获取当前用户信息）

### 类型定义

`public.api` 中定义的类型可以在其他文件中引用（通过 import），避免重复定义。

---

## 测试清单

### 功能测试

- [ ] 管理员验证码可以正常生成
- [ ] 管理员可以使用邮箱密码登录
- [ ] 验证码错误时登录失败
- [ ] 密码错误时登录失败
- [ ] 登录成功后 Cookie 自动设置
- [ ] 用户验证码可以正常生成
- [ ] 用户可以使用邮箱密码登录
- [ ] 邮箱验证码可以正常发送
- [ ] 邮箱验证码验证正确
- [ ] 用户可以通过邀请码注册
- [ ] 注册时邮箱验证码校验正确

### 安全测试

- [ ] 验证码 5 分钟后失效
- [ ] 验证码使用后立即失效
- [ ] 登录失败 5 次后账户锁定
- [ ] 同一邮箱 1 分钟内只能发送一次验证码
- [ ] 邀请码只能使用一次
- [ ] 密码传输使用 HTTPS
- [ ] Cookie 设置了 HttpOnly 属性
- [ ] Cookie 设置了 Secure 属性（生产环境）
- [ ] Cookie 设置了 SameSite=Strict 属性

### 性能测试

- [ ] 验证码生成速度 < 100ms
- [ ] 登录响应时间 < 500ms
- [ ] 邮件发送不阻塞登录流程（异步）

---

## 常见问题

### Q1: 为什么要单独创建 public.api？

**A**: 集中管理公开接口有以下优势：
1. 便于安全审计 - 一目了然知道哪些接口是公开的
2. 便于配置中间件 - 使用前缀匹配即可放行所有公开接口
3. 便于维护 - 避免公开接口分散在多个文件中

### Q2: 前端需要手动设置 Cookie 吗？

**A**: 不需要。后端通过 `Set-Cookie` 响应头自动设置 Cookie，前端只需要在请求中添加 `credentials: 'include'` 即可。

### Q3: 公开接口需要添加中间件吗？

**A**: 不需要。全局中间件通过路径前缀判断，`/api/public/` 开头的路径会自动跳过认证。

### Q4: 如何添加新的公开接口？

**A**:
1. 在 `public.api` 中添加接口定义
2. URL 路径必须以 `/api/public/` 开头
3. 重新生成代码：`goctl api go -api public.api -dir .`

---

## 参考文档

- [AKATM API 接口规范](./akatm-api.md)
- [全局中间件配置指南](./global-middleware-guide.md)
- [认证机制指南](./authentication-guide.md)
- [API 文件清单](./api-files-inventory.md)

---

**更新日期**: 2025-10-22
**版本**: v1.0
