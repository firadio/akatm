package dashboard

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 账户统计
func NewGetAccountStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountStatsLogic {
	return &GetAccountStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAccountStatsLogic) GetAccountStats(req *types.GetAccountStatsReq) (resp *types.GetAccountStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
