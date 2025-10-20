package managerProfile

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeeSettingsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取费率设置
func NewGetFeeSettingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeeSettingsLogic {
	return &GetFeeSettingsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFeeSettingsLogic) GetFeeSettings() (resp *types.GetFeeSettingsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
