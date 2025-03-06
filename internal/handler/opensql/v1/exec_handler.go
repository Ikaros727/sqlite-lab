package v1

import (
	"net/http"

	"github.com/Ikaros727/sqlite-lab/internal/logic/opensql/v1"
	"github.com/Ikaros727/sqlite-lab/internal/svc"
	"github.com/Ikaros727/sqlite-lab/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// ExecHandler
func ExecHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExecReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := v1.NewExecLogic(r.Context(), svcCtx)
		resp, err := l.Exec(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
