package account

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAccountStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新账户状态
func NewUpdateAccountStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAccountStatusLogic {
	return &UpdateAccountStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAccountStatusLogic) UpdateAccountStatus(req *types.UpdateAccountStatusReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
