package auditWithdrawal

import (
	"net/http"

	"akatm/api/adminGateway/internal/logic/auditWithdrawal"
	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 审核通过
func ApproveWithdrawalHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ApproveWithdrawalReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auditWithdrawal.NewApproveWithdrawalLogic(r.Context(), svcCtx)
		resp, err := l.ApproveWithdrawal(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
