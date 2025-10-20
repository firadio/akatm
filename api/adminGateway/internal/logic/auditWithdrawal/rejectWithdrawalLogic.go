package auditWithdrawal

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RejectWithdrawalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 审核拒绝
func NewRejectWithdrawalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RejectWithdrawalLogic {
	return &RejectWithdrawalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RejectWithdrawalLogic) RejectWithdrawal(req *types.RejectWithdrawalReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
