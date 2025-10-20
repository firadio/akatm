package menu

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除菜单
func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMenuLogic) DeleteMenu(req *types.IdReq) (resp *types.BaseResp, err error) {
	// 调用RPC服务删除菜单
	rpcResp, err := l.svcCtx.AdminRpc.DeleteMenu(l.ctx, &adminRpc.DeleteMenuReq{
		MenuId:    req.Id,
		RequestId: "delete-menu",
	})
	if err != nil {
		l.Errorf("调用RPC删除菜单失败: %v", err)
		return &types.BaseResp{
			Code:    500,
			Message: "删除菜单失败",
		}, nil
	}

	return &types.BaseResp{
		Code:    int32(rpcResp.Base.Code),
		Message: rpcResp.Base.Message,
	}, nil
}
