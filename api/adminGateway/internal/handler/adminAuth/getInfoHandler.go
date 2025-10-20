package adminAuth

import (
	"net/http"

	"akatm/api/adminGateway/internal/logic/adminAuth"
	"akatm/api/adminGateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取当前管理员信息
func GetInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := adminAuth.NewGetInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
