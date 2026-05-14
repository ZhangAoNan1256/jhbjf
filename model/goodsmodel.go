package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GoodsModel = (*customGoodsModel)(nil)

type (
	// GoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGoodsModel.
	GoodsModel interface {
		goodsModel
		withSession(session sqlx.Session) GoodsModel
		FindAll(ctx context.Context) ([]*Goods, error)
	}

	customGoodsModel struct {
		*defaultGoodsModel
	}
)

// NewGoodsModel returns a model for the database table.
func NewGoodsModel(conn sqlx.SqlConn) GoodsModel {
	return &customGoodsModel{
		defaultGoodsModel: newGoodsModel(conn),
	}
}

func (m *customGoodsModel) withSession(session sqlx.Session) GoodsModel {
	return NewGoodsModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customGoodsModel) FindAll(ctx context.Context) ([]*Goods, error) {
	query := fmt.Sprintf("select %s from %s order by id desc", goodsRows, m.table)
	var resp []*Goods
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
