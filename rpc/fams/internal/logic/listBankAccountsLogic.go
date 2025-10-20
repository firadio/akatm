package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBankAccountsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListBankAccountsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBankAccountsLogic {
	return &ListBankAccountsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取银行账户列表
func (l *ListBankAccountsLogic) ListBankAccounts(in *famsRpc.ListBankAccountsReq) (*famsRpc.ListBankAccountsResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.ListBankAccountsResp{}, nil
}
