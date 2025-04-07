// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1

package handler

import (
	"net/http"

	"IMM_server/imm_user/user_api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/user/user_info",
				Handler: UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/user/user_info",
				Handler: UserInfoUpdateHandler(serverCtx),
			},
		},
	)
}
