// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package coupon

import (
	"context"
	"fmt"

	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"
	"jhb-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCouponsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取我的优惠券
func NewGetUserCouponsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCouponsLogic {
	return &GetUserCouponsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserCouponsLogic) GetUserCoupons(req *types.UserIdReq) (resp *types.UserCouponListResp, err error) {
	// 查询用户的优惠券
	query := fmt.Sprintf("select uc.*, c.coupon_name from user_coupon uc left join coupon c on uc.coupon_id = c.id where uc.user_id = ? order by uc.exchange_time desc")
	
	type UserCouponWithInfo struct {
		model.UserCoupon
		CouponName string `db:"coupon_name"`
	}
	
	var userCoupons []UserCouponWithInfo
	err = l.svcCtx.SqlConn.QueryRowsCtx(l.ctx, &userCoupons, query, req.UserId)
	if err != nil {
		l.Errorf("Failed to get user coupons: %v", err)
		return nil, err
	}

	var list []types.UserCoupon
	for _, uc := range userCoupons {
		useTime := ""
		if uc.UseTime.Valid {
			useTime = uc.UseTime.Time.Format("2006-01-02 15:04:05")
		}
		
		list = append(list, types.UserCoupon{
			Id:           uc.Id,
			CouponName:   uc.CouponName,
			Status:       int(uc.Status),
			ExchangeTime: uc.ExchangeTime.Format("2006-01-02 15:04:05"),
			UseTime:      useTime,
		})
	}

	resp = &types.UserCouponListResp{
		Response: types.Response{
			Code:    0,
			Message: "success",
		},
		List: list,
	}

	return
}