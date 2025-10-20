package account

import (
	"context"

	"akatm/api/managerGateway/internal/svc"
	"akatm/api/managerGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAccountApplicationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建开户申请
func NewCreateAccountApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAccountApplicationLogic {
	return &CreateAccountApplicationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAccountApplicationLogic) CreateAccountApplication(req *types.CreateAccountApplicationReq) (resp *types.CreateAccountApplicationResp, err error) {
	// todo: add your logic here and delete this line

	return
}
