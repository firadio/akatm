package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCountryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCountryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCountryLogic {
	return &GetCountryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetCountry 获取国家
func (l *GetCountryLogic) GetCountry(in *adminRpc.GetCountryReq) (*adminRpc.GetCountryResp, error) {
	// 获取国家
	country, err := l.svcCtx.CountryRepository.GetByID(uint(in.CountryId))
	if err != nil {
		l.Errorf("获取国家失败: %v", err)
		return &adminRpc.GetCountryResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "获取国家失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 转换为响应格式
	countryInfo := &adminRpc.CountryInfo{
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
	}

	return &adminRpc.GetCountryResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "获取成功",
			RequestId: in.RequestId,
		},
		Data: countryInfo,
	}, nil
}
