package withdrawal

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateWithdrawalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 申请提现
func NewCreateWithdrawalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateWithdrawalLogic {
	return &CreateWithdrawalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateWithdrawalLogic) CreateWithdrawal(req *types.CreateWithdrawalReq) (resp *types.CreateWithdrawalResp, err error) {
	// todo: add your logic here and delete this line

	return
}
