# 全局中间件配置指南

## 概述

本项目采用**全局中间件**方式处理Cookie认证，而非在API定义文件中声明中间件。全局中间件在路由注册之前执行，优先级高于路由级别的处理器。

## 为什么使用全局中间件？

### 优势对比

| 特性 | 全局中间件 | 路由级中间件 |
|------|------------|------------|
| **配置位置** | 应用初始化代码 | .api定义文件 |
| **优先级** | 高于路由处理器 | 路由级别 |
| **维护性** | 集中管理，一处配置 | 分散在多个文件 |
| **灵活性** | 可动态判断是否需要认证 | 需要分离公开/私有路由 |
| **代码生成** | 不影响go-zero代码生成 | 需要在每个@server块声明 |

### 架构优势

1. **单一职责原则**：API定义文件只关注业务接口定义，认证逻辑由全局中间件统一处理
2. **集中管理**：所有认证逻辑、公开路由白名单在一处维护
3. **优先执行**：全局中间件先于路由匹配执行，提前拦截未认证请求
4. **易于测试**：中间件独立于业务逻辑，方便单元测试

## Go-Zero 全局中间件配置

### 1. 目录结构

```
api/admin/
├── admin.go              # 服务启动入口
├── middleware/
│   └── cookieauth.go     # Cookie认证中间件
├── docs/
│   ├── auth.api          # 认证接口（无middleware声明）
│   ├── staff.api         # 员工管理（无middleware声明）
│   └── ...
└── admin.api             # 主API文件
```

### 2. 中间件实现

创建 `api/admin/middleware/cookieauth.go`:

```go
package middleware

import (
    "net/http"
    "strings"

    "github.com/zeromicro/go-zero/rest/httpx"
)

// 公开路径白名单（不需要认证的路径）
var publicPaths = map[string]bool{
    "/api/admin/auth/captcha": true,
    "/api/admin/auth/login":   true,
}

type CookieAuthMiddleware struct {
    // 可以注入依赖，如session仓库、用户仓库等
}

func NewCookieAuthMiddleware() *CookieAuthMiddleware {
    return &CookieAuthMiddleware{}
}

func (m *CookieAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 1. 检查是否为公开路径
        if publicPaths[r.URL.Path] {
            next(w, r)
            return
        }

        // 2. 获取Cookie中的认证信息
        userIdCookie, err := r.Cookie("user_id")
        if err != nil {
            httpx.WriteJson(w, http.StatusUnauthorized, map[string]interface{}{
                "code":    401,
                "message": "未登录或会话已过期",
            })
            return
        }

        sessionTokenCookie, err := r.Cookie("session_token")
        if err != nil {
            httpx.WriteJson(w, http.StatusUnauthorized, map[string]interface{}{
                "code":    401,
                "message": "未登录或会话已过期",
            })
            return
        }

        userId := userIdCookie.Value
        sessionToken := sessionTokenCookie.Value

        // 3. 验证session（从数据库或Redis查询）
        // TODO: 实现session验证逻辑
        // - 查询session是否存在
        // - 检查session是否过期
        // - 验证token哈希值

        // 4. 将用户信息存入context，供后续handler使用
        ctx := r.Context()
        // ctx = context.WithValue(ctx, "userId", userId)
        // ctx = context.WithValue(ctx, "staffId", userId) // 对于admin服务

        // 5. 继续执行下一个handler
        next(w, r.WithContext(ctx))
    }
}
```

### 3. 全局中间件注册

在 `api/admin/admin.go` 中注册全局中间件:

```go
package main

import (
    "flag"
    "fmt"

    "github.com/zeromicro/go-zero/core/conf"
    "github.com/zeromicro/go-zero/rest"

    "your-project/api/admin/handler"
    "your-project/api/admin/middleware"
    "your-project/api/admin/svc"
    "your-project/config"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
    flag.Parse()

    var c config.Config
    conf.MustLoad(*configFile, &c)

    server := rest.MustNewServer(c.RestConf)
    defer server.Stop()

    ctx := svc.NewServiceContext(c)

    // ===== 关键：注册全局中间件 =====
    server.Use(middleware.NewCookieAuthMiddleware().Handle)

    // 注册路由
    handler.RegisterHandlers(server, ctx)

    fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
    server.Start()
}
```

### 4. 公开路由白名单管理

