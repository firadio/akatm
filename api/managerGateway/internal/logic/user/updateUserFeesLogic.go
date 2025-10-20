package user

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserFeesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户手续费
func NewUpdateUserFeesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserFeesLogic {
	return &UpdateUserFeesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserFeesLogic) UpdateUserFees(req *types.UpdateUserFeesReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
