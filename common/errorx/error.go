package errorx

import "google.golang.org/grpc/status"

var (
	ErrNoSuchUser      = status.Error(10001, "no such user")
	ErrWrongVerifyCode = status.Error(10002, "wrong verify code")
	ErrInvalidArgument = status.Error(10003, "invalid argument")
	ErrWrongPassword   = status.Error(10004, "wrong password")
)
