package managerAuth

import (
	"net/http"

	"akatm/api/managerGateway/internal/logic/managerAuth"
	"akatm/api/managerGateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 当前用户信息
func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := managerAuth.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
