package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBankCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBankCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBankCustomerLogic {
	return &GetBankCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取银行客户
func (l *GetBankCustomerLogic) GetBankCustomer(in *famsRpc.GetBankCustomerReq) (*famsRpc.GetBankCustomerResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.GetBankCustomerResp{}, nil
}
