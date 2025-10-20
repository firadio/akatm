package dashboard

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAlertsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 告警信息
func NewGetAlertsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAlertsLogic {
	return &GetAlertsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAlertsLogic) GetAlerts(req *types.GetAlertsReq) (resp *types.GetAlertsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
