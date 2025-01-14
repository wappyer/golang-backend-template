package router

import (
	"gitee.com/wappyer/golang-backend-template/internal/api/controller"
	mw "gitee.com/wappyer/golang-backend-template/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func Router(engine *gin.Engine) {
	apiRouter := engine.Group("/api")
	// 统一开启token验证 `登录校验`，`重复请求限制`等在 middleware.SkipAuthRoute 中配置
	apiRouter.Use(mw.Auth())

	// 用户相关
	adminCtrl := controller.NewAdminController()
	{
		apiRouter.POST("/admin/login", mw.MustUser, adminCtrl.Login) // 管理员登录
	}
}
