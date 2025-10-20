package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBankAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBankAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBankAccountLogic {
	return &UpdateBankAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新银行账户
func (l *UpdateBankAccountLogic) UpdateBankAccount(in *famsRpc.UpdateBankAccountReq) (*famsRpc.UpdateBankAccountResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.UpdateBankAccountResp{}, nil
}
