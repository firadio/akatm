package menu

import (
	"net/http"

	"akatm/api/adminGateway/internal/logic/menu"
	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 创建菜单
func CreateMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewCreateMenuLogic(r.Context(), svcCtx)
		resp, err := l.CreateMenu(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
