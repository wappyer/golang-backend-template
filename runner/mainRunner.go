package runner

import (
	"context"
	"errors"
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
	Conf   config.Config
	Server *http.Server
}

func NewMainRunner(conf config.Config) *MainRunner {
	return &MainRunner{
		Conf: conf,
	}
}

func (r *MainRunner) Initialize() error {
	// 初始化数据库
	if err := repository.Initialize(r.Conf.Db); err != nil {
		panic(fmt.Sprintf("[init] repository初始化失败：%s", err))
	}
	// 初始化日志
	logger.Initialize(r.Conf.Log, []string{global.ContextKeyTraceId, global.ContextKeyRole, global.ContextKeyLoginId})
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
	addr := fmt.Sprintf(":%s", r.Conf.Server.Port)
	r.Server = &http.Server{
		Addr:           addr,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   40 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.InfoF(context.Background(), "[init] 启动web服务 listening at %v", addr)
	err := r.Server.ListenAndServe()
	if err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			logger.InfoF(context.Background(), "[init] web服务退出.")
		} else {
			logger.ErrorF(context.Background(), "[init] web服务异常中断: %s.", err)
		}
	}
	return
}

func (r *MainRunner) Shutdown(ctx context.Context) {
	if err := r.Server.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("[init] 退出web服务失败: %s \n", err))
	}
}
