package managerProfile

import (
	"net/http"

	"akatm/api/managerGateway/internal/logic/managerProfile"
	"akatm/api/managerGateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取费率设置
func GetFeeSettingsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := managerProfile.NewGetFeeSettingsLogic(r.Context(), svcCtx)
		resp, err := l.GetFeeSettings()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
