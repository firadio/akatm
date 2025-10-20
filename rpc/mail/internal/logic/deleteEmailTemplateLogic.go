package logic

import (
	"context"

	"akatm/rpc/mail/internal/svc"
	"akatm/rpc/mail/pb/emailRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmailTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteEmailTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmailTemplateLogic {
	return &DeleteEmailTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除邮件模板
func (l *DeleteEmailTemplateLogic) DeleteEmailTemplate(in *emailRpc.DeleteEmailTemplateReq) (*emailRpc.DeleteEmailTemplateResp, error) {
	// todo: add your logic here and delete this line

	return &emailRpc.DeleteEmailTemplateResp{}, nil
}
