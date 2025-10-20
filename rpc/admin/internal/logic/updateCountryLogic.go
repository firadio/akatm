package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCountryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCountryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCountryLogic {
	return &UpdateCountryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCountry 更新国家
func (l *UpdateCountryLogic) UpdateCountry(in *adminRpc.UpdateCountryReq) (*adminRpc.UpdateCountryResp, error) {
	// 获取现有国家
	country, err := l.svcCtx.CountryRepository.GetByID(uint(in.CountryId))
	if err != nil {
		l.Errorf("获取国家失败: %v", err)
		return &adminRpc.UpdateCountryResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "获取国家失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 检查国家代码是否与其他国家冲突
	if country.Code != in.Code {
		existingCountry, err := l.svcCtx.CountryRepository.GetByCode(in.Code)
		if err == nil && existingCountry != nil && existingCountry.ID != country.ID {
			l.Errorf("国家代码已存在: %s", in.Code)
			return &adminRpc.UpdateCountryResp{
				Base: &adminRpc.BaseResp{
					Code:      400,
					Message:   "国家代码已存在",
					RequestId: in.RequestId,
				},
			}, nil
		}
	}

	// 更新国家信息
	country.Code = in.Code
	country.Name = in.Name
	country.NameEn = in.NameEn
	country.Currency = in.Currency
	country.PhoneCode = in.PhoneCode
	country.Sort = int(in.Sort)
	country.Description = in.Description

	if err := l.svcCtx.CountryRepository.Update(country); err != nil {
		l.Errorf("更新国家失败: %v", err)
		return &adminRpc.UpdateCountryResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "更新国家失败",
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

	return &adminRpc.UpdateCountryResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "更新成功",
			RequestId: in.RequestId,
		},
		Data: countryInfo,
	}, nil
}
