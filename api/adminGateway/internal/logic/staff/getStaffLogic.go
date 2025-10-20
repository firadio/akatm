package staff

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStaffLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 人员详情
func NewGetStaffLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStaffLogic {
	return &GetStaffLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStaffLogic) GetStaff(req *types.IdReq) (resp *types.GetStaffResp, err error) {
	// todo: add your logic here and delete this line

	return
}
