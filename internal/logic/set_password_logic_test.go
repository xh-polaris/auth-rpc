package logic

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	"github.com/xh-polaris/account-rpc/v3/internal/config"
	"github.com/xh-polaris/account-rpc/v3/internal/errorx"
	model2 "github.com/xh-polaris/account-rpc/v3/internal/model"
	"github.com/xh-polaris/account-rpc/v3/internal/model/mockmodel"
	"github.com/xh-polaris/account-rpc/v3/internal/svc"
	"github.com/xh-polaris/account-rpc/v3/pb"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSetPasswordLogic_SetPassword(t *testing.T) {
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
	l := NewSetPasswordLogic(ctx, svcCtx)

	t.Run("no such user", func(t *testing.T) {
		userModel.EXPECT().
			FindOne(gomock.Any(), "123").
			Return(nil, model2.ErrNotFound).
			Times(1)

		_, err := l.SetPassword(&pb.SetPasswordReq{
			UserId:   "123",
			Password: "",
		})
		assert.Equal(t, errorx.ErrNoSuchUser, err)
	})
	t.Run("valid request", func(t *testing.T) {
		id := primitive.NewObjectID()
		userModel.EXPECT().
			FindOne(gomock.Any(), id.Hex()).
			Return(&model2.User{
				ID:       id,
				UpdateAt: time.Now(),
				CreateAt: time.Now(),
				Password: "",
				Auth:     nil,
			}, nil).
			Times(1)
		userModel.EXPECT().
			Update(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		resp, err := l.SetPassword(&pb.SetPasswordReq{
			UserId:   id.Hex(),
			Password: "",
		})
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
}
