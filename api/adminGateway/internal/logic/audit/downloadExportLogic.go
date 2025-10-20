package audit

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadExportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 下载导出文件
func NewDownloadExportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadExportLogic {
	return &DownloadExportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadExportLogic) DownloadExport(req *types.GetExportTaskStatusReq) (resp *types.DownloadExportResp, err error) {
	// todo: add your logic here and delete this line

	return
}
