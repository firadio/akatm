package staff

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetStaffPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 重设密码
func NewResetStaffPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetStaffPasswordLogic {
	return &ResetStaffPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetStaffPasswordLogic) ResetStaffPassword(req *types.ResetStaffPasswordReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
