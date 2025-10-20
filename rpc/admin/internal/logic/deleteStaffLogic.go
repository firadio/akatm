package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStaffLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStaffLogic {
	return &DeleteStaffLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除员工
func (l *DeleteStaffLogic) DeleteStaff(in *adminRpc.DeleteStaffReq) (*adminRpc.DeleteStaffResp, error) {
	// todo: add your logic here and delete this line

	return &adminRpc.DeleteStaffResp{}, nil
}
