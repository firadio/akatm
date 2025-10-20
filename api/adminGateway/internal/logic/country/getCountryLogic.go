package country

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCountryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCountryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCountryLogic {
	return &GetCountryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCountryLogic) GetCountry(req *types.IdReq) (resp *types.GetCountryResp, err error) {
	// 调用RPC服务
	rpcResp, err := l.svcCtx.AdminRpc.GetCountry(l.ctx, &adminRpc.GetCountryReq{
		CountryId: req.Id,
		RequestId: "get-country",
	})

	if err != nil {
		l.Errorf("调用RPC服务失败: %v", err)
		return &types.GetCountryResp{
			Code:    500,
			Message: "获取国家失败",
		}, nil
	}

	// 转换响应
	countryInfo := types.CountryInfo{
		Id:          rpcResp.Data.Id,
		Code:        rpcResp.Data.Code,
		Name:        rpcResp.Data.Name,
		NameEn:      rpcResp.Data.NameEn,
		Currency:    rpcResp.Data.Currency,
		PhoneCode:   rpcResp.Data.PhoneCode,
		Sort:        int(rpcResp.Data.Sort),
		Status:      int(rpcResp.Data.Status),
		Description: rpcResp.Data.Description,
		CreatedAt:   rpcResp.Data.CreatedAt,
		UpdatedAt:   rpcResp.Data.UpdatedAt,
	}

	return &types.GetCountryResp{
		Code:    rpcResp.Base.Code,
		Message: rpcResp.Base.Message,
		Data:    countryInfo,
	}, nil
}
