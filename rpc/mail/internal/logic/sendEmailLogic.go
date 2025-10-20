package logic

import (
	"context"

	"akatm/rpc/mail/internal/svc"
	"akatm/rpc/mail/pb/emailRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 邮件发送相关 ============
func (l *SendEmailLogic) SendEmail(in *emailRpc.SendEmailReq) (*emailRpc.SendEmailResp, error) {
	// todo: add your logic here and delete this line

	return &emailRpc.SendEmailResp{}, nil
}
