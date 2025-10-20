package auditWithdrawal

import (
	"net/http"

	"akatm/api/adminGateway/internal/logic/auditWithdrawal"
	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 批量审核拒绝
func BatchRejectWithdrawalsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchRejectWithdrawalsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auditWithdrawal.NewBatchRejectWithdrawalsLogic(r.Context(), svcCtx)
		resp, err := l.BatchRejectWithdrawals(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
