package svc

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	mpConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/xh-polaris/account-rpc/internal/config"
	"github.com/xh-polaris/account-rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type A interface {
	miniprogram.MiniProgram
}
type ServiceContext struct {
	Config      config.Config
	UserModel   model.UserModel
	Redis       *redis.Redis
	MiniProgram *miniprogram.MiniProgram
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(monc.MustNewModel(c.Mongo.URL, c.Mongo.DB, model.UserCollectionName, c.CacheConf)),
		Redis:     c.Redis.NewRedis(),
		MiniProgram: wechat.NewWechat().GetMiniProgram(&mpConfig.Config{
			AppID:     c.MiniProgram.AppID,
			AppSecret: c.MiniProgram.AppSecret,
		}),
	}
}
