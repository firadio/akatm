package report

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPlatformReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 平台总览报表
func NewGetPlatformReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPlatformReportLogic {
	return &GetPlatformReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPlatformReportLogic) GetPlatformReport(req *types.GetPlatformReportReq) (resp *types.GetPlatformReportResp, err error) {
	// todo: add your logic here and delete this line

	return
}
