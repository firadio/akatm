package system

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取配置历史
func NewGetConfigHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigHistoryLogic {
	return &GetConfigHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigHistoryLogic) GetConfigHistory(req *types.GetConfigHistoryReq) (resp *types.GetConfigHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
