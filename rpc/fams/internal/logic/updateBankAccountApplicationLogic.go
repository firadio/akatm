package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateBankAccountApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBankAccountApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBankAccountApplicationLogic {
	return &UpdateBankAccountApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新银行账户申请
func (l *UpdateBankAccountApplicationLogic) UpdateBankAccountApplication(in *famsRpc.UpdateBankAccountApplicationReq) (*famsRpc.UpdateBankAccountApplicationResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.UpdateBankAccountApplicationResp{}, nil
}
