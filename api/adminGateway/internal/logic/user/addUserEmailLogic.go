package user

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加用户邮箱
func NewAddUserEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserEmailLogic {
	return &AddUserEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserEmailLogic) AddUserEmail(req *types.AddUserEmailReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
