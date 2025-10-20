package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBankAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBankAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBankAccountLogic {
	return &DeleteBankAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除银行账户
func (l *DeleteBankAccountLogic) DeleteBankAccount(in *famsRpc.DeleteBankAccountReq) (*famsRpc.DeleteBankAccountResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.DeleteBankAccountResp{}, nil
}
