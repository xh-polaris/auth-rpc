package svc

import (
	"github.com/xh-polaris/account-svc/api/internal/config"
	"github.com/xh-polaris/account-svc/rpc/account"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	AccountRPC account.Account
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		AccountRPC: account.NewAccount(zrpc.MustNewClient(c.AccountRPC)),
	}
}
