package auditWithdrawal

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWithdrawalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 提现审核详情
func NewGetWithdrawalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWithdrawalLogic {
	return &GetWithdrawalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWithdrawalLogic) GetWithdrawal(req *types.IdReq) (resp *types.GetWithdrawalResp, err error) {
	// todo: add your logic here and delete this line

	return
}
