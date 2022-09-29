package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/xh-polaris/account-svc/rpc/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/crypto/bcrypt"

	"github.com/xh-polaris/account-svc/model"
	"github.com/xh-polaris/account-svc/rpc/internal/svc"
	"github.com/xh-polaris/account-svc/rpc/pb"
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
	case model.EmailLoginType:
		fallthrough
	case model.PhoneLoginType:
		resp.UserId, err = l.signInByPassword(in)
	case model.WechatLoginType:
		resp.UserId, err = l.signInByWechat(in)
	default:
		return nil, errorx.ErrInvalidArgument
	}
	if err != nil {
		return nil, err
	}
	return
}

func (l *SignInLogic) signInByPassword(in *pb.SignInReq) (userId int64, err error) {
	userModel := l.svcCtx.UserModel
	userAuthModel := l.svcCtx.UserAuthModel

	// 检查是否设置了验证码，若设置了检查验证码是否合法
	ok, err := l.checkVerifyCode(in.Options, in.AuthValue)
	if err != nil {
		return
	}
	userAuth, err := userAuthModel.FindOneByAuthTypeAuthValue(l.ctx, in.AuthType, in.AuthValue)

	switch err {
	case nil:
	case sqlx.ErrNotFound:
		if !ok {
			return 0, errorx.ErrNoSuchUser
		}
		result, err := userModel.Insert(l.ctx, &model.User{})
		if err != nil {
			return 0, err
		}
		userId, _ = result.LastInsertId()
		_, err = userAuthModel.Insert(l.ctx, &model.UserAuth{
			UserId:    userId,
			AuthType:  in.AuthType,
			AuthValue: in.AuthValue,
		})
		return userId, err
	default:
		return
	}

	if ok {
		return userAuth.UserId, nil
	}

	// 验证码未通过，尝试密码登录
	user, err := userModel.FindOne(l.ctx, userAuth.UserId)
	if err != nil {
		return
	} else if !user.Password.Valid {
		return 0, errorx.ErrPasswordNotSet
	} else if bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(in.Password)) != nil {
		return 0, errorx.ErrWrongPassword
	}

	return user.Id, nil
}

func (l *SignInLogic) checkVerifyCode(opts []string, authValue string) (bool, error) {
	verifyCode, err := l.svcCtx.Redis.Hget(VerifyCodeKey, authValue)
	if err != nil {
		return false, err
	} else if len(opts) < 1 || verifyCode != opts[0] {
		return false, nil
	} else {
		return true, nil
	}
}

func (l *SignInLogic) signInByWechat(in *pb.SignInReq) (userId int64, err error) {
	opts := in.Options
	if len(opts) < 3 {
		return 0, errorx.ErrInvalidArgument
	}
	appId := opts[0]
	secret := opts[1]
	jsCode := opts[2]

	// 向微信开放接口提交临时code
	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appId, secret, jsCode))
	if err != nil {
		return
	}
	var buffer [512]byte
	n, err := resp.Body.Read(buffer[0:])
	if err != nil {
		return
	}
	respBody := &struct {
		OpenId     string `json:"openid"`
		SessionKey string `json:"session_key"`
		UnionId    string `json:"unionid"`
		ErrCode    int64  `json:"errcode"`
		ErrMsg     string `json:"errmsg"`
	}{}
	err = json.Unmarshal(buffer[0:n], respBody)
	if err != nil {
		return
	} else if respBody.ErrCode != 0 {
		return 0, errors.New(respBody.ErrMsg)
	}

	userAuthModel := l.svcCtx.UserAuthModel
	userContact, err := userAuthModel.FindOneByAuthTypeAuthValue(l.ctx, in.AuthType, respBody.UnionId)
	switch err {
	case nil:
	case model.ErrNotFound:
		userModel := l.svcCtx.UserModel
		result, err := userModel.Insert(l.ctx, &model.User{})
		if err != nil {
			return 0, err
		}
		userId, _ = result.LastInsertId()
		_, err = userAuthModel.Insert(l.ctx, &model.UserAuth{
			UserId:    userId,
			AuthType:  in.AuthType,
			AuthValue: in.AuthValue,
		})
		return userId, err
	default:
		return
	}

	return userContact.UserId, nil
}
