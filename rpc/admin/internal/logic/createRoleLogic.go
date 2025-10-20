package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/orm/table"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateRole 创建角色
func (l *CreateRoleLogic) CreateRole(in *adminRpc.CreateRoleReq) (*adminRpc.CreateRoleResp, error) {
	// 检查角色代码是否已存在
	exists, err := l.svcCtx.RoleRepository.ExistsByCode(in.Code)
	if err != nil {
		l.Errorf("检查角色代码失败: %v", err)
		return &adminRpc.CreateRoleResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "系统错误",
				RequestId: in.RequestId,
			},
		}, nil
	}
	if exists {
		return &adminRpc.CreateRoleResp{
			Base: &adminRpc.BaseResp{
				Code:      400,
				Message:   "角色代码已存在",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 创建角色
	role := &table.Role{
		Name:        in.Name,
		Code:        in.Code,
		Description: in.Description,
		Status:      1, // 默认启用
	}

	if err := l.svcCtx.RoleRepository.Create(role); err != nil {
		l.Errorf("创建角色失败: %v", err)
		return &adminRpc.CreateRoleResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "创建角色失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 返回创建的角色信息
	roleInfo := &adminRpc.RoleInfo{
		Id:          int64(role.ID),
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
		Status:      int32(role.Status),
		CreatedAt:   role.CreatedAt.Unix(),
	}

	return &adminRpc.CreateRoleResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "创建成功",
			RequestId: in.RequestId,
		},
		Data: roleInfo,
	}, nil
}
