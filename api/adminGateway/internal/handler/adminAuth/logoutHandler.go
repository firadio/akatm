package adminAuth

import (
	"net/http"

	"akatm/api/adminGateway/internal/logic/adminAuth"
	"akatm/api/adminGateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 管理员登出
func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := adminAuth.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
