package handler

import (
	"IMM_server/common/response"
	"IMM_server/imm_auth/auth_api/internal/logic"
	"IMM_server/imm_auth/auth_api/internal/svc"
	"IMM_server/imm_auth/auth_api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func open_loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OpenLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(r, w, nil, err)
			return
		}

		l := logic.NewOpen_loginLogic(r.Context(), svcCtx)
		resp, err := l.Open_login(&req)
		response.Response(r, w, resp, err)

	}
}