#### 方式一：硬编码（简单但不够灵活）

```go
var publicPaths = map[string]bool{
    "/api/admin/auth/captcha": true,
    "/api/admin/auth/login":   true,
}
```

#### 方式二：配置文件（推荐）

在 `etc/admin.yaml` 中配置:

```yaml
Name: admin-api
Host: 0.0.0.0
Port: 8888

# 公开路径（不需要认证）
PublicPaths:
  - /api/admin/auth/captcha
  - /api/admin/auth/login
```

中间件读取配置:

```go
type CookieAuthMiddleware struct {
    publicPaths map[string]bool
}

func NewCookieAuthMiddleware(publicPaths []string) *CookieAuthMiddleware {
    pathMap := make(map[string]bool)
    for _, path := range publicPaths {
        pathMap[path] = true
    }
    return &CookieAuthMiddleware{
        publicPaths: pathMap,
    }
}
```

#### 方式三：路径前缀匹配（最灵活）

```go
func (m *CookieAuthMiddleware) isPublicPath(path string) bool {
    // 精确匹配
    if m.publicPaths[path] {
        return true
    }

    // 前缀匹配（用于批量开放某个路径下的所有接口）
    publicPrefixes := []string{
        "/api/admin/auth/",  // /api/admin/auth/* 都是公开的
    }

    for _, prefix := range publicPrefixes {
        if strings.HasPrefix(path, prefix) {
            return true
        }
    }

    return false
}
```

## Session 验证逻辑

### 数据库表结构

