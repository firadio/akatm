package country

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCountryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCountryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCountryLogic {
	return &DeleteCountryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCountryLogic) DeleteCountry(req *types.IdReq) (resp *types.BaseResp, err error) {
	// 调用RPC服务
	rpcResp, err := l.svcCtx.AdminRpc.DeleteCountry(l.ctx, &adminRpc.DeleteCountryReq{
		CountryId: req.Id,
		RequestId: "delete-country",
	})

	if err != nil {
		l.Errorf("调用RPC服务失败: %v", err)
		return &types.BaseResp{
			Code:    500,
			Message: "删除国家失败",
		}, nil
	}

	return &types.BaseResp{
		Code:    rpcResp.Base.Code,
		Message: rpcResp.Base.Message,
	}, nil
}
