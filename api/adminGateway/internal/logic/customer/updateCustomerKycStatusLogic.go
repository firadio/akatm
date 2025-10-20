package customer

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCustomerKycStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新KYC状态
func NewUpdateCustomerKycStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCustomerKycStatusLogic {
	return &UpdateCustomerKycStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCustomerKycStatusLogic) UpdateCustomerKycStatus(req *types.UpdateCustomerKycStatusReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
