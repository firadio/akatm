package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBankCustomersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListBankCustomersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBankCustomersLogic {
	return &ListBankCustomersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取银行客户列表
func (l *ListBankCustomersLogic) ListBankCustomers(in *famsRpc.ListBankCustomersReq) (*famsRpc.ListBankCustomersResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.ListBankCustomersResp{}, nil
}
