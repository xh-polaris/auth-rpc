package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/xh-polaris/account-svc/rpc/pb"
	"time"

	"github.com/xh-polaris/account-svc/api/internal/svc"
	"github.com/xh-polaris/account-svc/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignIn(req *types.SignInReq) (resp *types.SignInResp, err error) {
	rpcResp, err := l.svcCtx.AccountRpc.SignIn(l.ctx, &pb.SignInReq{
		AuthType:  req.AuthType,
		AuthValue: req.AuthValue,
		Password:  req.Password,
		Options:   req.Options,
	})
	if err != nil {
		return
	}
	auth := l.svcCtx.Config.Auth
	resp.AccessToken, resp.AccessExpire, err = generateJwtToken(rpcResp.GetUserId(), auth.AccessSecret, auth.AccessExpire)
	return
}

func generateJwtToken(userId string, secret string, expire int64) (string, int64, error) {
	iat := time.Now().Unix()
	exp := iat + expire
	claims := make(jwt.MapClaims)
	claims["exp"] = exp
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}
	return tokenString, exp, nil
}
