package report

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSuperAgentReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 总代报表
func NewGetSuperAgentReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSuperAgentReportLogic {
	return &GetSuperAgentReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSuperAgentReportLogic) GetSuperAgentReport(req *types.GetSuperAgentReportReq) (resp *types.GetSuperAgentReportResp, err error) {
	// todo: add your logic here and delete this line

	return
}
