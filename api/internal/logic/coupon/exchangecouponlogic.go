// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package coupon

import (
	"context"
	"fmt"
	"time"

	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"
	"jhb-api/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ExchangeCouponLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 兑换优惠券
func NewExchangeCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeCouponLogic {
	return &ExchangeCouponLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExchangeCouponLogic) ExchangeCoupon(req *types.ExchangeCouponReq) (resp *types.ExchangeCouponResp, err error) {
	// 使用事务处理兑换逻辑
	err = l.svcCtx.SqlConn.Transact(func(session sqlx.Session) error {
		ctx := l.ctx
		// 1. 获取优惠券信息
		coupon, err := l.svcCtx.CouponModel.FindOne(ctx, req.CouponId)
		if err != nil {
			if err == model.ErrNotFound {
				return fmt.Errorf("优惠券不存在")
			}
			return err
		}

		// 2. 检查库存
		if coupon.Stock <= 0 {
			return fmt.Errorf("优惠券库存不足")
		}

		// 3. 获取用户积分
		userPoint, err := l.svcCtx.UserPointModel.FindOneByUserId(ctx, req.UserId)
		if err != nil {
			if err == model.ErrNotFound {
				return fmt.Errorf("用户积分账户不存在")
			}
			return err
		}

		// 4. 检查积分是否足够
		if userPoint.TotalPoint < coupon.NeedPoint {
			return fmt.Errorf("积分不足，需要 %d 积分", coupon.NeedPoint)
		}

		// 5. 扣减库存
		coupon.Stock--
		err = l.svcCtx.CouponModel.Update(ctx, coupon)
		if err != nil {
			return err
		}

		// 6. 扣减用户积分
		userPoint.TotalPoint -= coupon.NeedPoint
		err = l.svcCtx.UserPointModel.Update(ctx, userPoint)
		if err != nil {
			return err
		}

		// 7. 创建用户优惠券记录
		now := time.Now()
		userCoupon := &model.UserCoupon{
			UserId:       req.UserId,
			CouponId:     req.CouponId,
			Status:       0, // 未使用
			ExchangeTime: now,
		}
		_, err = l.svcCtx.UserCouponModel.Insert(ctx, userCoupon)
		if err != nil {
			return err
		}

		// 8. 插入积分日志
		pointLog := &model.PointLog{
			UserId:      req.UserId,
			ChangePoint: -coupon.NeedPoint,
			ChangeType:  "兑换优惠券",
			RelationId:  userCoupon.Id,
		}
		_, err = l.svcCtx.PointLogModel.Insert(ctx, pointLog)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		l.Errorf("Failed to exchange coupon: %v", err)
		return &types.ExchangeCouponResp{
			Response: types.Response{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	// 获取更新后的积分余额
	userPoint, _ := l.svcCtx.UserPointModel.FindOneByUserId(l.ctx, req.UserId)
	remainingPoints := int64(0)
	if userPoint != nil {
		remainingPoints = userPoint.TotalPoint
	}

	resp = &types.ExchangeCouponResp{
		Response: types.Response{
			Code:    0,
			Message: "兑换成功",
		},
		RemainingPoints: int(remainingPoints),
	}

	return
}
