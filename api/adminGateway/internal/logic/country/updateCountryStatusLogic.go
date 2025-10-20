package country

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCountryStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCountryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCountryStatusLogic {
	return &UpdateCountryStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCountryStatusLogic) UpdateCountryStatus(req *types.UpdateCountryStatusReq) (resp *types.BaseResp, err error) {
	// 调用RPC服务
	rpcResp, err := l.svcCtx.AdminRpc.UpdateCountryStatus(l.ctx, &adminRpc.UpdateCountryStatusReq{
		CountryId: req.Id,
		Status:    req.Status,
		RequestId: "update-country-status",
	})

	if err != nil {
		l.Errorf("调用RPC服务失败: %v", err)
		return &types.BaseResp{
			Code:    500,
			Message: "更新国家状态失败",
		}, nil
	}

	return &types.BaseResp{
		Code:    rpcResp.Base.Code,
		Message: rpcResp.Base.Message,
	}, nil
}
