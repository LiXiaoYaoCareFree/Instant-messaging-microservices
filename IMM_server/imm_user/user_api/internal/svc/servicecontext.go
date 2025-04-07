package svc

import (
	"IMM_server/imm_user/user_api/internal/config"
	"IMM_server/imm_user/user_rpc/types/user_rpc"
	"IMM_server/imm_user/user_rpc/users"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user_rpc.UsersClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: users.NewUsers(zrpc.MustNewClient(c.UserRpc)),
	}
}
