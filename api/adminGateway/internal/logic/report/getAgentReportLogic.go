package report

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAgentReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 代理报表
func NewGetAgentReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAgentReportLogic {
	return &GetAgentReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAgentReportLogic) GetAgentReport(req *types.GetAgentReportReq) (resp *types.GetAgentReportResp, err error) {
	// todo: add your logic here and delete this line

	return
}
