package v1

import (
	"net/http"

	"github.com/Ikaros727/sqlite-lab/internal/logic/sqlite/v1"
	"github.com/Ikaros727/sqlite-lab/internal/svc"
	"github.com/Ikaros727/sqlite-lab/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// CreateHandler // 执行 CREATE 方法
func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := v1.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
