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
	Env    config.Env
	Conf   config.Server
	Server *http.Server
}

func NewDocsRunner() *DocsRunner {
	return &DocsRunner{}
}

func (r *DocsRunner) Initialize() error {
	if !config.Conf.Docs.Enable {
		logger.InfoF(context.Background(), "[init] swagger服务未开启")
		return errors.New("swagger服务未开启")
	}
	r.Env = config.Conf.Env
	r.Conf = config.Conf.Docs
	return nil
}

func (r *DocsRunner) Run() {
	// 启动、监听日志端口
	r.WebServer()
}

func (r *DocsRunner) WebServer() {
	docsServer := &http.Server{}
	if r.Env != config.Pro {
		docsGinEngine := gin.New()
		docsGinEngine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		addr := fmt.Sprintf(":%s", r.Conf.Port)
		docsServer = &http.Server{
			Addr:           addr,
			Handler:        docsGinEngine,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   40 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		logger.InfoF(context.Background(), "[init] 启动swagger服务 listening at %v", addr)
		err := docsServer.ListenAndServe()
		if err != nil {
			panic(fmt.Sprintf("[init] 启动swagger服务失败: %s \n", err))
		}
		r.Server = docsServer
	}
}

func (r *DocsRunner) Shutdown(ctx context.Context) {
	if r.Server == nil {
		return
	}
	if err := r.Server.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("[init] 退出swagger服务失败: %s \n", err))
	}
}
