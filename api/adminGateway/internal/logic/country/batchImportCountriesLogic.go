package country

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchImportCountriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量导入国家
func NewBatchImportCountriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchImportCountriesLogic {
	return &BatchImportCountriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchImportCountriesLogic) BatchImportCountries(req *types.BatchImportCountriesReq) (resp *types.BatchImportCountriesResp, err error) {
	// todo: add your logic here and delete this line

	return
}
