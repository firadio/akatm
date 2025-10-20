package wallet

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWalletsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 钱包列表
func NewListWalletsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWalletsLogic {
	return &ListWalletsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWalletsLogic) ListWallets(req *types.ListWalletsReq) (resp *types.ListWalletsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
