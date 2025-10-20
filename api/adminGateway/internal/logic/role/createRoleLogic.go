package role

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建角色
func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.CreateRoleReq) (resp *types.CreateRoleResp, err error) {
	// 调用RPC服务创建角色
	rpcReq := &admin.CreateRoleReq{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Desc,
		RequestId:   "req-" + req.Name, // 生成请求ID
	}

	rpcResp, err := l.svcCtx.AdminRpc.CreateRole(l.ctx, rpcReq)
	if err != nil {
		l.Errorf("调用RPC创建角色失败: %v", err)
		return &types.CreateRoleResp{
			Code:    500,
			Message: "创建角色失败",
		}, nil
	}

	// 转换响应格式
	if rpcResp.Base.Code != 200 {
		return &types.CreateRoleResp{
			Code:    int32(rpcResp.Base.Code),
			Message: rpcResp.Base.Message,
		}, nil
	}

	return &types.CreateRoleResp{
		Code:    200,
		Message: "创建成功",
		Data: types.RoleDetail{
			Id:        rpcResp.Data.Id,
			Name:      rpcResp.Data.Name,
			Code:      rpcResp.Data.Code,
			Desc:      rpcResp.Data.Description,
			Status:    rpcResp.Data.Status,
			CreatedAt: rpcResp.Data.CreatedAt,
		},
	}, nil
}
