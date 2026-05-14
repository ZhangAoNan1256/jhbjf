// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package coupon

import (
	"context"

	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCouponListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取优惠券列表
func NewGetCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCouponListLogic {
	return &GetCouponListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCouponListLogic) GetCouponList() (resp *types.CouponListResp, err error) {
	coupons, err := l.svcCtx.CouponModel.FindAll(l.ctx)
	if err != nil {
		l.Errorf("Failed to get coupon list: %v", err)
		return nil, err
	}

	var list []types.CouponInfo
	for _, coupon := range coupons {
		list = append(list, types.CouponInfo{
			Id:             coupon.Id,
			CouponName:     coupon.CouponName,
			DiscountAmount: int(coupon.DiscountAmount),
			NeedPoint:      int(coupon.NeedPoint),
			Stock:          int(coupon.Stock),
			CreateTime:     coupon.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &types.CouponListResp{
		Response: types.Response{
			Code:    0,
			Message: "success",
		},
		List: list,
	}

	return
}