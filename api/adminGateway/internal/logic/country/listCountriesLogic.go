package country

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCountriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCountriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCountriesLogic {
	return &ListCountriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCountriesLogic) ListCountries(req *types.ListCountryReq) (resp *types.ListCountryResp, err error) {
	// 调用RPC服务
	rpcResp, err := l.svcCtx.AdminRpc.ListCountries(l.ctx, &adminRpc.ListCountriesReq{
		PageReq: &adminRpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Keyword:   req.Keyword,
		Status:    req.Status,
		RequestId: "list-countries",
	})

	if err != nil {
		l.Errorf("调用RPC服务失败: %v", err)
		return &types.ListCountryResp{
			Code:    500,
			Message: "获取国家列表失败",
		}, nil
	}

	// 转换响应
	var countries []types.CountryInfo
	for _, country := range rpcResp.Countries {
		countries = append(countries, types.CountryInfo{
			Id:          country.Id,
			Code:        country.Code,
			Name:        country.Name,
			NameEn:      country.NameEn,
			Currency:    country.Currency,
			PhoneCode:   country.PhoneCode,
			Sort:        int(country.Sort),
			Status:      int(country.Status),
			Description: country.Description,
			CreatedAt:   country.CreatedAt,
			UpdatedAt:   country.UpdatedAt,
		})
	}

	return &types.ListCountryResp{
		Code:    rpcResp.Base.Code,
		Message: rpcResp.Base.Message,
		Data: types.CountryListData{
			Total: rpcResp.PageResp.Total,
			List:  countries,
		},
	}, nil
}
