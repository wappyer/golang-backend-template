package application

import (
	"context"
	"gitee.com/wappyer/golang-backend-template/internal/application/appvo"
	"gitee.com/wappyer/golang-backend-template/internal/domain/service"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/errno"
)

type ExampleApplication struct {
	ExampleService *service.ExampleService
}

func NewExampleApplication() *ExampleApplication {
	return &ExampleApplication{
		ExampleService: service.NewExampleService(),
	}
}

func (a *ExampleApplication) Example(ctx context.Context, param appvo.ExampleReq) (appvo.ExampleResp, errno.Errno) {
	token, eno := a.ExampleService.Example(ctx, param.Phone, param.Password)
	return appvo.ExampleResp{
		Token: token,
	}, eno
}
