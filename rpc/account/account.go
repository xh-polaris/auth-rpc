// Code generated by goctl. DO NOT EDIT!
// Source: account.proto

package account

import (
	"context"

	"github.com/foliet/account/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserReq  = pb.GetUserReq
	GetUserResp = pb.GetUserResp

	Account interface {
		GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
	}

	defaultAccount struct {
		cli zrpc.Client
	}
)

func NewAccount(cli zrpc.Client) Account {
	return &defaultAccount{
		cli: cli,
	}
}

func (m *defaultAccount) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	client := pb.NewAccountClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}
