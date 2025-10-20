package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetRole 获取角色详情
func (l *GetRoleLogic) GetRole(in *adminRpc.GetRoleReq) (*adminRpc.GetRoleResp, error) {
	// 根据ID获取角色
	role, err := l.svcCtx.RoleRepository.GetByID(uint(in.RoleId))
	if err != nil {
		l.Errorf("获取角色失败: %v", err)
		return &adminRpc.GetRoleResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "获取角色失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 转换为响应格式
	roleInfo := &adminRpc.RoleInfo{
		Id:          int64(role.ID),
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
		Status:      int32(role.Status),
		CreatedAt:   role.CreatedAt.Unix(),
	}

	return &adminRpc.GetRoleResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "获取成功",
			RequestId: in.RequestId,
		},
		Data: roleInfo,
	}, nil
}
