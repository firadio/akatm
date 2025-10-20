package withdrawal

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelWithdrawalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 取消提现
func NewCancelWithdrawalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelWithdrawalLogic {
	return &CancelWithdrawalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelWithdrawalLogic) CancelWithdrawal(req *types.CancelWithdrawalReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
