package logic

import (
	"context"

	"akatm/rpc/mail/internal/svc"
	"akatm/rpc/mail/pb/emailRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEmailTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateEmailTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmailTemplateLogic {
	return &CreateEmailTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 邮件模板管理相关 ============
func (l *CreateEmailTemplateLogic) CreateEmailTemplate(in *emailRpc.CreateEmailTemplateReq) (*emailRpc.CreateEmailTemplateResp, error) {
	// todo: add your logic here and delete this line

	return &emailRpc.CreateEmailTemplateResp{}, nil
}
