package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBankAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBankAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBankAccountLogic {
	return &CreateBankAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 银行账户管理相关 ============
func (l *CreateBankAccountLogic) CreateBankAccount(in *famsRpc.CreateBankAccountReq) (*famsRpc.CreateBankAccountResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.CreateBankAccountResp{}, nil
}
