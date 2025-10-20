package invite

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInviteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 邀请详情
func NewGetInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInviteLogic {
	return &GetInviteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInviteLogic) GetInvite(req *types.IdReq) (resp *types.GetInviteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
