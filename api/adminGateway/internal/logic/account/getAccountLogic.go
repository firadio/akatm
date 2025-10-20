package account

import (
	"context"

	"akatm/api/adminGateway/internal/svc"
	"akatm/api/adminGateway/internal/types"
	"akatm/rpc/fams/pb/famsRpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 账户详情
func NewGetAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountLogic {
	return &GetAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAccountLogic) GetAccount(req *types.IdReq) (resp *types.GetAccountResp, err error) {
	// 调用FAMS RPC服务获取账户详情
	rpcResp, err := l.svcCtx.FamsRpc.GetBankAccount(l.ctx, &famsRpc.GetBankAccountReq{
		AccountId: req.Id,
		RequestId: "get-account-detail",
	})

	if err != nil {
		l.Errorf("调用FAMS RPC服务失败: %v", err)
		return &types.GetAccountResp{
			Code:    500,
			Message: "获取账户详情失败",
		}, nil
	}

	// 暂时返回空的层级信息，后续可以通过其他方式获取
	hierarchyInfo := types.HierarchyInfo{
		SuperAgentId:   0,
		SuperAgentName: "",
		AgentId:        0,
		AgentName:      "",
		ManagerId:      rpcResp.Data.ManagerId,
		ManagerName:    rpcResp.Data.ManagerName,
		CustomerId:     rpcResp.Data.CustomerId,
		CustomerName:   rpcResp.Data.CustomerName,
	}

	return &types.GetAccountResp{
		Code:    200,
		Message: "获取账户详情成功",
		Data: types.AccountDetail{
			Id:            rpcResp.Data.Id,
			AccountNumber: rpcResp.Data.AccountNumber,
			Currency:      rpcResp.Data.Currency,
			Status:        rpcResp.Data.Status,
			CreatedAt:     rpcResp.Data.CreatedAt,
			CustomerId:    rpcResp.Data.CustomerId,
			CustomerName:  rpcResp.Data.CustomerName,
			CustomerEmail: rpcResp.Data.CustomerEmail,
			CustomerPhone: rpcResp.Data.CustomerPhone,
			ManagerId:     rpcResp.Data.ManagerId,
			ManagerName:   rpcResp.Data.ManagerName,
			ManagerEmail:  rpcResp.Data.ManagerEmail,
			ManagerLevel:  rpcResp.Data.ManagerLevel,
			HierarchyInfo: hierarchyInfo,
		},
	}, nil
}
