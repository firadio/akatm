package logic

import (
	"context"

	"akatm/rpc/mail/internal/svc"
	"akatm/rpc/mail/pb/emailRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListEmailSendRecordsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListEmailSendRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListEmailSendRecordsLogic {
	return &ListEmailSendRecordsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取邮件发送记录列表
func (l *ListEmailSendRecordsLogic) ListEmailSendRecords(in *emailRpc.ListEmailSendRecordsReq) (*emailRpc.ListEmailSendRecordsResp, error) {
	// todo: add your logic here and delete this line

	return &emailRpc.ListEmailSendRecordsResp{}, nil
}
