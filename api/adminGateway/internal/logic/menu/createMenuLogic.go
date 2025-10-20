package menu

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建菜单
func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMenuLogic) CreateMenu(req *types.CreateMenuReq) (resp *types.CreateMenuResp, err error) {
	// 调用RPC服务创建菜单
	rpcResp, err := l.svcCtx.AdminRpc.CreateMenu(l.ctx, &adminRpc.CreateMenuReq{
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
		RequestId: "create-menu",
	})
	if err != nil {
		l.Errorf("调用RPC创建菜单失败: %v", err)
		return &types.CreateMenuResp{
			Code:    500,
			Message: "创建菜单失败",
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

	return &types.CreateMenuResp{
		Code:    int32(rpcResp.Base.Code),
		Message: rpcResp.Base.Message,
		Data:    menuInfo,
	}, nil
}
