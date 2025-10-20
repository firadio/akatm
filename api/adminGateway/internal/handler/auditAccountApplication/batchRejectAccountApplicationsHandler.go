package auditAccountApplication

import (
	"net/http"

	"akatm/api/adminGateway/internal/logic/auditAccountApplication"
	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 批量审核拒绝
func BatchRejectAccountApplicationsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BatchRejectAccountApplicationsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auditAccountApplication.NewBatchRejectAccountApplicationsLogic(r.Context(), svcCtx)
		resp, err := l.BatchRejectAccountApplications(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
