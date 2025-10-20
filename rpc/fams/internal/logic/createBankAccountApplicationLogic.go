package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBankAccountApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBankAccountApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBankAccountApplicationLogic {
	return &CreateBankAccountApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 银行账户申请相关 ============
func (l *CreateBankAccountApplicationLogic) CreateBankAccountApplication(in *famsRpc.CreateBankAccountApplicationReq) (*famsRpc.CreateBankAccountApplicationResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.CreateBankAccountApplicationResp{}, nil
}
