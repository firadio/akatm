package customer

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCustomerTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 为客户添加标签
func NewAddCustomerTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCustomerTagsLogic {
	return &AddCustomerTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCustomerTagsLogic) AddCustomerTags(req *types.AddCustomerTagsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
