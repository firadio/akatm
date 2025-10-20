package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCountriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCountriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCountriesLogic {
	return &ListCountriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ListCountries 获取国家列表
func (l *ListCountriesLogic) ListCountries(in *adminRpc.ListCountriesReq) (*adminRpc.ListCountriesResp, error) {
	// 获取国家列表
	countries, total, err := l.svcCtx.CountryRepository.List(
		in.PageReq.Page,
		in.PageReq.PageSize,
		in.Keyword,
		int8(in.Status),
	)
	if err != nil {
		l.Errorf("获取国家列表失败: %v", err)
		return &adminRpc.ListCountriesResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "获取国家列表失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 转换为响应格式
	var countryInfos []*adminRpc.CountryInfo
	for _, country := range countries {
		countryInfos = append(countryInfos, &adminRpc.CountryInfo{
			Id:          int64(country.ID),
			Code:        country.Code,
			Name:        country.Name,
			NameEn:      country.NameEn,
			Currency:    country.Currency,
			PhoneCode:   country.PhoneCode,
			Sort:        int32(country.Sort),
			Status:      int32(country.Status),
			Description: country.Description,
			CreatedAt:   country.CreatedAt.Unix(),
			UpdatedAt:   country.UpdatedAt.Unix(),
		})
	}

	return &adminRpc.ListCountriesResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "获取成功",
			RequestId: in.RequestId,
		},
		PageResp: &adminRpc.PageResp{
			Total:    total,
			Page:     in.PageReq.Page,
			PageSize: in.PageReq.PageSize,
		},
		Countries: countryInfos,
	}, nil
}
