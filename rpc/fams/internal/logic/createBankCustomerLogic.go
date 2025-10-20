package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateBankCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateBankCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateBankCustomerLogic {
	return &CreateBankCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 银行客户管理相关 ============
func (l *CreateBankCustomerLogic) CreateBankCustomer(in *famsRpc.CreateBankCustomerReq) (*famsRpc.CreateBankCustomerResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.CreateBankCustomerResp{}, nil
}
