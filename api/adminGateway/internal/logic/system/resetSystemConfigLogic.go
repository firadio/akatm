package system

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetSystemConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 重置系统配置
func NewResetSystemConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetSystemConfigLogic {
	return &ResetSystemConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetSystemConfigLogic) ResetSystemConfig() (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
