package handler

import (
	"net/http"

	"live/internal/logic"
	"live/internal/svc"
	"live/internal/types"
	"live/pkg/code/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LiveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLiveLogic(r.Context(), svcCtx)
		resp, err := l.Live(&req)
		response.Response(r, w, resp, err)
	}
}
