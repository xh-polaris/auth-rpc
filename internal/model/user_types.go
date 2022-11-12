package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Auth     []Auth             `bson:"auth,omitempty" json:"auth,omitempty"`
}

type Auth struct {
	Type  string `bson:"type" json:"type"`
	Value string `bson:"value" json:"value"`
}

const (
	EmailAuthType  = "email"
	PhoneAuthType  = "phone"
	WechatAuthType = "wechat"
)
