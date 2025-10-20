package auditAccountApplication

import (
	"net/http"

	"akatm/api/adminGateway/internal/logic/auditAccountApplication"
	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 批量审核通过
func BatchApproveAccountApplicationsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchApproveAccountApplicationsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auditAccountApplication.NewBatchApproveAccountApplicationsLogic(r.Context(), svcCtx)
		resp, err := l.BatchApproveAccountApplications(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
