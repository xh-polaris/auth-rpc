package logic

import (
	"context"
	"github.com/xh-polaris/account-rpc/v3/internal/errorx"
	"github.com/xh-polaris/account-rpc/v3/internal/model"
	"github.com/xh-polaris/account-rpc/v3/internal/svc"
	"github.com/xh-polaris/account-rpc/v3/pb"
	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetPasswordLogic {
	return &SetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetPasswordLogic) SetPassword(in *pb.SetPasswordReq) (*pb.SetPasswordResp, error) {
	userModel := l.svcCtx.UserModel
	user, err := userModel.FindOne(l.ctx, in.UserId)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.ErrNoSuchUser
	default:
		return nil, err
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashPassword)
	err = userModel.Update(l.ctx, user)
	if err != nil {
		return nil, err
	}
	return &pb.SetPasswordResp{}, nil
}
