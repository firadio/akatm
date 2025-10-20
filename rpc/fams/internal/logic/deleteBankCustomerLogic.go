package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBankCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteBankCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBankCustomerLogic {
	return &DeleteBankCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除银行客户
func (l *DeleteBankCustomerLogic) DeleteBankCustomer(in *famsRpc.DeleteBankCustomerReq) (*famsRpc.DeleteBankCustomerResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.DeleteBankCustomerResp{}, nil
}
