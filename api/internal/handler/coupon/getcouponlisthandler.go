// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package coupon

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jhb-api/api/internal/logic/coupon"
	"jhb-api/api/internal/svc"
)

// 获取优惠券列表
func GetCouponListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := coupon.NewGetCouponListLogic(r.Context(), svcCtx)
		resp, err := l.GetCouponList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
