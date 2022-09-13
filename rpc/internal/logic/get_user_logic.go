package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/crypto/bcrypt"

	"github.com/foliet/account/common/errorx"
	"github.com/foliet/account/model"
	"github.com/foliet/account/rpc/internal/svc"
	"github.com/foliet/account/rpc/pb"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const (
	PhonePasswordType = "phone"
	EmailPasswordType = "email"
	VerifyCodeKey     = "verify_code"
)

var backends = make(map[string]func([]string, *GetUserLogic) (int64, error))

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.GetUserReq) (*pb.GetUserResp, error) {
	if len(in.Options) < 1 {
		return nil, errorx.ErrInvalidArgument
	}
	userId, err := backends[in.Type](in.Options, l)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResp{UserId: userId}, nil
}

func init() {
	backends["password"] = getByPassword
	backends["wechat"] = getByWechat
}

func getByPassword(options []string, l *GetUserLogic) (int64, error) {
	if len(options) < 3 {
		return 0, errorx.ErrNoSuchUser
	}
	passwordType := options[0]
	typeValue := options[1]
	password := options[2]
	userModel := l.svcCtx.UserModel

	var user *model.User
	var err error
	if passwordType == PhonePasswordType {
		user, err = userModel.FindOneByPhone(l.ctx, typeValue)
	} else if passwordType == EmailPasswordType {
		user, err = userModel.FindOneByEmail(l.ctx, typeValue)
	} else {
		return 0, errorx.ErrInvalidArgument
	}

	switch err {
	case nil:
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
			return 0, errorx.ErrWrongPassword
		}
	case sqlx.ErrNotFound:
		if len(options) < 4 {
			return 0, errorx.ErrNoSuchUser
		}
		preVerifyCode := options[3]
		verifyCode, errGet := l.svcCtx.Redis.Hget(VerifyCodeKey, typeValue)
		if errGet != nil {
			return 0, errGet
		}
		if verifyCode != preVerifyCode {
			return 0, errorx.ErrWrongVerifyCode
		}
		hashPassword, errGen := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if errGen != nil {
			return 0, errGen
		}

		if passwordType == PhonePasswordType {
			user = &model.User{
				Phone:    typeValue,
				Email:    "",
				Password: string(hashPassword),
			}
		} else {
			user = &model.User{
				Phone:    "",
				Email:    typeValue,
				Password: string(hashPassword),
			}
		}

		result, errInsert := userModel.Insert(l.ctx, user)
		if errInsert != nil {
			return 0, errInsert
		}
		return result.LastInsertId()
	}

	return user.Id, nil
}

func getByWechat(options []string, l *GetUserLogic) (int64, error) {
	//TODO(feat) Implement wechat login
	return 0, nil
}
