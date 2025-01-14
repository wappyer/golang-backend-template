package runner

import (
	"context"
	"fmt"
	"gitee.com/wappyer/golang-backend-template/config"
	"gitee.com/wappyer/golang-backend-template/global"
	"gitee.com/wappyer/golang-backend-template/internal/api/middleware"
	"gitee.com/wappyer/golang-backend-template/internal/api/router"
	"gitee.com/wappyer/golang-backend-template/internal/domain/repository"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/errno"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/logger"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/validate"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MainRunner struct {
	Server *http.Server
}

func NewMainRunner() *MainRunner {
	return &MainRunner{}
}

func (r *MainRunner) Initialize() error {
	// 初始化数据库
	if err := repository.Initialize(config.Conf.Db); err != nil {
		panic(fmt.Sprintf("[init] repository初始化失败：%s", err))
	}
	// 初始化日志
	logger.Initialize(config.Conf.Log, []string{global.ContextKeyTraceId, global.ContextKeyRole, global.ContextKeyLoginId})
	// 初始化验证器
	validate.Initialize()
	// 注册错误码
	errno.Initialize()
	return nil
}

func (r *MainRunner) Run() {

	// 启动web服务（提供外部restful接口）
	r.WebServer()
}

func (r *MainRunner) WebServer() {
	engine := gin.New()

	// 开启跨域
	engine.Use(middleware.Cors())

	// 自定义GinPanicWriter
	engine.Use(gin.RecoveryWithWriter(&middleware.GinPanicWriter{}))

	// 请求信息打印
	//engine.Use(middleware.Logger(r.Conf.Server))

	router.Router(engine)
	addr := fmt.Sprintf(":%s", config.Conf.Server.Port)
	webServer := &http.Server{
		Addr:           addr,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   40 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.InfoF(context.Background(), "[init] 启动web服务 listening at %v", addr)
	err := webServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("[init] 启动web服务失败: %s \n", err))
	}
	r.Server = webServer
}

func (r *MainRunner) Shutdown(ctx context.Context) {
	if err := r.Server.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("[init] 退出web服务失败: %s \n", err))
	}
}
