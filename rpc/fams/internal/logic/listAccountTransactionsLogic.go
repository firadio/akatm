package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAccountTransactionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListAccountTransactionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAccountTransactionsLogic {
	return &ListAccountTransactionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取账户流水
func (l *ListAccountTransactionsLogic) ListAccountTransactions(in *famsRpc.ListAccountTransactionsReq) (*famsRpc.ListAccountTransactionsResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.ListAccountTransactionsResp{}, nil
}
