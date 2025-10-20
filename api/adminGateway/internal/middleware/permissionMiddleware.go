package middleware

import (
	"context"
	"net/http"

	"akatm/rpc/admin/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuPermissionMiddleware struct {
	adminRpc admin.Admin
}

func NewMenuPermissionMiddleware(adminRpc admin.Admin) *MenuPermissionMiddleware {
	return &MenuPermissionMiddleware{
		adminRpc: adminRpc,
	}
}

// Handle 菜单权限验证中间件（仅用于前端菜单访问控制）
func (m *MenuPermissionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求头获取用户ID（JWT解析后应该设置）
		userID := r.Header.Get("X-User-ID")
		if userID == "" {
			http.Error(w, "missing user id", http.StatusUnauthorized)
			return
		}

		// 从请求头获取菜单路径
		menuPath := r.Header.Get("X-Menu-Path")
		if menuPath == "" {
			// 如果没有指定菜单路径，则跳过菜单权限检查
			next(w, r)
			return
		}

		// 检查用户是否有访问该菜单的权限
		hasMenuAccess, err := m.checkMenuAccess(userID, menuPath)
		if err != nil {
			logx.Errorf("菜单权限检查失败: %v", err)
			http.Error(w, "menu permission check failed", http.StatusInternalServerError)
			return
		}

		if !hasMenuAccess {
			http.Error(w, "insufficient menu access", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}

// checkMenuAccess 检查用户是否有访问指定菜单的权限
func (m *MenuPermissionMiddleware) checkMenuAccess(userID, menuPath string) (bool, error) {
	// 1. 获取员工的所有角色
	rpcReq := &admin.GetStaffRolesReq{
		StaffId:   1, // 这里需要从userID解析出staffID
		RequestId: "menu-check-" + userID,
	}

	rpcResp, err := m.adminRpc.GetStaffRoles(context.TODO(), rpcReq)
	if err != nil {
		return false, err
	}

	if rpcResp.Base.Code != 200 {
		return false, nil
	}

	// 2. 检查角色是否有访问该菜单的权限
	for _, role := range rpcResp.Roles {
		// 获取角色的所有菜单
		menuReq := &admin.GetRoleMenusReq{
			RoleId:    role.Id,
			RequestId: "menu-check-" + string(rune(role.Id)),
		}

		menuResp, err := m.adminRpc.GetRoleMenus(context.TODO(), menuReq)
		if err != nil {
			continue
		}

		// 检查菜单路径是否匹配
		for _, menu := range menuResp.Menus {
			if menu.Path == menuPath {
				return true, nil
			}
		}
	}

	return false, nil
}

// RequireMenuAccess 创建需要特定菜单访问权限的中间件
func RequireMenuAccess(adminRpc admin.Admin, menuPath string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 设置菜单路径到请求头
			r.Header.Set("X-Menu-Path", menuPath)

			// 调用菜单权限验证中间件
			middleware := NewMenuPermissionMiddleware(adminRpc)
			middleware.Handle(next)(w, r)
		}
	}
}
