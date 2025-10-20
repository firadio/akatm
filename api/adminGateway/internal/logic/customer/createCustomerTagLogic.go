package customer

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCustomerTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建客户标签
func NewCreateCustomerTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCustomerTagLogic {
	return &CreateCustomerTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCustomerTagLogic) CreateCustomerTag(req *types.CreateCustomerTagReq) (resp *types.CreateCustomerTagResp, err error) {
	// todo: add your logic here and delete this line

	return
}
