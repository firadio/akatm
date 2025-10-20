package report

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMyReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 我的报表
func NewGetMyReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMyReportLogic {
	return &GetMyReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMyReportLogic) GetMyReport(req *types.GetMyReportReq) (resp *types.GetMyReportResp, err error) {
	// todo: add your logic here and delete this line

	return
}
