package user

import (
	"net/http"

	"akatm/api/adminGateway/internal/logic/user"
	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 添加用户邮箱
func AddUserEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddUserEmailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewAddUserEmailLogic(r.Context(), svcCtx)
		resp, err := l.AddUserEmail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
