package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderInfoModel = (*customOrderInfoModel)(nil)

type (
	// OrderInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderInfoModel.
	OrderInfoModel interface {
		orderInfoModel
		withSession(session sqlx.Session) OrderInfoModel
		FindByUserId(ctx context.Context, userId int64) ([]*OrderInfo, error)
	}

	customOrderInfoModel struct {
		*defaultOrderInfoModel
	}
)

// NewOrderInfoModel returns a model for the database table.
func NewOrderInfoModel(conn sqlx.SqlConn) OrderInfoModel {
	return &customOrderInfoModel{
		defaultOrderInfoModel: newOrderInfoModel(conn),
	}
}

func (m *customOrderInfoModel) withSession(session sqlx.Session) OrderInfoModel {
	return NewOrderInfoModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customOrderInfoModel) FindByUserId(ctx context.Context, userId int64) ([]*OrderInfo, error) {
	query := fmt.Sprintf("select %s from %s where user_id = ? order by create_time desc", orderInfoRows, m.table)
	var resp []*OrderInfo
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
