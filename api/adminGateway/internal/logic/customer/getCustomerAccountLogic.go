package customer

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomerAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 客户账户详情
func NewGetCustomerAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCustomerAccountLogic {
	return &GetCustomerAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCustomerAccountLogic) GetCustomerAccount(req *types.GetCustomerAccountReq) (resp *types.GetCustomerAccountResp, err error) {
	// todo: add your logic here and delete this line

	return
}
