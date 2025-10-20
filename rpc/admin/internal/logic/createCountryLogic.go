package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/orm/table"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCountryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCountryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCountryLogic {
	return &CreateCountryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateCountry 创建国家
func (l *CreateCountryLogic) CreateCountry(in *adminRpc.CreateCountryReq) (*adminRpc.CreateCountryResp, error) {
	// 检查国家代码是否已存在
	existingCountry, err := l.svcCtx.CountryRepository.GetByCode(in.Code)
	if err == nil && existingCountry != nil {
		l.Errorf("国家代码已存在: %s", in.Code)
		return &adminRpc.CreateCountryResp{
			Base: &adminRpc.BaseResp{
				Code:      400,
				Message:   "国家代码已存在",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 创建国家
	country := &table.Country{
		Code:        in.Code,
		Name:        in.Name,
		NameEn:      in.NameEn,
		Currency:    in.Currency,
		PhoneCode:   in.PhoneCode,
		Sort:        int(in.Sort),
		Status:      1, // 默认启用
		Description: in.Description,
	}

	if err := l.svcCtx.CountryRepository.Create(country); err != nil {
		l.Errorf("创建国家失败: %v", err)
		return &adminRpc.CreateCountryResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "创建国家失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 返回创建的国家信息
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

	return &adminRpc.CreateCountryResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "创建成功",
			RequestId: in.RequestId,
		},
		Data: countryInfo,
	}, nil
}
