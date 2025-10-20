package logic

import (
	"context"

	"akatm/rpc/mail/internal/svc"
	"akatm/rpc/mail/pb/emailRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResendEmailLogic {
	return &ResendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 重发邮件
func (l *ResendEmailLogic) ResendEmail(in *emailRpc.ResendEmailReq) (*emailRpc.ResendEmailResp, error) {
	// todo: add your logic here and delete this line

	return &emailRpc.ResendEmailResp{}, nil
}
