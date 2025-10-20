package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCountryStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCountryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCountryStatusLogic {
	return &UpdateCountryStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCountryStatus 更新国家状态
func (l *UpdateCountryStatusLogic) UpdateCountryStatus(in *adminRpc.UpdateCountryStatusReq) (*adminRpc.UpdateCountryStatusResp, error) {
	// 检查国家是否存在
	country, err := l.svcCtx.CountryRepository.GetByID(uint(in.CountryId))
	if err != nil {
		l.Errorf("获取国家失败: %v", err)
		return &adminRpc.UpdateCountryStatusResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "国家不存在",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 更新国家状态
	if err := l.svcCtx.CountryRepository.UpdateStatus(uint(in.CountryId), int8(in.Status)); err != nil {
		l.Errorf("更新国家状态失败: %v", err)
		return &adminRpc.UpdateCountryStatusResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "更新国家状态失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	statusText := "禁用"
	if in.Status == 1 {
		statusText = "启用"
	}

	l.Infof("成功更新国家状态: %s (ID: %d) -> %s", country.Name, in.CountryId, statusText)

	return &adminRpc.UpdateCountryStatusResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "更新成功",
			RequestId: in.RequestId,
		},
	}, nil
}
