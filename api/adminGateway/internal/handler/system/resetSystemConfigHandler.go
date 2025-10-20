package system

import (
	"net/http"

	"akatm/api/adminGateway/internal/logic/system"
	"akatm/api/adminGateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 重置系统配置
func ResetSystemConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := system.NewResetSystemConfigLogic(r.Context(), svcCtx)
		resp, err := l.ResetSystemConfig()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
