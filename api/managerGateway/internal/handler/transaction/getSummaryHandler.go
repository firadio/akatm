package transaction

import (
	"net/http"

	"akatm/api/managerGateway/internal/logic/transaction"
	"akatm/api/managerGateway/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 资金总览
func GetSummaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := transaction.NewGetSummaryLogic(r.Context(), svcCtx)
		resp, err := l.GetSummary()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
