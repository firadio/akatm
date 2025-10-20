package report

import (
	"net/http"

	"akatm/api/adminGateway/internal/logic/report"
	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 总代报表
func GetSuperAgentReportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSuperAgentReportReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := report.NewGetSuperAgentReportLogic(r.Context(), svcCtx)
		resp, err := l.GetSuperAgentReport(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
