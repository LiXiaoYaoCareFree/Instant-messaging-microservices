package svc

import (
	"IMM_server/core"
	"IMM_server/imm_auth/auth_api/internal/config"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := core.InitGorm(c.Mysql.DataSource)
	client := core.InitRedis(c.Redis.Addr, c.Redis.Pwd, c.Redis.DB)
	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
		Redis:  client,
	}
}
