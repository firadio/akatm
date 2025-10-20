package staff

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStaffStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新状态
func NewUpdateStaffStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStaffStatusLogic {
	return &UpdateStaffStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStaffStatusLogic) UpdateStaffStatus(req *types.UpdateStaffStatusReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
