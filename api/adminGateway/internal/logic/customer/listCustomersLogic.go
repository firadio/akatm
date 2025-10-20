package customer

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCustomersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 客户列表
func NewListCustomersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCustomersLogic {
	return &ListCustomersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCustomersLogic) ListCustomers(req *types.ListCustomersReq) (resp *types.ListCustomersResp, err error) {
	// todo: add your logic here and delete this line

	return
}
