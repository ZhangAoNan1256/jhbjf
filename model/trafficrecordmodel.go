package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TrafficRecordModel = (*customTrafficRecordModel)(nil)

type (
	// TrafficRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTrafficRecordModel.
	TrafficRecordModel interface {
		trafficRecordModel
		withSession(session sqlx.Session) TrafficRecordModel
		FindByUserId(ctx context.Context, userId int64) ([]*TrafficRecord, error)
	}

	customTrafficRecordModel struct {
		*defaultTrafficRecordModel
	}
)

// NewTrafficRecordModel returns a model for the database table.
func NewTrafficRecordModel(conn sqlx.SqlConn) TrafficRecordModel {
	return &customTrafficRecordModel{
		defaultTrafficRecordModel: newTrafficRecordModel(conn),
	}
}

func (m *customTrafficRecordModel) withSession(session sqlx.Session) TrafficRecordModel {
	return NewTrafficRecordModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTrafficRecordModel) FindByUserId(ctx context.Context, userId int64) ([]*TrafficRecord, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? order by `traffic_time` desc", trafficRecordRows, m.table)
	var resp []*TrafficRecord
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
