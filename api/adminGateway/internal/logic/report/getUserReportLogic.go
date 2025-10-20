package report

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户报表
func NewGetUserReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserReportLogic {
	return &GetUserReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserReportLogic) GetUserReport(req *types.GetUserReportReq) (resp *types.GetUserReportResp, err error) {
	// todo: add your logic here and delete this line

	return
}
