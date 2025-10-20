package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUserWalletsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserWalletsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserWalletsLogic {
	return &ListUserWalletsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户钱包列表
func (l *ListUserWalletsLogic) ListUserWallets(in *famsRpc.ListUserWalletsReq) (*famsRpc.ListUserWalletsResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.ListUserWalletsResp{}, nil
}
