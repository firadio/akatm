package fund

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFundSummaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFundSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFundSummaryLogic {
	return &GetFundSummaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFundSummaryLogic) GetFundSummary(req *types.GetFundSummaryReq) (resp *types.GetFundSummaryResp, err error) {
	// 调用FAMS RPC服务
	rpcResp, err := l.svcCtx.FamsRpc.GetFundSummary(l.ctx, &famsRpc.GetFundSummaryReq{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		UserType:  req.UserType,
		Currency:  req.Currency,
		RequestId: "get-fund-summary",
	})

	if err != nil {
		l.Errorf("调用FAMS RPC服务失败: %v", err)
		return &types.GetFundSummaryResp{
			Code:    500,
			Message: "获取资金统计总览失败",
		}, nil
	}

	// 转换响应
	var currencyStats []types.CurrencyStat
	for _, stat := range rpcResp.Data.CurrencyStats {
		currencyStats = append(currencyStats, types.CurrencyStat{
			Currency:         stat.Currency,
			DepositAmount:    stat.DepositAmount,
			WithdrawalAmount: stat.WithdrawalAmount,
			FeeAmount:        stat.FeeAmount,
			TransactionCount: stat.TransactionCount,
		})
	}

	fundSummaryData := types.FundSummaryData{
		TotalDeposit:     rpcResp.Data.TotalDeposit,
		TotalWithdrawal:  rpcResp.Data.TotalWithdrawal,
		TotalFee:         rpcResp.Data.TotalFee,
		ActiveUsers:      rpcResp.Data.ActiveUsers,
		TransactionCount: rpcResp.Data.TransactionCount,
		CurrencyStats:    currencyStats,
	}

	return &types.GetFundSummaryResp{
		Code:    rpcResp.Base.Code,
		Message: rpcResp.Base.Message,
		Data:    fundSummaryData,
	}, nil
}
