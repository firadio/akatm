package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBankCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBankCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBankCustomerLogic {
	return &UpdateBankCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新银行客户
func (l *UpdateBankCustomerLogic) UpdateBankCustomer(in *famsRpc.UpdateBankCustomerReq) (*famsRpc.UpdateBankCustomerResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.UpdateBankCustomerResp{}, nil
}
