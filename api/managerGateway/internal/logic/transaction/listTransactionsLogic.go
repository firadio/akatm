package transaction

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTransactionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 流水列表
func NewListTransactionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTransactionsLogic {
	return &ListTransactionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListTransactionsLogic) ListTransactions(req *types.ListTransactionsReq) (resp *types.ListTransactionsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
