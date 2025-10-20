package auditAccountApplication

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApproveAccountApplicationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 审核通过
func NewApproveAccountApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApproveAccountApplicationLogic {
	return &ApproveAccountApplicationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApproveAccountApplicationLogic) ApproveAccountApplication(req *types.ApproveAccountApplicationReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
