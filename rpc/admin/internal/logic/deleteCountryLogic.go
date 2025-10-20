package logic

import (
	"context"

	"akatm/rpc/admin/internal/svc"
	"akatm/rpc/admin/pb/adminRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCountryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCountryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCountryLogic {
	return &DeleteCountryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteCountry 删除国家
func (l *DeleteCountryLogic) DeleteCountry(in *adminRpc.DeleteCountryReq) (*adminRpc.DeleteCountryResp, error) {
	// 检查国家是否存在
	country, err := l.svcCtx.CountryRepository.GetByID(uint(in.CountryId))
	if err != nil {
		l.Errorf("获取国家失败: %v", err)
		return &adminRpc.DeleteCountryResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "国家不存在",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// TODO: 检查是否有相关的开户申请使用此国家
	// 这里可以添加业务逻辑检查，如果有关联数据则不允许删除

	// 删除国家
	if err := l.svcCtx.CountryRepository.Delete(uint(in.CountryId)); err != nil {
		l.Errorf("删除国家失败: %v", err)
		return &adminRpc.DeleteCountryResp{
			Base: &adminRpc.BaseResp{
				Code:      500,
				Message:   "删除国家失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	l.Infof("成功删除国家: %s (ID: %d)", country.Name, in.CountryId)

	return &adminRpc.DeleteCountryResp{
		Base: &adminRpc.BaseResp{
			Code:      200,
			Message:   "删除成功",
			RequestId: in.RequestId,
		},
	}, nil
}
