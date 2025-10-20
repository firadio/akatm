package country

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCountryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCountryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCountryLogic {
	return &CreateCountryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCountryLogic) CreateCountry(req *types.CreateCountryReq) (resp *types.CreateCountryResp, err error) {
	// 调用RPC服务
	rpcResp, err := l.svcCtx.AdminRpc.CreateCountry(l.ctx, &adminRpc.CreateCountryReq{
		Code:        req.Code,
		Name:        req.Name,
		NameEn:      req.NameEn,
		Currency:    req.Currency,
		PhoneCode:   req.PhoneCode,
		Sort:        int32(req.Sort),
		Description: req.Description,
		RequestId:   "create-country",
	})

	if err != nil {
		l.Errorf("调用RPC服务失败: %v", err)
		return &types.CreateCountryResp{
			Code:    500,
			Message: "创建国家失败",
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

	return &types.CreateCountryResp{
		Code:    rpcResp.Base.Code,
		Message: rpcResp.Base.Message,
		Data:    countryInfo,
	}, nil
}
