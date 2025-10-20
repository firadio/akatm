package user

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 下级用户列表
func NewGetSubUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubUsersLogic {
	return &GetSubUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubUsersLogic) GetSubUsers(req *types.GetSubUsersReq) (resp *types.ListUsersResp, err error) {
	// todo: add your logic here and delete this line

	return
}
