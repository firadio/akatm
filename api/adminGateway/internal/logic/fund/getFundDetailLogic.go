package fund

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFundDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFundDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFundDetailLogic {
	return &GetFundDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFundDetailLogic) GetFundDetail(req *types.IdReq) (resp *types.GetFundDetailResp, err error) {
	// 调用FAMS RPC服务
	rpcResp, err := l.svcCtx.FamsRpc.GetFundDetail(l.ctx, &famsRpc.GetFundDetailReq{
		FundDetailId: req.Id,
		RequestId:    "get-fund-detail",
	})

	if err != nil {
		l.Errorf("调用FAMS RPC服务失败: %v", err)
		return &types.GetFundDetailResp{
			Code:    500,
			Message: "获取资金明细失败",
		}, nil
	}

	// 转换响应
	fundDetailInfo := types.FundDetailInfo{
		Id:                rpcResp.Data.Id,
		TransactionNumber: rpcResp.Data.TransactionNumber,
		UserId:            rpcResp.Data.UserId,
		UserEmail:         rpcResp.Data.UserEmail,
		UserType:          rpcResp.Data.UserType,
		ParentUserId:      rpcResp.Data.ParentUserId,
		ParentUserEmail:   rpcResp.Data.ParentUserEmail,
		TransactionType:   rpcResp.Data.TransactionType,
		Amount:            rpcResp.Data.Amount,
		Currency:          rpcResp.Data.Currency,
		Fee:               rpcResp.Data.Fee,
		ActualAmount:      rpcResp.Data.ActualAmount,
		Status:            rpcResp.Data.Status,
		Description:       rpcResp.Data.Description,
		CreatedAt:         rpcResp.Data.CreatedAt,
		UpdatedAt:         rpcResp.Data.UpdatedAt,
	}

	return &types.GetFundDetailResp{
		Code:    rpcResp.Base.Code,
		Message: rpcResp.Base.Message,
		Data:    fundDetailInfo,
	}, nil
}
