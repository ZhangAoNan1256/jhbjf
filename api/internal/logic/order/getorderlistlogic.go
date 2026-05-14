// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package order

import (
	"context"
	"database/sql"
	"fmt"

	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"
	"jhb-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取订单列表
func NewGetOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListLogic {
	return &GetOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderListLogic) GetOrderList(req *types.UserIdReq) (resp *types.OrderListResp, err error) {
	// 查询用户订单，关联商品和优惠券信息
	query := fmt.Sprintf(`
		select o.*, g.goods_name, c.coupon_name 
		from order_info o 
		left join goods g on o.goods_id = g.id 
		left join coupon c on o.coupon_id = c.id 
		where o.user_id = ? 
		order by o.create_time desc
	`)

	type OrderWithInfo struct {
		model.OrderInfo
		GoodsName  sql.NullString `db:"goods_name"`
		CouponName sql.NullString `db:"coupon_name"`
	}

	var orders []OrderWithInfo
	err = l.svcCtx.SqlConn.QueryRowsCtx(l.ctx, &orders, query, req.UserId)
	if err != nil {
		l.Errorf("Failed to get order list: %v", err)
		return nil, err
	}

	var list []types.OrderInfo
	for _, o := range orders {
		list = append(list, types.OrderInfo{
			Id:         o.Id,
			OrderNo:    o.OrderNo,
			UserId:     o.UserId,
			GoodsName:  o.GoodsName.String,
			CouponName: o.CouponName.String,
			UsePoint:   int(o.UsePoint),
			CreateTime: o.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &types.OrderListResp{
		Response: types.Response{
			Code:    0,
			Message: "success",
		},
		List: list,
	}

	return
}
