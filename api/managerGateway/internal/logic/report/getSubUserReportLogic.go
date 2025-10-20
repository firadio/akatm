package report

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubUserReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 下级用户报表
func NewGetSubUserReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubUserReportLogic {
	return &GetSubUserReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubUserReportLogic) GetSubUserReport(req *types.GetSubUserReportReq) (resp *types.GetSubUserReportResp, err error) {
	// todo: add your logic here and delete this line

	return
}