```sql
CREATE TABLE `admin_session` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `staff_id` BIGINT UNSIGNED NOT NULL COMMENT '员工ID（雪花算法）',
  `session_token` VARCHAR(64) NOT NULL COMMENT 'Session Token（SHA-256哈希值）',
  `expires_at` DATETIME NOT NULL COMMENT '过期时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_token` (`session_token`),
  KEY `idx_staff_id` (`staff_id`),
  KEY `idx_expires_at` (`expires_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员会话表';
```

### 中间件验证逻辑完整示例

```go
func (m *CookieAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 1. 检查是否为公开路径
        if m.isPublicPath(r.URL.Path) {
            next(w, r)
            return
        }

        // 2. 获取Cookie
        userIdCookie, err := r.Cookie("user_id")
        if err != nil {
            m.unauthorized(w, "未登录或会话已过期")
            return
        }

        sessionTokenCookie, err := r.Cookie("session_token")
        if err != nil {
            m.unauthorized(w, "未登录或会话已过期")
            return
        }

        userId := userIdCookie.Value
        sessionToken := sessionTokenCookie.Value

        // 3. 解析用户ID
        staffId, err := strconv.ParseInt(userId, 10, 64)
        if err != nil {
            m.unauthorized(w, "无效的用户ID")
            return
        }

        // 4. 查询session（从数据库或Redis）
        session, err := m.sessionRepo.FindByToken(r.Context(), sessionToken)
        if err != nil {
            m.unauthorized(w, "会话不存在或已失效")
            return
        }

        // 5. 验证session
        if session.StaffId != staffId {
            m.unauthorized(w, "会话与用户不匹配")
            return
        }

        if time.Now().After(session.ExpiresAt) {
            m.unauthorized(w, "会话已过期")
            return
        }

        // 6. 将用户信息存入context
        ctx := r.Context()
        ctx = context.WithValue(ctx, "staffId", staffId)
        ctx = context.WithValue(ctx, "sessionToken", sessionToken)

        // 7. 继续执行
        next(w, r.WithContext(ctx))
    }
}

func (m *CookieAuthMiddleware) unauthorized(w http.ResponseWriter, message string) {
    httpx.WriteJson(w, http.StatusUnauthorized, map[string]interface{}{
        "code":    401,
        "message": message,
    })
}
```

## 在Handler中获取用户信息

中间件验证通过后，handler可以从context中获取用户信息:

```go
package handler

import (
    "net/http"

    "github.com/zeromicro/go-zero/rest/httpx"
    "your-project/api/admin/logic"
    "your-project/api/admin/svc"
    "your-project/api/admin/types"
)

func GetStaffInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 从context中获取staffId（由全局中间件设置）
        staffId := r.Context().Value("staffId").(int64)

        l := logic.NewGetStaffInfoLogic(r.Context(), svcCtx)
        resp, err := l.GetStaffInfo(staffId)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
        } else {
            httpx.OkJsonCtx(r.Context(), w, resp)
        }
    }
}
```

## 多服务架构下的中间件复用

如果有多个服务（admin、iam、fams），可以将中间件提取到公共包:

```
common/
├── middleware/
│   ├── cookieauth.go     # Cookie认证中间件
│   └── config.go         # 中间件配置
api/
├── admin/
│   └── admin.go          # 引用common/middleware
├── iam/
│   └── iam.go            # 引用common/middleware
└── fams/
    └── fams.go           # 引用common/middleware
```

每个服务注册时传入不同的公开路径:

```go
// api/admin/admin.go
server.Use(middleware.NewCookieAuthMiddleware([]string{
    "/api/admin/auth/captcha",
    "/api/admin/auth/login",
}).Handle)

// api/iam/iam.go
server.Use(middleware.NewCookieAuthMiddleware([]string{
    "/api/iam/auth/captcha",
    "/api/iam/auth/login",
}).Handle)

// api/fams/fams.go
server.Use(middleware.NewCookieAuthMiddleware([]string{
    // FAMS服务可能所有接口都需要认证
}).Handle)
```

## 最佳实践总结

### ✅ 推荐做法

1. **全局中间件注册**：在服务启动代码中使用 `server.Use()` 注册
2. **配置化白名单**：公开路径通过配置文件管理，便于维护
3. **Context传递用户信息**：中间件验证通过后，将用户信息存入context
4. **分离认证与授权**：中间件只做认证（验证身份），权限控制在业务逻辑中处理
5. **统一错误响应**：认证失败返回统一的401响应格式

### ❌ 避免做法

1. **不要在.api文件中声明middleware**：这会导致代码生成时产生路由级中间件
2. **不要硬编码敏感信息**：如session密钥、过期时间等应通过配置管理
3. **不要在handler中重复验证**：认证已由全局中间件完成，handler只需从context获取用户信息
4. **不要混用多种认证方式**：项目统一使用Cookie认证，不要同时支持JWT等其他方式

## 测试验证

### 测试公开接口（无需Cookie）

```bash
# 获取验证码（公开接口）
curl http://localhost:8888/api/admin/auth/captcha

# 登录（公开接口）
curl -X POST http://localhost:8888/api/admin/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "password123",
    "captchaId": "xxx",
    "captcha": "1234"
  }' \
  -c cookies.txt  # 保存响应Cookie
```

### 测试私有接口（需要Cookie）

```bash
# 获取员工信息（需要认证）
curl http://localhost:8888/api/admin/staff/1 \
  -b cookies.txt  # 使用保存的Cookie

# 无Cookie访问（应返回401）
curl http://localhost:8888/api/admin/staff/1
# 预期响应: {"code":401,"message":"未登录或会话已过期"}
```

### 单元测试中间件

```go
package middleware

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestCookieAuthMiddleware_PublicPath(t *testing.T) {
    middleware := NewCookieAuthMiddleware([]string{
        "/api/admin/auth/captcha",
        "/api/admin/auth/login",
    })

    handler := middleware.Handle(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })

    // 测试公开路径
    req := httptest.NewRequest("GET", "/api/admin/auth/captcha", nil)
    w := httptest.NewRecorder()
    handler(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status 200, got %d", w.Code)
    }
}

func TestCookieAuthMiddleware_PrivatePath(t *testing.T) {
    middleware := NewCookieAuthMiddleware([]string{
        "/api/admin/auth/captcha",
        "/api/admin/auth/login",
    })

    handler := middleware.Handle(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })

    // 测试私有路径（无Cookie）
    req := httptest.NewRequest("GET", "/api/admin/staff/1", nil)
    w := httptest.NewRecorder()
    handler(w, req)

    if w.Code != http.StatusUnauthorized {
        t.Errorf("Expected status 401, got %d", w.Code)
    }
}
```

## 参考资料

- [Go-Zero 官方文档 - 中间件](https://go-zero.dev/docs/tutorials/http/middleware/middleware)
- [Go-Zero 官方文档 - 认证](https://go-zero.dev/docs/tutorials/http/jwt/jwt)
- Cookie安全最佳实践: HttpOnly, Secure, SameSite
- Session管理: Token哈希、过期时间、自动续期
