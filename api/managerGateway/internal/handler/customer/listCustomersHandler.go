package customer

import (
	"net/http"

	"akatm/api/managerGateway/internal/logic/customer"
	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 客户列表
func ListCustomersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListCustomersReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := customer.NewListCustomersLogic(r.Context(), svcCtx)
		resp, err := l.ListCustomers(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
