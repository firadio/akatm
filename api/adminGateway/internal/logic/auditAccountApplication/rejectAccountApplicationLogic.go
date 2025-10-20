package auditAccountApplication

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RejectAccountApplicationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 审核拒绝
func NewRejectAccountApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RejectAccountApplicationLogic {
	return &RejectAccountApplicationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RejectAccountApplicationLogic) RejectAccountApplication(req *types.RejectAccountApplicationReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
