// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package point

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jhb-api/api/internal/logic/point"
	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"
)

// 积分核算
func CalculatePointHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CalculatePointReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := point.NewCalculatePointLogic(r.Context(), svcCtx)
		resp, err := l.CalculatePoint(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
