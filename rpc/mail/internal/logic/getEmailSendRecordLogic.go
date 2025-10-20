package logic

import (
	"context"

	"akatm/rpc/mail/internal/svc"
	"akatm/rpc/mail/pb/emailRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailSendRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmailSendRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailSendRecordLogic {
	return &GetEmailSendRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 邮件发送记录相关 ============
func (l *GetEmailSendRecordLogic) GetEmailSendRecord(in *emailRpc.GetEmailSendRecordReq) (*emailRpc.GetEmailSendRecordResp, error) {
	// todo: add your logic here and delete this line

	return &emailRpc.GetEmailSendRecordResp{}, nil
}
