package runner

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/wappyer/golang-backend-template/config"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

type DocsRunner struct {
	Conf   config.Server
	Server *http.Server
}

func NewDocsRunner(conf config.Server) *DocsRunner {
	return &DocsRunner{
		Conf: conf,
	}
}

func (r *DocsRunner) Initialize() error {
	if !r.Conf.Enable {
		logger.InfoF(context.Background(), "[init] swagger服务未开启")
		return errors.New("swagger服务未开启")
	}
	return nil
}

func (r *DocsRunner) Run() {
	// 启动、监听日志端口
	r.WebServer()
}

func (r *DocsRunner) WebServer() {
	r.Server = &http.Server{}
	if !config.EnvIsPro() {
		docsGinEngine := gin.New()
		docsGinEngine.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		addr := fmt.Sprintf(":%s", r.Conf.Port)
		r.Server = &http.Server{
			Addr:           addr,
			Handler:        docsGinEngine,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   40 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		logger.InfoF(context.Background(), "[init] 启动swagger服务 listening at %v", addr)
		err := r.Server.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				logger.InfoF(context.Background(), "[init] swagger服务退出.")
			} else {
				logger.ErrorF(context.Background(), "[init] swagger服务异常中断: %s.", err)
			}
		}
	}
	return
}

func (r *DocsRunner) Shutdown(ctx context.Context) {
	if r.Server == nil {
		return
	}
	if err := r.Server.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("[init] 退出swagger服务失败: %s \n", err))
	}
}
