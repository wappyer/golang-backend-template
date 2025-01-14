package middleware

import (
	"context"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/logger"
)

type GinPanicWriter struct {
}

func (w GinPanicWriter) Write(p []byte) (n int, err error) {
	logger.ErrorF(context.Background(), "gin panic recovered :%v", string(p))
	return 0, err
}
