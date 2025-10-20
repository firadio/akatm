package staff

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStaffLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加人员
func NewCreateStaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStaffLogic {
	return &CreateStaffLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStaffLogic) CreateStaff(req *types.CreateStaffReq) (resp *types.CreateStaffResp, err error) {
	// 调用RPC服务创建员工
	rpcReq := &admin.CreateStaffReq{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		UserId:    0, // 暂时设为0，后续可以关联用户
		RequestId: "req-" + req.Email,
	}

	rpcResp, err := l.svcCtx.AdminRpc.CreateStaff(l.ctx, rpcReq)
	if err != nil {
		l.Errorf("调用RPC创建员工失败: %v", err)
		return &types.CreateStaffResp{
			Code:    500,
			Message: "创建员工失败",
		}, nil
	}

	// 转换响应格式
	if rpcResp.Base.Code != 200 {
		return &types.CreateStaffResp{
			Code:    int32(rpcResp.Base.Code),
			Message: rpcResp.Base.Message,
		}, nil
	}

	return &types.CreateStaffResp{
		Code:    200,
		Message: "创建成功",
		Data: types.StaffDetail{
			Id:        rpcResp.Data.Id,
			Name:      rpcResp.Data.Name,
			Username:  req.Username, // 使用请求中的用户名
			Email:     rpcResp.Data.Email,
			Status:    rpcResp.Data.Status,
			CreatedAt: rpcResp.Data.CreatedAt,
		},
	}, nil
}
