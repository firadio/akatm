package account

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAccountsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 账户列表
func NewListAccountsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAccountsLogic {
	return &ListAccountsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAccountsLogic) ListAccounts(req *types.ListAccountsReq) (resp *types.ListAccountsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
