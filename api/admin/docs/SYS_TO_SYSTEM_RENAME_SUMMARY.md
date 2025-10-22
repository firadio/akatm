# sys → system 重命名修改总结

## 修改时间
2025-10-22

## 修改目的
统一将 `sys` 改为 `system`，使命名更加清晰和规范。

---

## 文件重命名

### 重命名的文件（4个）

| 原文件名 | 新文件名 | 说明 |
|---------|---------|------|
| `system.api` | `admin_system_config.api` | 系统配置管理 |
| `admin_sys_staff.api` | `admin_system_staff.api` | 员工管理 |
| `admin_sys_role.api` | `admin_system_role.api` | 角色管理 |
| `admin_sys_menu.api` | `admin_system_menu.api` | 菜单管理 |

---

## 文件内容修改

### 1. URL 路径修改

**修改前**:
```
prefix: /api/admin/sys
```

**修改后**:
```
prefix: /api/admin/system
```

**影响文件**:
- admin_system_staff.api
- admin_system_role.api  
- admin_system_menu.api
- admin_system_config.api

### 2. Group 名称修改

| 文件 | 原 group | 新 group |
|------|---------|----------|
| admin_system_staff.api | `sysStaff` / `adminSystemUser` | `adminSystemStaff` |
| admin_system_role.api | `sysRole` | `adminSystemRole` |
| admin_system_menu.api | `sysMenu` | `adminSystemMenu` |
| admin_system_config.api | `system` | `adminSystemConfig` |

### 3. 表名修改

数据库表名也需要相应修改（如果适用）：

| 原表名 | 新表名 |
|--------|--------|
| `admin_sys_staffs` | `admin_system_staffs` |
| `admin_sys_roles` | `admin_system_roles` |
| `admin_sys_menus` | `admin_system_menus` |
| `admin_sys_staff_sessions` | `admin_system_staff_sessions` |
| `admin_sys_configs` | `admin_system_configs` |

---

## URL 端点变化示例

### 员工管理
- **旧**: `POST /api/admin/sys/staff`
- **新**: `POST /api/admin/system/staff`

### 角色管理
- **旧**: `GET /api/admin/sys/role`
- **新**: `GET /api/admin/system/role`

### 菜单管理
- **旧**: `PUT /api/admin/sys/menu/:id`
- **新**: `PUT /api/admin/system/menu/:id`

### 系统配置
- **旧**: `GET /api/system/config`
- **新**: `GET /api/admin/system/config`

---

## 文档更新

已更新 `docs/api-files-inventory.md`:

1. **总文件数**: 20 → 24
2. **Admin RPC 文件数**: 5 → 7
3. **IAM RPC 文件数**: 6 → 7
4. **分类统计**: `sys` → `system`
5. **Group 命名示例**: `adminSysStaff` → `adminSystemStaff`

---

## 验证结果

### 文件清单
```
admin_system_config.api     ✅
admin_system_menu.api       ✅
admin_system_role.api       ✅
admin_system_staff.api      ✅
```

### URL 前缀
所有文件的 prefix 都已更新为: `/api/admin/system` ✅

### Group 命名
所有 group 都遵循 `adminSystem{Resource}` 格式 ✅

---

## 影响范围

### 前端影响
需要更新以下 API 调用路径：
- `/api/admin/sys/*` → `/api/admin/system/*`
- `/api/system/*` → `/api/admin/system/*`

### 后端影响
1. 重新运行 `goctl` 生成代码
2. 更新路由注册
3. 更新业务逻辑中的引用
4. 数据库表名迁移（如需要）

### 数据库迁移
如果表名需要修改，需要执行数据库迁移脚本：
```sql
-- 示例（仅供参考，实际需根据情况调整）
RENAME TABLE admin_sys_staffs TO admin_system_staffs;
RENAME TABLE admin_sys_roles TO admin_system_roles;
RENAME TABLE admin_sys_menus TO admin_system_menus;
RENAME TABLE admin_sys_configs TO admin_system_configs;
```

---

## 后续工作

### 必须完成 ⚠️
- [ ] 使用 `goctl api go` 重新生成 API 代码
- [ ] 更新前端 API 调用路径
- [ ] 更新 API 文档/Swagger
- [ ] 执行数据库迁移（如表名有变化）
- [ ] 进行完整的回归测试

### 建议完成 💡
- [ ] 更新 Postman 集合
- [ ] 更新 API 使用文档
- [ ] 通知前端团队进行相应调整
- [ ] 更新部署脚本（如有硬编码路径）

---

## 回滚方案

如需回滚，可使用 git：
```bash
# 回滚 API 文件
git checkout HEAD -- api/admin/docs/admin_system_*.api

# 回滚文档
git checkout HEAD -- docs/api-files-inventory.md
```

或手动回滚：
```bash
cd api/admin/docs
mv admin_system_config.api system.api
mv admin_system_staff.api admin_sys_staff.api
mv admin_system_role.api admin_sys_role.api
mv admin_system_menu.api admin_sys_menu.api

# 然后手动修改文件内的 prefix 和 group
```

---

**执行人**: Claude  
**审核状态**: 待审核  
**版本**: v1.0  
**更新日期**: 2025-10-22
