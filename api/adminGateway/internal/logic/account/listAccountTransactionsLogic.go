package account

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAccountTransactionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 账户流水
func NewListAccountTransactionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAccountTransactionsLogic {
	return &ListAccountTransactionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAccountTransactionsLogic) ListAccountTransactions(req *types.ListAccountTransactionsReq) (resp *types.ListAccountTransactionsResp, err error) {
	// 调用FAMS RPC服务获取账户流水
	rpcResp, err := l.svcCtx.FamsRpc.ListAccountTransactions(l.ctx, &famsRpc.ListAccountTransactionsReq{
		AccountId: req.Id,
		PageReq: &famsRpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		TransactionType: req.TransactionType,
		Status:          req.Status,
		StartTime:       req.StartTime,
		EndTime:         req.EndTime,
		MinAmount:       req.MinAmount,
		MaxAmount:       req.MaxAmount,
		Keyword:         req.Keyword,
		RequestId:       "list-account-transactions",
	})

	if err != nil {
		l.Errorf("调用FAMS RPC服务失败: %v", err)
		return &types.ListAccountTransactionsResp{
			Code:    500,
			Message: "获取账户流水失败",
		}, nil
	}

	// 转换响应
	var transactions []types.TransactionBrief
	for _, tx := range rpcResp.Data.List {
		transactions = append(transactions, types.TransactionBrief{
			Id:                   tx.Id,
			TransactionNumber:    tx.TransactionNumber,
			Type:                 tx.Type,
			Amount:               tx.Amount,
			Symbol:               tx.Symbol,
			Status:               tx.Status,
			Fee:                  tx.Fee,
			ActualAmount:         tx.ActualAmount,
			Description:          tx.Description,
			RelatedAccountId:     tx.RelatedAccountId,
			RelatedAccountNumber: tx.RelatedAccountNumber,
			CreatedAt:            tx.CreatedAt,
			UpdatedAt:            tx.UpdatedAt,
		})
	}

	return &types.ListAccountTransactionsResp{
		Code:    200,
		Message: "获取账户流水成功",
		Data: types.TransactionListData{
			Total: rpcResp.Data.Total,
			List:  transactions,
		},
	}, nil
}
