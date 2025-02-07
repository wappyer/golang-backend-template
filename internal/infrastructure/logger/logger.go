package logger

import (
	"context"
	"go.uber.org/zap"
)

const (
	TplErrorDatabase = "数据服务错误 error: %s."
)

func withValue(ctx context.Context) *zap.SugaredLogger {
	s := zap.S()
	c := GetClientIns()
	for _, key := range c.CtxValue {
		if value := ctx.Value(key); value != nil {
			s = s.With(zap.Any(key, value))
		}
	}
	return s
}

func DebugF(ctx context.Context, format string, args ...interface{}) {
	withValue(ctx).Debugf(format, args...)
}

func InfoF(ctx context.Context, format string, args ...interface{}) {
	withValue(ctx).Infof(format, args...)
}

func WarnF(ctx context.Context, format string, args ...interface{}) {
	withValue(ctx).Warnf(format, args...)
}

func ErrorF(ctx context.Context, format string, args ...interface{}) {
	withValue(ctx).Errorf(format, args...)
}
