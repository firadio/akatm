package auditDeposit

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListDepositsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 入账审核列表
func NewListDepositsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListDepositsLogic {
	return &ListDepositsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListDepositsLogic) ListDeposits(req *types.ListDepositsReq) (resp *types.ListDepositsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
