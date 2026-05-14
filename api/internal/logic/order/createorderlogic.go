// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package order

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

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建订单
func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderReq) (resp *types.CreateOrderResp, err error) {
	// 使用事务处理订单创建
	var orderId int64
	var orderNo string
	var usePoint int64
	var remainingPoints int64

	err = l.svcCtx.SqlConn.Transact(func(session sqlx.Session) error {
		ctx := l.ctx
		// 1. 获取商品信息
		goods, err := l.svcCtx.GoodsModel.FindOne(ctx, req.GoodsId)
		if err != nil {
			if err == model.ErrNotFound {
				return fmt.Errorf("商品不存在")
			}
			return err
		}

		// 2. 检查库存
		if goods.Stock <= 0 {
			return fmt.Errorf("商品库存不足")
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
		if userPoint.TotalPoint < goods.NeedPoint {
			return fmt.Errorf("积分不足，需要 %d 积分", goods.NeedPoint)
		}

		// 5. 扣减商品库存
		goods.Stock--
		err = l.svcCtx.GoodsModel.Update(ctx, goods)
		if err != nil {
			return err
		}

		// 6. 扣减用户积分
		usePoint = goods.NeedPoint
		userPoint.TotalPoint -= usePoint
		err = l.svcCtx.UserPointModel.Update(ctx, userPoint)
		if err != nil {
			return err
		}

		remainingPoints = userPoint.TotalPoint

		// 7. 创建订单
		now := time.Now()
		orderNo = fmt.Sprintf("ORD%d%s", now.Unix(), fmt.Sprintf("%06d", now.Nanosecond()/1000))
		orderInfo := &model.OrderInfo{
			OrderNo:    orderNo,
			UserId:     req.UserId,
			GoodsId:    req.GoodsId,
			CouponId:   0, // 默认不使用优惠券
			UsePoint:   usePoint,
			CreateTime: now,
		}
		result, err := l.svcCtx.OrderInfoModel.Insert(ctx, orderInfo)
		if err != nil {
			return err
		}

		orderId, _ = result.LastInsertId()

		// 8. 插入积分日志
		pointLog := &model.PointLog{
			UserId:      req.UserId,
			ChangePoint: -usePoint,
			ChangeType:  "商品消费",
			RelationId:  orderId,
		}
		_, err = l.svcCtx.PointLogModel.Insert(ctx, pointLog)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		l.Errorf("Failed to create order: %v", err)
		return &types.CreateOrderResp{
			Response: types.Response{
				Code:    500,
				Message: err.Error(),
			},
		}, nil
	}

	resp = &types.CreateOrderResp{
		Response: types.Response{
			Code:    0,
			Message: "订单创建成功",
		},
		OrderId:         orderId,
		OrderNo:         orderNo,
		UsePoint:        int(usePoint),
		RemainingPoints: int(remainingPoints),
	}

	return
}
