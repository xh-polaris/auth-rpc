package logic

import (
	"context"

	"github.com/xh-polaris/auth-rpc/internal/errorx"
	"github.com/xh-polaris/auth-rpc/internal/model"
	"github.com/xh-polaris/auth-rpc/internal/svc"
	"github.com/xh-polaris/auth-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendVerifyCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerifyCodeLogic {
	return &SendVerifyCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendVerifyCodeLogic) SendVerifyCode(in *pb.SendVerifyCodeReq) (*pb.SendVerifyCodeResp, error) {
	var verifyCode string
	switch in.AuthType {
	case model.PhoneAuthType:
		verifyCode = "1234"
	case model.EmailAuthType:
		verifyCode = "6666"
	default:
		return nil, errorx.ErrInvalidArgument
	}
	err := l.svcCtx.Redis.Hset(VerifyCodeKey, in.AuthId, verifyCode)
	if err != nil {
		return nil, err
	}
	return &pb.SendVerifyCodeResp{}, nil
}
