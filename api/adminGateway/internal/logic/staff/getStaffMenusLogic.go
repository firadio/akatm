package staff

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStaffMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 人员菜单
func NewGetStaffMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStaffMenusLogic {
	return &GetStaffMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStaffMenusLogic) GetStaffMenus(req *types.IdReq) (resp *types.GetStaffMenusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
