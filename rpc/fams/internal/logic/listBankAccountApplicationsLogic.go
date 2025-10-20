package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListBankAccountApplicationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListBankAccountApplicationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBankAccountApplicationsLogic {
	return &ListBankAccountApplicationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取银行账户申请列表
func (l *ListBankAccountApplicationsLogic) ListBankAccountApplications(in *famsRpc.ListBankAccountApplicationsReq) (*famsRpc.ListBankAccountApplicationsResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.ListBankAccountApplicationsResp{}, nil
}
