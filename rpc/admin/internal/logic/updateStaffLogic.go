package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStaffLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStaffLogic {
	return &UpdateStaffLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新员工
func (l *UpdateStaffLogic) UpdateStaff(in *adminRpc.UpdateStaffReq) (*adminRpc.UpdateStaffResp, error) {
	// todo: add your logic here and delete this line

	return &adminRpc.UpdateStaffResp{}, nil
}
