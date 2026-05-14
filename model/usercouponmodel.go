package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserCouponModel = (*customUserCouponModel)(nil)

type (
	// UserCouponModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserCouponModel.
	UserCouponModel interface {
		userCouponModel
		withSession(session sqlx.Session) UserCouponModel
		FindByUserId(ctx context.Context, userId int64) ([]*UserCoupon, error)
	}

	customUserCouponModel struct {
		*defaultUserCouponModel
	}
)

// NewUserCouponModel returns a model for the database table.
func NewUserCouponModel(conn sqlx.SqlConn) UserCouponModel {
	return &customUserCouponModel{
		defaultUserCouponModel: newUserCouponModel(conn),
	}
}

func (m *customUserCouponModel) withSession(session sqlx.Session) UserCouponModel {
	return NewUserCouponModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customUserCouponModel) FindByUserId(ctx context.Context, userId int64) ([]*UserCoupon, error) {
	query := fmt.Sprintf("select %s from %s where user_id = ? order by exchange_time desc", userCouponRows, m.table)
	var resp []*UserCoupon
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
