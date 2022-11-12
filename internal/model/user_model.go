package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
)

const UserCollectionName = "user"

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindOneByAuth(ctx context.Context, auth Auth) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the mongo.
func NewUserModel(conn *monc.Model) UserModel {
	return &customUserModel{
		defaultUserModel: newDefaultUserModel(conn),
	}
}

func (m *customUserModel) FindOneByAuth(ctx context.Context, auth Auth) (*User, error) {
	var data User
	err := m.conn.FindOneNoCache(ctx, &data, bson.M{"auth": auth})
	switch err {
	case nil:
		return &data, nil
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
