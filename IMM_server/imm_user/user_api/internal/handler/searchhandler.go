package handler

import (
	"IMM_server/common/response"
	"IMM_server/imm_user/user_api/internal/logic"
	"IMM_server/imm_user/user_api/internal/svc"
	"IMM_server/imm_user/user_api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func searchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		response.Response(r, w, resp, err)

	}
}
