package customer

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveCustomerTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 移除客户标签
func NewRemoveCustomerTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveCustomerTagsLogic {
	return &RemoveCustomerTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveCustomerTagsLogic) RemoveCustomerTags(req *types.RemoveCustomerTagsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
