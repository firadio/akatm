package managerAuth

import (
	"net/http"

	"akatm/api/managerGateway/internal/logic/managerAuth"
	"akatm/api/managerGateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取图片验证码
func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := managerAuth.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
