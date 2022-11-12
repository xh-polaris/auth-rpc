// Code generated by goctl. DO NOT EDIT!
// Source: account.proto

package server

import (
	"context"
	logic2 "github.com/xh-polaris/account-rpc/internal/logic"
	"github.com/xh-polaris/account-rpc/internal/svc"
	pb2 "github.com/xh-polaris/account-rpc/pb"
)

type AccountServer struct {
	svcCtx *svc.ServiceContext
	pb2.UnimplementedAccountServer
}

func NewAccountServer(svcCtx *svc.ServiceContext) *AccountServer {
	return &AccountServer{
		svcCtx: svcCtx,
	}
}

func (s *AccountServer) SignIn(ctx context.Context, in *pb2.SignInReq) (*pb2.SignInResp, error) {
	l := logic2.NewSignInLogic(ctx, s.svcCtx)
	return l.SignIn(in)
}

func (s *AccountServer) SetPassword(ctx context.Context, in *pb2.SetPasswordReq) (*pb2.SetPasswordResp, error) {
	l := logic2.NewSetPasswordLogic(ctx, s.svcCtx)
	return l.SetPassword(in)
}

func (s *AccountServer) SendVerifyCode(ctx context.Context, in *pb2.SendVerifyCodeReq) (*pb2.SendVerifyCodeResp, error) {
	l := logic2.NewSendVerifyCodeLogic(ctx, s.svcCtx)
	return l.SendVerifyCode(in)
}
