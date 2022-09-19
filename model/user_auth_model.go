package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserAuthModel = (*customUserAuthModel)(nil)

const (
	PhoneType  = "phone"
	EmailType  = "email"
	WechatType = "wechat"
)

type (
	// UserAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthModel.
	UserAuthModel interface {
		userAuthModel
	}

	customUserAuthModel struct {
		*defaultUserAuthModel
	}
)

// NewUserAuthModel returns a model for the database table.
func NewUserAuthModel(conn sqlx.SqlConn) UserAuthModel {
	return &customUserAuthModel{
		defaultUserAuthModel: newUserAuthModel(conn),
	}
}
