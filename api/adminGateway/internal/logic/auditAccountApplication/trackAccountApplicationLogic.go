package auditAccountApplication

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TrackAccountApplicationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 申请流程跟踪
func NewTrackAccountApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TrackAccountApplicationLogic {
	return &TrackAccountApplicationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TrackAccountApplicationLogic) TrackAccountApplication(req *types.IdReq) (resp *types.TrackAccountApplicationResp, err error) {
	// todo: add your logic here and delete this line

	return
}
