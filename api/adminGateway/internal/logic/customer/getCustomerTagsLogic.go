package customer

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomerTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取客户标签
func NewGetCustomerTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCustomerTagsLogic {
	return &GetCustomerTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCustomerTagsLogic) GetCustomerTags(req *types.IdReq) (resp *types.GetCustomerTagsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
