package dashboard

import (
	"net/http"

	"akatm/api/adminGateway/internal/logic/dashboard"
	"akatm/api/adminGateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 实时监控
func GetRealtimeMonitorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := dashboard.NewGetRealtimeMonitorLogic(r.Context(), svcCtx)
		resp, err := l.GetRealtimeMonitor()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
