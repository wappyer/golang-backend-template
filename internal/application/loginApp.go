package application

import (
	"context"
	"gitee.com/wappyer/golang-backend-template/internal/application/appvo"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/errno"
)

type LoginApplication struct {
}

func NewLoginApplication() *LoginApplication {
	return &LoginApplication{}
}

func (a *LoginApplication) Login(ctx context.Context, param appvo.LoginReq) (appvo.LoginResp, errno.Errno) {

	return appvo.LoginResp{}, errno.Errno{}
}
