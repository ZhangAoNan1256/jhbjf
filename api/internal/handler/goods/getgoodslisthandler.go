// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package goods

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jhb-api/api/internal/logic/goods"
	"jhb-api/api/internal/svc"
)

// 获取商品列表
func GetGoodsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := goods.NewGetGoodsListLogic(r.Context(), svcCtx)
		resp, err := l.GetGoodsList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
