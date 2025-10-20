package auditAccountApplication

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountApplicationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 开户申请详情
func NewGetAccountApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountApplicationLogic {
	return &GetAccountApplicationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAccountApplicationLogic) GetAccountApplication(req *types.IdReq) (resp *types.GetAccountApplicationResp, err error) {
	// todo: add your logic here and delete this line

	return
}
