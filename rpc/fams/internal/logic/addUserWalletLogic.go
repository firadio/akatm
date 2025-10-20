package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserWalletLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserWalletLogic {
	return &AddUserWalletLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 增加资金
func (l *AddUserWalletLogic) AddUserWallet(in *famsRpc.AddUserWalletReq) (*famsRpc.AddUserWalletResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.AddUserWalletResp{}, nil
}
