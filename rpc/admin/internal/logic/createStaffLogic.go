package logic

import (
	"context"
	"crypto/md5"
	"fmt"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/orm/table"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStaffLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStaffLogic {
	return &CreateStaffLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateStaff 创建员工
func (l *CreateStaffLogic) CreateStaff(in *adminRpc.CreateStaffReq) (*adminRpc.CreateStaffResp, error) {
	// 检查邮箱是否已存在
	exists, err := l.svcCtx.StaffRepository.ExistsByEmail(in.Email)
	if err != nil {
		l.Errorf("检查邮箱失败: %v", err)
		return &adminRpc.CreateStaffResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "系统错误",
				RequestId: in.RequestId,
			},
		}, nil
	}
	if exists {
		return &adminRpc.CreateStaffResp{
			Base: &adminRpc.BaseResp{
				Code:      400,
				Message:   "邮箱已存在",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 密码加密
	hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(in.Password)))

	// 创建员工
	staff := &table.Staff{
		Name:     in.Name,
		Email:    in.Email,
		Password: hashedPassword,
		Status:   1, // 默认启用
	}

	if err := l.svcCtx.StaffRepository.Create(staff); err != nil {
		l.Errorf("创建员工失败: %v", err)
		return &adminRpc.CreateStaffResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "创建员工失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 返回创建的员工信息
	staffInfo := &adminRpc.StaffInfo{
		Id:        int64(staff.ID),
		Name:      staff.Name,
		Email:     staff.Email,
		Status:    int32(staff.Status),
		CreatedAt: staff.CreatedAt.Unix(),
		UpdatedAt: staff.UpdatedAt.Unix(),
	}

	return &adminRpc.CreateStaffResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "创建成功",
			RequestId: in.RequestId,
		},
		Data: staffInfo,
	}, nil
}
