# API 时间戳字段统一为 int64 修改总结

## 修改时间
2025-10-22

## 修改范围
- **文件数量**: 24个 .api 文件
- **影响字段**: 118个时间戳相关字段

## 主要修改内容

### 1. 时间戳字段类型统一
将所有时间相关字段从 `string` 类型修改为 `int64` 类型（毫秒级时间戳）

#### 修改的字段类型：
- `CreatedAt` - 创建时间
- `UpdatedAt` - 更新时间
- `DeletedAt` - 删除时间
- `StartTime` - 开始时间
- `EndTime` - 结束时间
- `ExpireTime` / `ExpiresAt` - 过期时间
- `LoginTime` / `LogoutTime` - 登录/登出时间
- `ApplicationTime` - 申请时间
- `AuditTime` - 审核时间
- `DepositTime` - 存款时间
- `WithdrawTime` - 提现时间
- `TransactionTime` - 交易时间
- `UploadTime` - 上传时间
- `SettlementTime` / `LastSettlementTime` - 结算时间
- `AccountOpeningDate` - 开户日期
- `EstimatedCompletionTime` - 预计完成时间
- `AverageProcessTime` - 平均处理时间
- `Time` / `Date` - 通用时间/日期字段

#### 修改示例：
```go
// 修改前
CreatedAt string `json:"createdAt"` // 创建时间
StartTime string `form:"startTime,optional"` // 开始时间

// 修改后
CreatedAt int64 `json:"createdAt"` // 创建时间（毫秒级时间戳）
StartTime int64 `form:"startTime,optional"` // 开始时间（毫秒级时间戳）
```

### 2. 删除 JWT 和中间件声明
从所有 @server 配置中删除以下内容：
- `jwt: Auth`
- `middleware: SignCheck, JwtAuth`

#### 修改示例：
```go
// 修改前
@server(
    prefix: /api/admin
    group: audit
    jwt: Auth
    middleware: SignCheck, JwtAuth
    tags: "日志审核"
)

// 修改后
@server(
    prefix: /api/admin
    group: audit
    tags: "日志审核"
)
```

### 3. 删除请求头说明
从所有接口的 description 中删除：
- `。请求头: Authorization, X-Timestamp, X-Sign`

#### 修改示例：
```go
// 修改前
description: "分页查询系统操作日志。请求头: Authorization, X-Timestamp, X-Sign"

// 修改后
description: "分页查询系统操作日志"
```

## 验证结果

### 清理完成度
- ✅ JWT/Middleware 残留: 0 个
- ✅ 请求头说明残留: 0 个
- ✅ String 类型时间字段残留: 0 个
- ✅ Int64 时间戳字段总数: 118 个

### 影响的文件列表
1. admin_auth.api
2. admin_data_country.api
3. admin_sys_menu.api
4. admin_sys_role.api
5. admin_sys_staff.api
6. audit.api
7. dashboard.api
8. fams_agent_earnings.api
9. fams_bank_account.api
10. fams_bank_account_application.api
11. fams_bank_customer.api
12. fams_bank_deposit.api
13. fams_bank_withdrawal.api
14. fams_fund_detail.api
15. fams_user_wallet.api
16. iam_user.api
17. iam_user_country.api
18. iam_user_credential.api
19. iam_user_email.api
20. iam_user_invite.api
21. iam_user_profile.api
22. iam_user_session.api
23. public.api
24. report.api
25. system.api

## 注意事项

### 时间戳格式说明
- **类型**: int64
- **单位**: 毫秒（milliseconds）
- **示例**: `1729587600000` (2024-10-22 12:00:00 UTC)

### 前后端对接注意点
1. **前端发送时间**: 需要将 Date 对象转换为毫秒级时间戳
   ```javascript
   const timestamp = new Date().getTime(); // 毫秒级
   ```

2. **前端接收时间**: 需要将时间戳转换为 Date 对象
   ```javascript
   const date = new Date(timestamp); // timestamp 是毫秒级
   ```

3. **后端存储**: 数据库层面仍使用 gorm.Model，保留 CreatedAt/UpdatedAt/DeletedAt 的 time.Time 类型

4. **API 层转换**: 在 handler 层进行类型转换
   ```go
   // API 响应
   CreatedAt: model.CreatedAt.UnixMilli() // time.Time -> int64 毫秒
   
   // API 请求解析
   createdAt := time.UnixMilli(req.CreatedAt) // int64 毫秒 -> time.Time
   ```

## 后续工作

### 已完成 ✅
- [x] 所有时间字段统一为 int64 类型
- [x] 删除所有 JWT 和中间件声明
- [x] 删除所有请求头说明
- [x] 验证修改完整性

### 待完成 ⏳
- [ ] 使用 goctl 重新生成 API 代码
- [ ] 更新业务逻辑层的时间戳处理
- [ ] 更新前端对接代码
- [ ] 更新 API 文档
- [ ] 测试时间戳字段的序列化/反序列化

## 回滚方案

如需回滚，可使用 git 版本控制：
```bash
git checkout -- api/admin/docs/*.api
```

---

**修改人员**: Claude  
**审核状态**: 待审核  
**版本**: v1.0
