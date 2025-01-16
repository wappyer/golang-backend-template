package application

import (
	"context"
	"gitee.com/wappyer/golang-backend-template/internal/application/appvo"
	"gitee.com/wappyer/golang-backend-template/internal/domain/service"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/errno"
)

type AdminApplication struct {
	AdminService *service.AdminService
}

func NewAdminApplication() *AdminApplication {
	return &AdminApplication{
		AdminService: service.NewAdminService(),
	}
}

func (a *AdminApplication) Admin(ctx context.Context, param appvo.LoginReq) (appvo.LoginResp, errno.Errno) {
	token, eno := a.AdminService.Admin(ctx, param.Phone, param.Password)
	return appvo.LoginResp{
		Token: token,
	}, eno
}
