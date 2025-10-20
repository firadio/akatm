package managerAuth

import (
	"net/http"

	"akatm/api/managerGateway/internal/logic/managerAuth"
	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 邀请注册
func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ManagerRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := managerAuth.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
