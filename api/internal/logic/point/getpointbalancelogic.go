// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package point

import (
	"context"

	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"
	"jhb-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPointBalanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取积分余额
func NewGetPointBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPointBalanceLogic {
	return &GetPointBalanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPointBalanceLogic) GetPointBalance(req *types.UserIdReq) (resp *types.PointBalanceResp, err error) {
	userPoint, err := l.svcCtx.UserPointModel.FindOneByUserId(l.ctx, req.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			// 如果用户没有积分记录，返回0
			resp = &types.PointBalanceResp{
				Response: types.Response{
					Code:    0,
					Message: "success",
				},
				TotalPoint: 0,
			}
			return
		}
		l.Errorf("Failed to get point balance: %v", err)
		return nil, err
	}

	resp = &types.PointBalanceResp{
		Response: types.Response{
			Code:    0,
			Message: "success",
		},
		TotalPoint: userPoint.TotalPoint,
	}

	return
}