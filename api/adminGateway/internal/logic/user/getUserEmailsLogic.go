package user

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserEmailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户邮箱列表
func NewGetUserEmailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserEmailsLogic {
	return &GetUserEmailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserEmailsLogic) GetUserEmails(req *types.IdReq) (resp *types.GetUserEmailsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
