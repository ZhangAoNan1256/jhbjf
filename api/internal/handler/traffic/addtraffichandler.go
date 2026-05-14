// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package traffic

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jhb-api/api/internal/logic/traffic"
	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"
)

// 添加里程记录
func AddTrafficHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddTrafficReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := traffic.NewAddTrafficLogic(r.Context(), svcCtx)
		resp, err := l.AddTraffic(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
