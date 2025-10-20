package country

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCountryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCountryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCountryLogic {
	return &UpdateCountryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCountryLogic) UpdateCountry(req *types.UpdateCountryReq) (resp *types.BaseResp, err error) {
	// 调用RPC服务
	rpcResp, err := l.svcCtx.AdminRpc.UpdateCountry(l.ctx, &adminRpc.UpdateCountryReq{
		CountryId:   req.Id,
		Code:        req.Code,
		Name:        req.Name,
		NameEn:      req.NameEn,
		Currency:    req.Currency,
		PhoneCode:   req.PhoneCode,
		Sort:        int32(req.Sort),
		Description: req.Description,
		RequestId:   "update-country",
	})

	if err != nil {
		l.Errorf("调用RPC服务失败: %v", err)
		return &types.BaseResp{
			Code:    500,
			Message: "更新国家失败",
		}, nil
	}

	return &types.BaseResp{
		Code:    rpcResp.Base.Code,
		Message: rpcResp.Base.Message,
	}, nil
}
