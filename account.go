package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/xh-polaris/account-rpc/v3/internal/config"
	"github.com/xh-polaris/account-rpc/v3/internal/server"
	"github.com/xh-polaris/account-rpc/v3/internal/svc"
	"github.com/xh-polaris/account-rpc/v3/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var configFile = flag.String("f", "etc/account.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterAccountServer(grpcServer, server.NewAccountServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	s.AddUnaryInterceptors(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			se, ok := status.FromError(err)
			if !ok {
				logx.Error(se)
			}
		}
		return
	})
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
