package logic

import (
	"context"

	"akatm/rpc/mail/internal/svc"
	"akatm/rpc/mail/pb/emailRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmailTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailTemplateLogic {
	return &GetEmailTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取邮件模板
func (l *GetEmailTemplateLogic) GetEmailTemplate(in *emailRpc.GetEmailTemplateReq) (*emailRpc.GetEmailTemplateResp, error) {
	// todo: add your logic here and delete this line

	return &emailRpc.GetEmailTemplateResp{}, nil
}
