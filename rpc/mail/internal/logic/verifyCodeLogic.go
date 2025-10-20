package logic

import (
	"context"

	"akatm/rpc/mail/internal/svc"
	"akatm/rpc/mail/pb/emailRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCodeLogic {
	return &VerifyCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证验证码
func (l *VerifyCodeLogic) VerifyCode(in *emailRpc.VerifyCodeReq) (*emailRpc.VerifyCodeResp, error) {
	// todo: add your logic here and delete this line

	return &emailRpc.VerifyCodeResp{}, nil
}
