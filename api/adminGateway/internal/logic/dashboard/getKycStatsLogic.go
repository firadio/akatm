package dashboard

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetKycStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// KYC统计
func NewGetKycStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetKycStatsLogic {
	return &GetKycStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetKycStatsLogic) GetKycStats(req *types.GetKycStatsReq) (resp *types.GetKycStatsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
