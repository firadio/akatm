package logic

import (
	"context"

	"akatm/rpc/fams/internal/svc"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFundSummaryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFundSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFundSummaryLogic {
	return &GetFundSummaryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetFundSummary 获取资金统计总览
func (l *GetFundSummaryLogic) GetFundSummary(in *famsRpc.GetFundSummaryReq) (*famsRpc.GetFundSummaryResp, error) {
	// 获取资金统计总览
	summaryData, err := l.svcCtx.FundDetailRepository.GetFundSummary(
		in.StartTime,
		in.EndTime,
		in.UserType,
		in.Currency,
	)
	if err != nil {
		l.Errorf("获取资金统计总览失败: %v", err)
		return &famsRpc.GetFundSummaryResp{
			Base: &famsRpc.BaseResp{
				Code:      500,
				Message:   "获取资金统计总览失败",
				RequestId: in.RequestId,
			},
		}, nil
	}

	// 转换为响应格式
	var currencyStats []*famsRpc.CurrencyStat
	for _, stat := range summaryData.CurrencyStats {
		currencyStats = append(currencyStats, &famsRpc.CurrencyStat{
			Currency:         stat.Currency,
			DepositAmount:    stat.DepositAmount,
			WithdrawalAmount: stat.WithdrawalAmount,
			FeeAmount:        stat.FeeAmount,
			TransactionCount: stat.TransactionCount,
		})
	}

	fundSummaryData := &famsRpc.FundSummaryData{
		TotalDeposit:     summaryData.TotalDeposit,
		TotalWithdrawal:  summaryData.TotalWithdrawal,
		TotalFee:         summaryData.TotalFee,
		ActiveUsers:      summaryData.ActiveUsers,
		TransactionCount: summaryData.TransactionCount,
		CurrencyStats:    currencyStats,
	}

	return &famsRpc.GetFundSummaryResp{
		Base: &famsRpc.BaseResp{
			Code:      200,
			Message:   "获取成功",
			RequestId: in.RequestId,
		},
		Data: fundSummaryData,
	}, nil
}
