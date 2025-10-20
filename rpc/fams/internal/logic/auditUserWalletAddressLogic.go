package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuditUserWalletAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuditUserWalletAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuditUserWalletAddressLogic {
	return &AuditUserWalletAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 审核钱包地址
func (l *AuditUserWalletAddressLogic) AuditUserWalletAddress(in *famsRpc.AuditUserWalletAddressReq) (*famsRpc.AuditUserWalletAddressResp, error) {
	// todo: add your logic here and delete this line

	return &famsRpc.AuditUserWalletAddressResp{}, nil
}
