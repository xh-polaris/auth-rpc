package logic

import (
	"context"
	"github.com/xh-polaris/account-rpc/v2/internal/errorx"
	"github.com/xh-polaris/account-rpc/v2/internal/model"
	"github.com/xh-polaris/account-rpc/v2/internal/svc"
	"github.com/xh-polaris/account-rpc/v2/pb"
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
	err := l.svcCtx.Redis.Hset(VerifyCodeKey, in.AuthValue, verifyCode)
	if err != nil {
		return nil, err
	}
	return &pb.SendVerifyCodeResp{}, nil
}
