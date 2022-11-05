package logic

import (
	"context"
	"errors"

	"github.com/xh-polaris/account-svc/model"
	"github.com/xh-polaris/account-svc/rpc/errorx"
	"github.com/xh-polaris/account-svc/rpc/internal/svc"
	"github.com/xh-polaris/account-svc/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"golang.org/x/crypto/bcrypt"
)

type SignInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const (
	VerifyCodeKey = "verify_code"
)

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignInLogic) SignIn(in *pb.SignInReq) (resp *pb.SignInResp, err error) {
	resp = &pb.SignInResp{}
	switch in.AuthType {
	case model.EmailAuthType:
		fallthrough
	case model.PhoneAuthType:
		resp.UserId, err = l.signInByPassword(in)
	case model.WechatAuthType:
		resp.UserId, err = l.signInByWechat(in)
	default:
		return nil, errorx.ErrInvalidArgument
	}
	if err != nil {
		return nil, err
	}
	return
}

func (l *SignInLogic) signInByPassword(in *pb.SignInReq) (string, error) {
	userModel := l.svcCtx.UserModel

	// 检查是否设置了验证码，若设置了检查验证码是否合法
	ok, err := l.checkVerifyCode(in.Options, in.AuthValue)
	if err != nil {
		return "", err
	}

	auth := model.Auth{
		Type:  in.AuthType,
		Value: in.AuthValue,
	}
	user, err := userModel.FindOneByAuth(l.ctx, auth)

	switch err {
	case nil:
	case model.ErrNotFound:
		if !ok {
			return "", errorx.ErrNoSuchUser
		}

		user = &model.User{Auth: []model.Auth{auth}}
		err := userModel.Insert(l.ctx, user)
		if err != nil {
			return "", err
		}
		return user.ID.Hex(), nil
	default:
		return "", err
	}

	if ok {
		return user.ID.Hex(), nil
	}

	// 验证码未通过，尝试密码登录
	if user.Password == "" || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)) != nil {
		return "", errorx.ErrWrongPassword
	}

	return user.ID.Hex(), nil
}

func (l *SignInLogic) checkVerifyCode(opts []string, authValue string) (bool, error) {
	verifyCode, err := l.svcCtx.Redis.Hget(VerifyCodeKey, authValue)
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	} else if len(opts) < 1 || verifyCode != opts[0] {
		return false, nil
	} else {
		return true, nil
	}
}

func (l *SignInLogic) signInByWechat(in *pb.SignInReq) (string, error) {
	opts := in.Options
	if len(opts) < 1 {
		return "", errorx.ErrInvalidArgument
	}
	jsCode := opts[0]

	// 向微信开放接口提交临时code
	res, err := l.svcCtx.MiniProgram.GetAuth().Code2SessionContext(l.ctx, jsCode)
	if err != nil {
		return "", err
	} else if res.ErrCode != 0 {
		return "", errors.New(res.ErrMsg)
	}

	userModel := l.svcCtx.UserModel
	auth := model.Auth{
		Type:  in.AuthType,
		Value: in.AuthValue,
	}
	user, err := userModel.FindOneByAuth(l.ctx, auth)
	switch err {
	case nil:
	case model.ErrNotFound:
		user = &model.User{Auth: []model.Auth{auth}}
		err := userModel.Insert(l.ctx, &model.User{})
		if err != nil {
			return "", err
		}
		return user.ID.Hex(), nil
	default:
		return "", err
	}

	return user.ID.Hex(), nil
}
