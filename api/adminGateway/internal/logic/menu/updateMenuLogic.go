package menu

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新菜单
func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuLogic) UpdateMenu(req *types.UpdateMenuReq) (resp *types.BaseResp, err error) {
	// 调用RPC服务更新菜单
	rpcResp, err := l.svcCtx.AdminRpc.UpdateMenu(l.ctx, &adminRpc.UpdateMenuReq{
		MenuId:    req.Id,
		Name:      req.Name,
		Path:      req.Path,
		Icon:      req.Icon,
		Component: req.Component,
		Role:      req.Role,
		Label:     req.Label,
		Alias:     req.Alias,
		Type:      int32(req.Type),
		ParentId:  req.ParentId,
		Sort:      int32(req.Sort),
		Status:    1, // 默认启用
		RequestId: "update-menu",
	})
	if err != nil {
		l.Errorf("调用RPC更新菜单失败: %v", err)
		return &types.BaseResp{
			Code:    500,
			Message: "更新菜单失败",
		}, nil
	}

	return &types.BaseResp{
		Code:    int32(rpcResp.Base.Code),
		Message: rpcResp.Base.Message,
	}, nil
}
