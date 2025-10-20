package logic

import (
	"context"

	"akatm/rpc/iam/internal/svc"
	"akatm/rpc/iam/pb/iamRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListInvitesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListInvitesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListInvitesLogic {
	return &ListInvitesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取邀请列表
func (l *ListInvitesLogic) ListInvites(in *iamRpc.ListInvitesReq) (*iamRpc.ListInvitesResp, error) {
	// todo: add your logic here and delete this line

	return &iamRpc.ListInvitesResp{}, nil
}
