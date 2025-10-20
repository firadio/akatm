package country

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetImportTemplateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取导入模板
func NewGetImportTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetImportTemplateLogic {
	return &GetImportTemplateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetImportTemplateLogic) GetImportTemplate(req *types.GetImportTemplateReq) (resp *types.GetImportTemplateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
