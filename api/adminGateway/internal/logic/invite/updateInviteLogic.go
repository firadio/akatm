package invite

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInviteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改邀请链接
func NewUpdateInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInviteLogic {
	return &UpdateInviteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInviteLogic) UpdateInvite(req *types.UpdateInviteReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
