package logic

import (
	"context"
	"github.com/xh-polaris/account-svc/api/internal/svc"
	"github.com/xh-polaris/account-svc/api/internal/types"
	"github.com/xh-polaris/account-svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendVerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerifyCodeLogic {
	return &SendVerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendVerifyCodeLogic) SendVerifyCode(req *types.SendVerifyCodeReq) (*types.SendVerifyCodeResp, error) {
	_, err := l.svcCtx.AccountRPC.SendVerifyCode(l.ctx, &pb.SendVerifyCodeReq{
		AuthType:  req.AuthType,
		AuthValue: req.AuthValue,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
