package customer

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCustomerTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 客户标签列表
func NewListCustomerTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCustomerTagsLogic {
	return &ListCustomerTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCustomerTagsLogic) ListCustomerTags(req *types.ListCustomerTagsReq) (resp *types.ListCustomerTagsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
