package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CouponModel = (*customCouponModel)(nil)

type (
	// CouponModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCouponModel.
	CouponModel interface {
		couponModel
		withSession(session sqlx.Session) CouponModel
		FindAll(ctx context.Context) ([]*Coupon, error)
	}

	customCouponModel struct {
		*defaultCouponModel
	}
)

// NewCouponModel returns a model for the database table.
func NewCouponModel(conn sqlx.SqlConn) CouponModel {
	return &customCouponModel{
		defaultCouponModel: newCouponModel(conn),
	}
}

func (m *customCouponModel) withSession(session sqlx.Session) CouponModel {
	return NewCouponModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customCouponModel) FindAll(ctx context.Context) ([]*Coupon, error) {
	query := fmt.Sprintf("select %s from %s order by id desc", couponRows, m.table)
	var resp []*Coupon
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
