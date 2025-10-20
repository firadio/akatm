package auditDeposit

import (
	"net/http"

	"akatm/api/adminGateway/internal/logic/auditDeposit"
	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 入账审核列表
func ListDepositsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListDepositsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auditDeposit.NewListDepositsLogic(r.Context(), svcCtx)
		resp, err := l.ListDeposits(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
