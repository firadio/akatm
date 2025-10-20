package report

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomersReportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 客户报表
func NewGetCustomersReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCustomersReportLogic {
	return &GetCustomersReportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCustomersReportLogic) GetCustomersReport(req *types.GetCustomersReportReq) (resp *types.GetCustomersReportResp, err error) {
	// todo: add your logic here and delete this line

	return
}
