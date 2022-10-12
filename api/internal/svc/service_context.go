package svc

import (
	"github.com/xh-polaris/account-svc/api/internal/config"
	"github.com/xh-polaris/account-svc/rpc/account"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	AccountRpc account.Account
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		AccountRpc: account.NewAccount(zrpc.MustNewClient(c.AccountRpc)),
	}
}
