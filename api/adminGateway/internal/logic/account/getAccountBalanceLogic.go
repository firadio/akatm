package account

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountBalanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 账户余额
func NewGetAccountBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountBalanceLogic {
	return &GetAccountBalanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAccountBalanceLogic) GetAccountBalance(req *types.IdReq) (resp *types.GetAccountBalanceResp, err error) {
	// todo: add your logic here and delete this line

	return
}
