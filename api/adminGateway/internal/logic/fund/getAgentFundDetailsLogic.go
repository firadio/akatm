package fund

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAgentFundDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAgentFundDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAgentFundDetailsLogic {
	return &GetAgentFundDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAgentFundDetailsLogic) GetAgentFundDetails(req *types.GetAgentFundDetailsReq) (resp *types.ListFundDetailsResp, err error) {
	// 调用FAMS RPC服务
	rpcResp, err := l.svcCtx.FamsRpc.GetAgentFundDetails(l.ctx, &famsRpc.GetAgentFundDetailsReq{
		AgentId: req.AgentId,
		PageReq: &famsRpc.PageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		IncludeSubAgents: req.IncludeSubAgents,
		TransactionType:  req.TransactionType,
		Status:           req.Status,
		StartTime:        req.StartTime,
		EndTime:          req.EndTime,
		RequestId:        "get-agent-fund-details",
	})

	if err != nil {
		l.Errorf("调用FAMS RPC服务失败: %v", err)
		return &types.ListFundDetailsResp{
			Code:    500,
			Message: "获取代理资金明细失败",
		}, nil
	}

	// 转换响应
	var fundDetails []types.FundDetailInfo
	for _, fundDetail := range rpcResp.FundDetails {
		fundDetails = append(fundDetails, types.FundDetailInfo{
			Id:                fundDetail.Id,
			TransactionNumber: fundDetail.TransactionNumber,
			UserId:            fundDetail.UserId,
			UserEmail:         fundDetail.UserEmail,
			UserType:          fundDetail.UserType,
			ParentUserId:      fundDetail.ParentUserId,
			ParentUserEmail:   fundDetail.ParentUserEmail,
			TransactionType:   fundDetail.TransactionType,
			Amount:            fundDetail.Amount,
			Currency:          fundDetail.Currency,
			Fee:               fundDetail.Fee,
			ActualAmount:      fundDetail.ActualAmount,
			Status:            fundDetail.Status,
			Description:       fundDetail.Description,
			CreatedAt:         fundDetail.CreatedAt,
			UpdatedAt:         fundDetail.UpdatedAt,
		})
	}

	return &types.ListFundDetailsResp{
		Code:    rpcResp.Base.Code,
		Message: rpcResp.Base.Message,
		Data: types.FundDetailsListData{
			Total: rpcResp.PageResp.Total,
			List:  fundDetails,
		},
	}, nil
}
