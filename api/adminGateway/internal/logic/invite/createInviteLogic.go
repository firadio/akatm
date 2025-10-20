package invite

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateInviteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 生成邀请链接
func NewCreateInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateInviteLogic {
	return &CreateInviteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateInviteLogic) CreateInvite(req *types.CreateInviteReq) (resp *types.CreateInviteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
