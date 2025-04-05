package handler

import (
	"IMM_server/common/response"
	"IMM_server/imm_auth/auth_api/internal/logic"
	"IMM_server/imm_auth/auth_api/internal/svc"
	"net/http"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthenticationLogic(r.Context(), svcCtx) // 接收客户端请求
		token := r.Header.Get("token")                         // 从请求头中获取 token
		resp, err := l.Authentication(token)                   // 调用逻辑层的 Authentication() 方法处理认证
		response.Response(r, w, resp, err)                     // 返回结果给客户端
	}
}
