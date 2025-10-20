package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFundDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFundDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFundDetailLogic {
	return &GetFundDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetFundDetail 获取资金明细详情
func (l *GetFundDetailLogic) GetFundDetail(in *famsRpc.GetFundDetailReq) (*famsRpc.GetFundDetailResp, error) {
	// 获取资金明细
	fundDetail, err := l.svcCtx.FundDetailRepository.GetByID(uint(in.FundDetailId))
	if err != nil {
		l.Errorf("获取资金明细失败: %v", err)
		return &famsRpc.GetFundDetailResp{
			Base: &famsRpc.BaseResp{
				Code:      500,
				Message:   "获取资金明细失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 转换为响应格式
	fundDetailInfo := &famsRpc.FundDetailInfo{
		Id:                int64(fundDetail.ID),
		TransactionNumber: fundDetail.TransactionNumber,
		UserId:            int64(fundDetail.UserId),
		UserEmail:         "", // TODO: 需要关联查询用户邮箱
		UserType:          fundDetail.UserType,
		ParentUserId:      int64(fundDetail.ParentUserId),
		ParentUserEmail:   "", // TODO: 需要关联查询父级用户邮箱
		TransactionType:   fundDetail.TransactionType,
		Amount:            fundDetail.Amount.String(),
		Currency:          fundDetail.Currency,
		Fee:               fundDetail.Fee.String(),
		ActualAmount:      fundDetail.ActualAmount.String(),
		Status:            fundDetail.Status,
		Description:       fundDetail.Description,
		CreatedAt:         fundDetail.CreatedAt.Unix(),
		UpdatedAt:         fundDetail.UpdatedAt.Unix(),
	}

	return &famsRpc.GetFundDetailResp{
		Base: &famsRpc.BaseResp{
			Code:      200,
			Message:   "获取成功",
			RequestId: in.RequestId,
		},
		Data: fundDetailInfo,
	}, nil
}
