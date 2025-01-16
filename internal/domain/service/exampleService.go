package service

import (
	"context"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/errno"
)

type ExampleService struct {
}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}

func (s *ExampleService) Example(ctx context.Context, phone, password string) (string, errno.Errno) {
	return "token", errno.Errno{}
}
