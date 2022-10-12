package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mongo struct {
		URL string
		DB  string
	}
	CacheConf   cache.CacheConf
	MiniProgram struct {
		AppID     string
		AppSecret string
	}
}
