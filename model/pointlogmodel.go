package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PointLogModel = (*customPointLogModel)(nil)

type (
	// PointLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPointLogModel.
	PointLogModel interface {
		pointLogModel
		withSession(session sqlx.Session) PointLogModel
		FindByUserId(ctx context.Context, userId int64) ([]*PointLog, error)
	}

	customPointLogModel struct {
		*defaultPointLogModel
	}
)

// NewPointLogModel returns a model for the database table.
func NewPointLogModel(conn sqlx.SqlConn) PointLogModel {
	return &customPointLogModel{
		defaultPointLogModel: newPointLogModel(conn),
	}
}

func (m *customPointLogModel) withSession(session sqlx.Session) PointLogModel {
	return NewPointLogModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customPointLogModel) FindByUserId(ctx context.Context, userId int64) ([]*PointLog, error) {
	query := fmt.Sprintf("select %s from %s where user_id = ? order by create_time desc", pointLogRows, m.table)
	var resp []*PointLog
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
