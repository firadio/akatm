package menu

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 菜单详情
func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuLogic) GetMenu(req *types.IdReq) (resp *types.GetMenuResp, err error) {
	// 调用RPC服务获取菜单
	rpcResp, err := l.svcCtx.AdminRpc.GetMenu(l.ctx, &adminRpc.GetMenuReq{
		MenuId:    req.Id,
		RequestId: "get-menu",
	})
	if err != nil {
		l.Errorf("调用RPC获取菜单失败: %v", err)
		return &types.GetMenuResp{
			Code:    500,
			Message: "获取菜单失败",
		}, nil
	}

	// 转换响应格式
	menuInfo := types.MenuInfo{
		Id:        rpcResp.Data.Id,
		Name:      rpcResp.Data.Name,
		Path:      rpcResp.Data.Path,
		Icon:      rpcResp.Data.Icon,
		Component: rpcResp.Data.Component,
		Role:      rpcResp.Data.Role,
		Label:     rpcResp.Data.Label,
		Alias:     rpcResp.Data.Alias,
		Type:      int(rpcResp.Data.Type),
		ParentId:  rpcResp.Data.ParentId,
		Sort:      int(rpcResp.Data.Sort),
	}

	return &types.GetMenuResp{
		Code:    int32(rpcResp.Base.Code),
		Message: rpcResp.Base.Message,
		Data:    menuInfo,
	}, nil
}
