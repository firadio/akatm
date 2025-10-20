package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBankAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBankAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBankAccountLogic {
	return &GetBankAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取银行账户
func (l *GetBankAccountLogic) GetBankAccount(in *famsRpc.GetBankAccountReq) (*famsRpc.GetBankAccountResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.GetBankAccountResp{}, nil
}
