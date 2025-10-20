package logic

import (
	"context"

	"akatm/rpc/mail/internal/svc"
	"akatm/rpc/mail/pb/emailRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmailTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateEmailTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailTemplateLogic {
	return &UpdateEmailTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新邮件模板
func (l *UpdateEmailTemplateLogic) UpdateEmailTemplate(in *emailRpc.UpdateEmailTemplateReq) (*emailRpc.UpdateEmailTemplateResp, error) {
	// todo: add your logic here and delete this line

	return &emailRpc.UpdateEmailTemplateResp{}, nil
}
