// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package goods

import (
	"context"
	"fmt"

	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"
	"jhb-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取商品列表
func NewGetGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsListLogic {
	return &GetGoodsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsListLogic) GetGoodsList() (resp *types.GoodsListResp, err error) {
	// 查询所有商品
	query := fmt.Sprintf("select * from goods order by id desc")
	var goods []*model.Goods
	err = l.svcCtx.SqlConn.QueryRowsCtx(l.ctx, &goods, query)
	if err != nil {
		l.Errorf("Failed to get goods list: %v", err)
		return nil, err
	}

	var list []types.GoodsInfo
	for _, g := range goods {
		list = append(list, types.GoodsInfo{
			Id:         g.Id,
			GoodsName:  g.GoodsName,
			NeedPoint:  int(g.NeedPoint),
			Stock:      int(g.Stock),
			CreateTime: g.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &types.GoodsListResp{
		Response: types.Response{
			Code:    0,
			Message: "success",
		},
		List: list,
	}

	return
}