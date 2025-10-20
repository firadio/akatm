package logic

import (
	"context"

	"akatm/rpc/mail/internal/svc"
	"akatm/rpc/mail/pb/emailRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailSendStatsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmailSendStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailSendStatsLogic {
	return &GetEmailSendStatsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ============ 邮件统计相关 ============
func (l *GetEmailSendStatsLogic) GetEmailSendStats(in *emailRpc.GetEmailSendStatsReq) (*emailRpc.GetEmailSendStatsResp, error) {
	// todo: add your logic here and delete this line

	return &emailRpc.GetEmailSendStatsResp{}, nil
}
