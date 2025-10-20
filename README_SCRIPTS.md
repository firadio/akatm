# AKATM 项目脚本使用说明

本项目提供了多个自动化脚本，用于代码生成和项目管理。所有脚本都支持单个服务生成，提高开发效率。

## 📋 脚本概览

| 脚本名称     | 功能描述          | 支持参数                         |
| ------------ | ----------------- | -------------------------------- |
| `gen_api.sh` | 生成 API 网关代码 | admin, manager, all              |
| `gen_doc.sh` | 生成 Swagger 文档 | admin, manager, all              |
| `gen_rpc.sh` | 生成 RPC 服务代码 | iam, admin, fams, mail, all      |
| `gen_orm.sh` | 生成 ORM 模型代码 | iam, admin, fams, all + 环境参数 |
| `gen_all.sh` | 一键生成所有代码  | 无参数                           |

## 🚀 快速开始

### 1. 生成所有代码

```bash
# 一键生成所有 API、RPC、ORM 代码
./gen_all.sh
```

### 2. 分步生成

```bash
# 1. 生成 API 网关
./gen_api.sh all

# 2. 生成 Swagger 文档
./gen_doc.sh all

# 3. 生成 RPC 服务
./gen_rpc.sh all

# 4. 生成 ORM 模型
./gen_orm.sh all
```

## 📖 详细使用说明

### 🔧 gen_api.sh - API 网关生成

生成 go-zero API 网关代码，包括 Handler、Logic、Config 等。

**使用方法：**

```bash
# 生成所有 API 网关
./gen_api.sh all

# 生成指定服务网关
./gen_api.sh admin      # 生成管理端网关
./gen_api.sh manager    # 生成经理端网关

# 查看帮助
./gen_api.sh
```

**输出目录：**

- 管理端：`api/adminGateway/`
- 经理端：`api/managerGateway/`

**生成内容：**

- HTTP 路由和处理器
- 业务逻辑层
- 配置文件
- 中间件集成

---

### 📚 gen_doc.sh - Swagger 文档生成

生成标准的 OpenAPI 3.0 格式 Swagger 文档。

**使用方法：**

```bash
# 生成所有 Swagger 文档
./gen_doc.sh all

# 生成指定服务文档
./gen_doc.sh admin      # 生成管理端 Swagger
./gen_doc.sh manager    # 生成经理端 Swagger

# 查看帮助
./gen_doc.sh
```

**输出文件：**

- 管理端：`docs/swagger/admin-swagger.yaml`
- 经理端：`docs/swagger/manager-swagger.yaml`

**查看文档：**

- 在线工具：https://editor.swagger.io/
- 本地工具：swagger-ui-serve、swagger-codegen

---

### 🔗 gen_rpc.sh - RPC 服务生成

生成 go-zero RPC 服务代码，包括服务端和客户端。

**使用方法：**

```bash
# 生成所有 RPC 服务
./gen_rpc.sh all

# 生成指定 RPC 服务
./gen_rpc.sh iam        # 生成 IAM 服务
./gen_rpc.sh admin      # 生成 Admin 服务
./gen_rpc.sh fams       # 生成 FAMS 服务
./gen_rpc.sh mail       # 生成 Mail 服务

# 查看帮助
./gen_rpc.sh
```

**服务说明：**

- `iam` - 身份认证服务（用户管理、邀请、注册）
- `admin` - 管理端服务（RBAC、人员管理）
- `fams` - 金融账户服务（银行账户、钱包、交易）
- `mail` - 邮件服务（验证码、通知）

**输出目录：**

- `rpc/{service}/internal/server/` - 服务端代码
- `rpc/{service}/client/` - 客户端代码

---

### 🗄️ gen_orm.sh - ORM 模型生成

生成 GORM 模型代码，支持多环境配置。

**使用方法：**

```bash
# 生成所有 ORM 模型（默认环境）
./gen_orm.sh all

# 生成指定服务 ORM
./gen_orm.sh iam        # 生成 IAM ORM
./gen_orm.sh admin      # 生成 Admin ORM
./gen_orm.sh fams       # 生成 FAMS ORM

# 指定环境生成
./gen_orm.sh iam dev    # 使用 dev 环境配置
./gen_orm.sh admin test # 使用 test 环境配置
./gen_orm.sh fams pro   # 使用 pro 环境配置

# 查看帮助
./gen_orm.sh
```

**环境配置：**

- 默认：使用 `.env` 文件
- `dev`：使用 `.env.dev` 文件
- `test`：使用 `.env.test` 文件
- `pro`：使用 `.env.pro` 文件

