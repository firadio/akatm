package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListStaffsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListStaffsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListStaffsLogic {
	return &ListStaffsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取员工列表
func (l *ListStaffsLogic) ListStaffs(in *adminRpc.ListStaffsReq) (*adminRpc.ListStaffsResp, error) {
	// todo: add your logic here and delete this line

	return &adminRpc.ListStaffsResp{}, nil
}
