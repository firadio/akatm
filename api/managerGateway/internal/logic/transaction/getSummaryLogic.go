package transaction

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSummaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 资金总览
func NewGetSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSummaryLogic {
	return &GetSummaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSummaryLogic) GetSummary() (resp *types.GetSummaryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
