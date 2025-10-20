package country

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCountrySortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新国家排序
func NewUpdateCountrySortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCountrySortLogic {
	return &UpdateCountrySortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCountrySortLogic) UpdateCountrySort(req *types.UpdateCountrySortReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
