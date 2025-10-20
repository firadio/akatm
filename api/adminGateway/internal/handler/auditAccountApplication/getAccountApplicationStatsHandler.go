package auditAccountApplication

import (
	"net/http"

	"akatm/api/adminGateway/internal/logic/auditAccountApplication"
	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 申请统计
func GetAccountApplicationStatsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAccountApplicationStatsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auditAccountApplication.NewGetAccountApplicationStatsLogic(r.Context(), svcCtx)
		resp, err := l.GetAccountApplicationStats(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
