package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PointRuleModel = (*customPointRuleModel)(nil)

type (
	// PointRuleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPointRuleModel.
	PointRuleModel interface {
		pointRuleModel
		withSession(session sqlx.Session) PointRuleModel
		GetDefaultRule(ctx context.Context) (*PointRule, error)
	}

	customPointRuleModel struct {
		*defaultPointRuleModel
	}
)

// NewPointRuleModel returns a model for the database table.
func NewPointRuleModel(conn sqlx.SqlConn) PointRuleModel {
	return &customPointRuleModel{
		defaultPointRuleModel: newPointRuleModel(conn),
	}
}

func (m *customPointRuleModel) withSession(session sqlx.Session) PointRuleModel {
	return NewPointRuleModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customPointRuleModel) GetDefaultRule(ctx context.Context) (*PointRule, error) {
	query := fmt.Sprintf("select %s from %s where is_default = 1 limit 1", pointRuleRows, m.table)
	var resp PointRule
	err := m.conn.QueryRowCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
