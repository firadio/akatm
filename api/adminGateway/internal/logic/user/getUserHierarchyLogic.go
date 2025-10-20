package user

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserHierarchyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户层级关系
func NewGetUserHierarchyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserHierarchyLogic {
	return &GetUserHierarchyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserHierarchyLogic) GetUserHierarchy(req *types.IdReq) (resp *types.GetUserHierarchyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
