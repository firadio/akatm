package customer

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCustomerTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新客户标签
func NewUpdateCustomerTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomerTagLogic {
	return &UpdateCustomerTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCustomerTagLogic) UpdateCustomerTag(req *types.UpdateCustomerTagReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
