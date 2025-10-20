package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFundDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFundDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFundDetailsLogic {
	return &GetUserFundDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserFundDetails 获取用户资金明细
func (l *GetUserFundDetailsLogic) GetUserFundDetails(in *famsRpc.GetUserFundDetailsReq) (*famsRpc.ListFundDetailsResp, error) {
	// 获取用户资金明细
	fundDetails, total, err := l.svcCtx.FundDetailRepository.GetUserFundDetails(
		uint(in.UserId),
		in.PageReq.Page,
		in.PageReq.PageSize,
		in.TransactionType,
		in.Status,
		in.StartTime,
		in.EndTime,
	)
	if err != nil {
		l.Errorf("获取用户资金明细失败: %v", err)
		return &famsRpc.ListFundDetailsResp{
			Base: &famsRpc.BaseResp{
				Code:      500,
				Message:   "获取用户资金明细失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 转换为响应格式
	var fundDetailInfos []*famsRpc.FundDetailInfo
	for _, fundDetail := range fundDetails {
		fundDetailInfos = append(fundDetailInfos, &famsRpc.FundDetailInfo{
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
		})
	}

	return &famsRpc.ListFundDetailsResp{
		Base: &famsRpc.BaseResp{
			Code:      200,
			Message:   "获取成功",
			RequestId: in.RequestId,
		},
		PageResp: &famsRpc.PageResp{
			Total:    total,
			Page:     in.PageReq.Page,
			PageSize: in.PageReq.PageSize,
		},
		FundDetails: fundDetailInfos,
	}, nil
}
