package middleware

import (
	"fmt"
	"gitee.com/wappyer/golang-backend-template/config"
	"gitee.com/wappyer/golang-backend-template/global"
	"gitee.com/wappyer/golang-backend-template/internal/api/controller"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/errno"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/jwt"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/redis"
	"github.com/gin-gonic/gin"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头中获取角色信息
		role := ctx.Request.Header.Get("User-Role")
		ctx.Set(global.ContextKeyRole, role)

		// 校验token获取用户信息
		token := ctx.Request.Header.Get("Authorization")
		claims, err := jwt.GetClientIns().VerifyToken(token)
		if err == nil {
			// 校验通过保存用户信息
			ctx.Set(global.ContextKeyLoginId, claims.LoginId)
		}

		ctx.Next()
	}
}

// MustLogin 判断登录
func MustLogin(ctx *gin.Context) {
	if ctx.GetString(global.ContextKeyLoginId) == "" {
		controller.FailedByCode(ctx, errno.CodeIllegalToken)
		ctx.Abort()
		return
	}
	ctx.Next()
}

// MustUser 角色必须为用户
func MustUser(ctx *gin.Context) {
	if ctx.GetString(global.ContextKeyRole) != global.RoleUser {
		controller.FailedByCode(ctx, errno.CodeIllegalRoleRequest)
		ctx.Abort()
		return
	}
	ctx.Next()
}

// MustAdmin 角色必须为管理员
func MustAdmin(ctx *gin.Context) {
	if ctx.GetString(global.ContextKeyRole) != global.RoleAdmin {
		controller.FailedByCode(ctx, errno.CodeIllegalRoleRequest)
		ctx.Abort()
		return
	}
	ctx.Next()
}

// MustNotPro 非线上环境
func MustNotPro(ctx *gin.Context) {
	if config.Conf.Env == config.Pro {
		controller.FailedByCode(ctx, errno.CodeIllegalRequest)
		ctx.Abort()
		return
	}
	ctx.Next()
}

// UnReRequest 3s内不可重复请求
func UnReRequest(ctx *gin.Context) {
	userId := ctx.GetString(global.ContextKeyLoginId)
	if userId == "" {
		ctx.Next()
	}

	method := ctx.Request.Method
	url := ctx.Request.URL.Path
	checkRoute := method + url

	key := fmt.Sprintf("%s%s", redis.KeysCommonRepeatRequest, userId)
	val, _ := redis.GetClientIns().Get(ctx, key)
	if val == checkRoute {
		controller.FailedByCode(ctx, errno.CodeRepeatRequest)
		ctx.Abort()
	}
	_ = redis.GetClientIns().Set(ctx, key, checkRoute, time.Second*3)
	ctx.Next()
}
