package auditDeposit

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RejectDepositLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 审核拒绝
func NewRejectDepositLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RejectDepositLogic {
	return &RejectDepositLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RejectDepositLogic) RejectDeposit(req *types.RejectDepositReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
