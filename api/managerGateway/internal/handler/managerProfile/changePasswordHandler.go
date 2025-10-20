package managerProfile

import (
	"net/http"

	"akatm/api/managerGateway/internal/logic/managerProfile"
	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 修改登录密码
func ChangePasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangePasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := managerProfile.NewChangePasswordLogic(r.Context(), svcCtx)
		resp, err := l.ChangePassword(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
