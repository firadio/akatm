package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBankAccountApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBankAccountApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBankAccountApplicationLogic {
	return &GetBankAccountApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取银行账户申请
func (l *GetBankAccountApplicationLogic) GetBankAccountApplication(in *famsRpc.GetBankAccountApplicationReq) (*famsRpc.GetBankAccountApplicationResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.GetBankAccountApplicationResp{}, nil
}
