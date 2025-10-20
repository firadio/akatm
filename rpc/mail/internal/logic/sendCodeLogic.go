package logic

import (
	"context"

	"akatm/rpc/mail/internal/svc"
	"akatm/rpc/mail/pb/emailRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送验证码邮件
func (l *SendCodeLogic) SendCode(in *emailRpc.SendCodeReq) (*emailRpc.SendCodeResp, error) {
	// todo: add your logic here and delete this line

	return &emailRpc.SendCodeResp{}, nil
}
