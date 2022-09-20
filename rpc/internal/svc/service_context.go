package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/xh-polaris/account-svc/model"
	"github.com/xh-polaris/account-svc/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
	Redis         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(sqlConn),
		UserAuthModel: model.NewUserAuthModel(sqlConn),
		Redis:         c.Redis.NewRedis(),
	}
}
