// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"jhb-api/api/internal/config"
	"jhb-api/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config             config.Config
	SqlConn            sqlx.SqlConn
	SysUserModel       model.SysUserModel
	TrafficRecordModel model.TrafficRecordModel
	UserPointModel     model.UserPointModel
	PointLogModel      model.PointLogModel
	PointRuleModel     model.PointRuleModel
	CouponModel        model.CouponModel
	UserCouponModel    model.UserCouponModel
	GoodsModel         model.GoodsModel
	OrderInfoModel     model.OrderInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:             c,
		SqlConn:            conn,
		SysUserModel:       model.NewSysUserModel(conn),
		TrafficRecordModel: model.NewTrafficRecordModel(conn),
		UserPointModel:     model.NewUserPointModel(conn),
		PointLogModel:      model.NewPointLogModel(conn),
		PointRuleModel:     model.NewPointRuleModel(conn),
		CouponModel:        model.NewCouponModel(conn),
		UserCouponModel:    model.NewUserCouponModel(conn),
		GoodsModel:         model.NewGoodsModel(conn),
		OrderInfoModel:     model.NewOrderInfoModel(conn),
	}
}
