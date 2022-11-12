package logic

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/xh-polaris/account-rpc/internal/config"
	"github.com/xh-polaris/account-rpc/internal/errorx"
	model2 "github.com/xh-polaris/account-rpc/internal/model"
	"github.com/xh-polaris/account-rpc/internal/model/mockmodel"
	"github.com/xh-polaris/account-rpc/internal/svc"
	"github.com/xh-polaris/account-rpc/pb"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSignInLogic_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)

	r, err := miniredis.Run()
	assert.Nil(t, err)

	userModel := mockmodel.NewMockUserModel(ctrl)
	svcCtx := &svc.ServiceContext{
		Config:    config.Config{},
		UserModel: userModel,
		Redis:     redis.New(r.Addr()),
	}
	ctx := context.Background()
	l := NewSignInLogic(ctx, svcCtx)

	t.Run("invalid auth type", func(t *testing.T) {
		_, err := l.SignIn(&pb.SignInReq{
			AuthType: "gitlab",
			AuthId:   "12306",
			Password: "",
			Params:   nil,
		})
		assert.Equal(t, errorx.ErrInvalidArgument, err)
	})
	t.Run("auth by phone or email", func(t *testing.T) {
		err := svcCtx.Redis.Hset(VerifyCodeKey, "123@abc.com", "66666")
		assert.Nil(t, err)

		t.Run("no such user", func(t *testing.T) {
			userModel.EXPECT().
				FindOneByAuth(gomock.Any(), model2.Auth{Type: "phone", Value: "12306"}).
				Return(nil, model2.ErrNotFound).
				Times(1)

			_, err := l.SignIn(&pb.SignInReq{
				AuthType: "phone",
				AuthId:   "12306",
				Password: "",
				Params:   nil,
			})
			assert.Equal(t, errorx.ErrNoSuchUser, err)
		})
		t.Run("wrong password", func(t *testing.T) {
			userModel.EXPECT().
				FindOneByAuth(gomock.Any(), model2.Auth{Type: "phone", Value: "12306"}).
				Return(&model2.User{
					ID:       primitive.NewObjectID(),
					UpdateAt: time.Now(),
					CreateAt: time.Now(),
					Password: "$2a$10$KTaZRvmPE2MUfVOhjofOou8UgKAZEIkCftj3//iRFQCOpnAfLiDl2",
					Auth:     []model2.Auth{{"phone", "12306"}},
				}, nil).
				Times(2)

			_, err := l.SignIn(&pb.SignInReq{
				AuthType: "phone",
				AuthId:   "12306",
				Password: "",
				Params:   nil,
			})
			assert.Equal(t, errorx.ErrWrongPassword, err)
			_, err = l.SignIn(&pb.SignInReq{
				AuthType: "phone",
				AuthId:   "12306",
				Password: "123",
				Params:   nil,
			})
			assert.Equal(t, errorx.ErrWrongPassword, err)
		})
		t.Run("auth by email and password", func(t *testing.T) {
			userModel.EXPECT().
				FindOneByAuth(gomock.Any(), model2.Auth{Type: "email", Value: "123@abc.com"}).
				Return(&model2.User{
					ID:       primitive.NewObjectID(),
					UpdateAt: time.Now(),
					CreateAt: time.Now(),
					Password: "$2a$10$KTaZRvmPE2MUfVOhjofOou8UgKAZEIkCftj3//iRFQCOpnAfLiDl2",
					Auth:     []model2.Auth{{"email", "123@abc.com"}},
				}, nil).
				Times(1)

			resp, err := l.SignIn(&pb.SignInReq{
				AuthType: "email",
				AuthId:   "123@abc.com",
				Password: "123456Abc.",
				Params:   nil,
			})
			assert.Nil(t, err)
			assert.NotNil(t, resp)
		})
		t.Run("auth by email and verify code", func(t *testing.T) {
			userModel.EXPECT().
				FindOneByAuth(gomock.Any(), model2.Auth{Type: "email", Value: "123@abc.com"}).
				Return(nil, model2.ErrNotFound).
				Times(1)
			userModel.EXPECT().
				Insert(gomock.Any(), gomock.Any()).
				Return(nil).
				Times(1)
			userModel.EXPECT().
				FindOneByAuth(gomock.Any(), model2.Auth{Type: "email", Value: "123@abc.com"}).
				Return(&model2.User{
					ID:       primitive.NewObjectID(),
					UpdateAt: time.Now(),
					CreateAt: time.Now(),
					Password: "$2a$10$vJaijEGmaM4hgMF/55heder6dsEh7B6P8SdMnoDOMbRCJtBv6xD32",
					Auth:     []model2.Auth{{"email", "123@abc.com"}},
				}, nil).
				Times(1)

			resp, err := l.SignIn(&pb.SignInReq{
				AuthType: "email",
				AuthId:   "123@abc.com",
				Password: "123321",
				Params:   []string{"66666"},
			})
			assert.Nil(t, err)
			assert.NotNil(t, resp)
			resp, err = l.SignIn(&pb.SignInReq{
				AuthType: "email",
				AuthId:   "123@abc.com",
				Password: "241312",
				Params:   []string{"66666"},
			})
			assert.Nil(t, err)
			assert.NotNil(t, resp)
		})
	})
	t.Run("auth by wechat", func(t *testing.T) {
		t.Run("invalid options", func(t *testing.T) {
			_, err := l.SignIn(&pb.SignInReq{
				AuthType: "wechat",
				AuthId:   "12138",
				Password: "",
				Params:   nil,
			})
			assert.Equal(t, errorx.ErrInvalidArgument, err)
		})
	})
}
