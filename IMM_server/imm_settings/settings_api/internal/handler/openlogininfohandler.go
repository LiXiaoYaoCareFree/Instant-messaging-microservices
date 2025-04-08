package handler

import (
	"IMM_server/common/response"
	"IMM_server/imm_settings/settings_api/internal/logic"
	"IMM_server/imm_settings/settings_api/internal/svc"
	"net/http"
)

func open_login_infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewOpen_login_infoLogic(r.Context(), svcCtx)
		resp, err := l.Open_login_info()
		response.Response(r, w, resp, err)

	}
}
