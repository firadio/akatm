package staff

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStaffLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新人员
func NewUpdateStaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStaffLogic {
	return &UpdateStaffLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStaffLogic) UpdateStaff(req *types.UpdateStaffReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
