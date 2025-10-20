package auditDeposit

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDepositLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 入账审核详情
func NewGetDepositLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepositLogic {
	return &GetDepositLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDepositLogic) GetDeposit(req *types.IdReq) (resp *types.GetDepositResp, err error) {
	// todo: add your logic here and delete this line

	return
}
