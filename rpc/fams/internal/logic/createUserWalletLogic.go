package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserWalletLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserWalletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserWalletLogic {
	return &CreateUserWalletLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 用户钱包管理相关 ============
func (l *CreateUserWalletLogic) CreateUserWallet(in *famsRpc.CreateUserWalletReq) (*famsRpc.CreateUserWalletResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.CreateUserWalletResp{}, nil
}
