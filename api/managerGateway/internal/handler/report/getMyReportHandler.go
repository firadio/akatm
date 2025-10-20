package report

import (
	"net/http"

	"akatm/api/managerGateway/internal/logic/report"
	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 我的报表
func GetMyReportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMyReportReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := report.NewGetMyReportLogic(r.Context(), svcCtx)
		resp, err := l.GetMyReport(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
