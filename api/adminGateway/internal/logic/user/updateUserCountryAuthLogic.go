package user

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserCountryAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户国家授权
func NewUpdateUserCountryAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserCountryAuthLogic {
	return &UpdateUserCountryAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserCountryAuthLogic) UpdateUserCountryAuth(req *types.UpdateUserCountryAuthReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
