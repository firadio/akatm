package role

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分配菜单
func NewAssignMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignMenusLogic {
	return &AssignMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignMenusLogic) AssignMenus(req *types.AssignMenusReq) (resp *types.BaseResp, err error) {
	// 批量分配菜单给角色
	for _, menuId := range req.MenuIds {
		rpcReq := &admin.AssignMenuToRoleReq{
			RoleId:    req.Id,
			MenuId:    menuId,
			RequestId: "req-" + string(rune(req.Id)),
		}

		rpcResp, err := l.svcCtx.AdminRpc.AssignMenuToRole(l.ctx, rpcReq)
		if err != nil {
			l.Errorf("调用RPC分配菜单失败: %v", err)
			return &types.BaseResp{
				Code:    500,
				Message: "分配菜单失败",
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