**配置文件位置：**

```
rpc/{service}/orm/gen/
├── .env          # 默认环境
├── .env.dev      # 开发环境
├── .env.test     # 测试环境
└── .env.pro      # 生产环境
```

**配置格式：**

```bash
# 数据库连接
DSN=user:password@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local

# 表前缀
TABLE_PREFIX=akatm_
```

---

### 🎯 gen_all.sh - 一键生成

一键执行所有代码生成脚本。

**使用方法：**

```bash
# 生成所有代码
./gen_all.sh
```

**执行顺序：**

1. 生成 API 网关 (`gen_api.sh all`)
2. 生成 Swagger 文档 (`gen_doc.sh all`)
3. 生成 RPC 服务 (`gen_rpc.sh all`)
4. 生成 ORM 模型 (`gen_orm.sh all`)

## ⚙️ 环境要求

### 必需工具

- **goctl**: go-zero 代码生成工具

  ```bash
  go install github.com/zeromicro/go-zero/tools/goctl@latest
  ```

- **GORM**: ORM 代码生成工具
  ```bash
  go install gorm.io/gorm/cmd/gorm-gen@latest
  ```

### 环境变量

确保设置了正确的环境变量：

```bash
export ENV=dev  # 可选：dev, test, pro
```

## 🔧 配置说明

### API 网关配置

每个 API 网关都有对应的配置文件：

- `api/adminGateway/etc/admin.yaml`
- `api/managerGateway/etc/manager.yaml`

### RPC 服务配置

每个 RPC 服务都有对应的配置文件：

- `rpc/iam/etc/iam.yaml`
- `rpc/admin/etc/admin.yaml`
- `rpc/fams/etc/fams.yaml`
- `rpc/mail/etc/mail.yaml`

### 数据库配置

ORM 生成需要配置数据库连接信息：

```bash
# 在 rpc/{service}/orm/gen/.env* 文件中配置
DSN=user:password@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local
TABLE_PREFIX=akatm_
```

## 🚨 常见问题

### 1. goctl 命令未找到

```bash
# 安装 goctl
go install github.com/zeromicro/go-zero/tools/goctl@latest

# 验证安装
goctl --version
```

### 2. GORM 生成失败

```bash
# 安装 gorm-gen
go install gorm.io/gorm/cmd/gorm-gen@latest

# 检查数据库连接
# 确保 .env 文件中的 DSN 配置正确
```

### 3. 类型重复定义错误

- 检查 `.api` 文件中是否有重复的类型定义
- 确保每个模块使用唯一的类型名称
- 使用模块前缀避免冲突（如 `AdminRoleInfo`、`ManagerRoleInfo`）

### 4. 权限问题

```bash
# 给脚本添加执行权限
chmod +x *.sh
```

## 📁 项目结构

```
akatm/
├── api/                    # API 网关
│   ├── admin/              # 管理端 API 定义
│   ├── manager/            # 经理端 API 定义
│   ├── adminGateway/       # 管理端网关代码
│   └── managerGateway/     # 经理端网关代码
├── rpc/                    # RPC 服务
│   ├── iam/               # 身份认证服务
│   ├── admin/             # 管理端服务
│   ├── fams/              # 金融账户服务
│   └── mail/              # 邮件服务
├── common/                # 公共库
├── docs/                  # 文档
│   └── swagger/          # Swagger 文档
├── gen_*.sh              # 生成脚本
└── README_SCRIPTS.md     # 本文档
```

## 🎯 最佳实践

### 1. 开发流程

```bash
# 1. 修改 .api 文件
# 2. 重新生成 API 网关
./gen_api.sh admin

# 3. 重新生成文档
./gen_doc.sh admin

# 4. 修改 .proto 文件
# 5. 重新生成 RPC 服务
./gen_rpc.sh iam

# 6. 修改数据表定义
# 7. 重新生成 ORM
./gen_orm.sh iam dev
```

### 2. 环境管理

- 开发环境使用 `dev` 配置
- 测试环境使用 `test` 配置
- 生产环境使用 `pro` 配置

### 3. 代码生成原则

- 只生成正在开发的服务，避免不必要的代码变更
- 生成前确保配置文件正确
- 生成后检查代码质量

## 📞 支持

如果遇到问题，请检查：

1. 工具是否正确安装
2. 配置文件是否正确
3. 网络连接是否正常
4. 权限是否足够

---

**版本**: 1.0.0  
**更新时间**: 2024-10-18  
**适用项目**: AKATM 金融账户管理系统
