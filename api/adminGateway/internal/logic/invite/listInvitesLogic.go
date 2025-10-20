package invite

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListInvitesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 邀请链接列表
func NewListInvitesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListInvitesLogic {
	return &ListInvitesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListInvitesLogic) ListInvites(req *types.ListInvitesReq) (resp *types.ListInvitesResp, err error) {
	// todo: add your logic here and delete this line

	return
}
