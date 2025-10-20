package logic

import (
	"context"

	"akatm/rpc/mail/internal/svc"
	"akatm/rpc/mail/pb/emailRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeSendStatsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCodeSendStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodeSendStatsLogic {
	return &GetCodeSendStatsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取验证码发送统计
func (l *GetCodeSendStatsLogic) GetCodeSendStats(in *emailRpc.GetCodeSendStatsReq) (*emailRpc.GetCodeSendStatsResp, error) {
	// todo: add your logic here and delete this line

	return &emailRpc.GetCodeSendStatsResp{}, nil
}
