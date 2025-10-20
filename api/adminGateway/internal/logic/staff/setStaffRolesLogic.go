package staff

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetStaffRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 设置角色
func NewSetStaffRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetStaffRolesLogic {
	return &SetStaffRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetStaffRolesLogic) SetStaffRoles(req *types.SetStaffRolesReq) (resp *types.BaseResp, err error) {
	// 批量分配角色给员工
	for _, roleId := range req.RoleIds {
		rpcReq := &admin.AssignRoleToStaffReq{
			StaffId:   req.Id,
			RoleId:    roleId,
			RequestId: "req-" + string(rune(req.Id)),
		}

		rpcResp, err := l.svcCtx.AdminRpc.AssignRoleToStaff(l.ctx, rpcReq)
		if err != nil {
			l.Errorf("调用RPC分配角色失败: %v", err)
			return &types.BaseResp{
				Code:    500,
				Message: "分配角色失败",
			}, nil
		}

		if rpcResp.Base.Code != 200 {
			return &types.BaseResp{
				Code:    int32(rpcResp.Base.Code),
				Message: rpcResp.Base.Message,
			}, nil
		}
	}

	return &types.BaseResp{
		Code:    200,
		Message: "分配成功",
	}, nil
}
