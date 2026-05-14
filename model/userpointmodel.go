package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserPointModel = (*customUserPointModel)(nil)

type (
	// UserPointModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPointModel.
	UserPointModel interface {
		userPointModel
		withSession(session sqlx.Session) UserPointModel
	}

	customUserPointModel struct {
		*defaultUserPointModel
	}
)

// NewUserPointModel returns a model for the database table.
func NewUserPointModel(conn sqlx.SqlConn) UserPointModel {
	return &customUserPointModel{
		defaultUserPointModel: newUserPointModel(conn),
	}
}

func (m *customUserPointModel) withSession(session sqlx.Session) UserPointModel {
	return NewUserPointModel(sqlx.NewSqlConnFromSession(session))
}
