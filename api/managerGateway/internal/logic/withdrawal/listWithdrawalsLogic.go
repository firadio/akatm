package withdrawal

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWithdrawalsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 提现列表
func NewListWithdrawalsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWithdrawalsLogic {
	return &ListWithdrawalsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWithdrawalsLogic) ListWithdrawals(req *types.ListWithdrawalsReq) (resp *types.ListWithdrawalsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
