// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package coupon

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jhb-api/api/internal/logic/coupon"
	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"
)

// 兑换优惠券
func ExchangeCouponHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExchangeCouponReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := coupon.NewExchangeCouponLogic(r.Context(), svcCtx)
		resp, err := l.ExchangeCoupon(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
