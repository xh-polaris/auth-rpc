package svc

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/miniprogram"
	mpConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"github.com/xh-polaris/account-svc/model"
	"github.com/xh-polaris/account-svc/rpc/internal/config"
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
		UserModel: model.NewUserModel(monc.MustNewModel(c.Mongo.Url, c.Mongo.DB, model.UserCollectionName, c.CacheConf)),
		Redis:     c.Redis.NewRedis(),
		MiniProgram: wechat.NewWechat().GetMiniProgram(&mpConfig.Config{
			AppID:     c.MiniProgram.AppID,
			AppSecret: c.MiniProgram.AppSecret,
		}),
	}
}
