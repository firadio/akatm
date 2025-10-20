package logic

import (
	"context"

	"akatm/rpc/mail/internal/svc"
	"akatm/rpc/mail/pb/emailRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListEmailTemplatesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListEmailTemplatesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListEmailTemplatesLogic {
	return &ListEmailTemplatesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取邮件模板列表
func (l *ListEmailTemplatesLogic) ListEmailTemplates(in *emailRpc.ListEmailTemplatesReq) (*emailRpc.ListEmailTemplatesResp, error) {
	// todo: add your logic here and delete this line

	return &emailRpc.ListEmailTemplatesResp{}, nil
}
