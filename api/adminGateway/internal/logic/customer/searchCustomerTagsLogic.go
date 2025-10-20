package customer

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCustomerTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 搜索可用标签
func NewSearchCustomerTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCustomerTagsLogic {
	return &SearchCustomerTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchCustomerTagsLogic) SearchCustomerTags(req *types.SearchCustomerTagsReq) (resp *types.SearchCustomerTagsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
