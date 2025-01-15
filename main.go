package main

import (
	"flag"
	"fmt"
	"gitee.com/wappyer/golang-backend-template/config"
	_ "gitee.com/wappyer/golang-backend-template/docs"
	"gitee.com/wappyer/golang-backend-template/runner"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"
)

// @title					golang后端接口项目模版
// @version				1.0
// @description.markdown	api
// @license.name			项目地址
// @license.url			https://gitee.com/wappyer/golang-backend-template.git
// @host					127.0.0.1:8100
// @BasePath				/api
func main() {
	// 通过flag指定运行环境
	flagEnv := ""
	flag.StringVar(&flagEnv, "env", "", "请指定运行环境，如：-env dev")
	flag.Parse()
	if env, ok := config.IsEnv(flagEnv); ok {
		config.Conf.Env = env
	} else {
		panic("请指定运行环境pro/test/dev，如：-env dev")
	}

	// 加载全局配置文件
	_, filename, _, _ := runtime.Caller(0)
	rootPath := path.Dir(filename)
	InitConfig(fmt.Sprintf("%s/config/%s/%s", rootPath, flagEnv, "config.yaml"))

	// 启动服务
	runnersFactory := runner.NewRunnersFactory()
	runnersFactory.RegisterRunner(runner.NewMainRunner(config.Conf))          // web接口服务 runner
	runnersFactory.RegisterRunner(runner.NewDocsRunner(config.Conf.Docs))     // swagger文档服务 runner
	runnersFactory.RegisterRunner(runner.NewPprofRunner(config.Conf.Monitor)) // pprof监控服务 runner
	runnersFactory.Run()

	/**
	 * 优雅退出
	 * SIGHUP	1	Term	终端控制进程结束(终端连接断开)
	 * SIGINT	2	Term	用户发送INTR字符(Ctrl+C)触发
	 * SIGTERM	15	Term	结束程序(可以被捕获、阻塞或忽略)
	 * SIGQUIT	3	Core	用户发送QUIT字符(Ctrl+/)触发
	 * SIGUSR1	30,10,16	Term	用户保留
	 * SIGUSR2	31,12,17	Term	用户保留
	 */
	exitChan := make(chan os.Signal)
	signal.Notify(exitChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	select {
	case s := <-exitChan:
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			runnersFactory.Shutdown()
			break
		default:
			fmt.Printf("exit by other signal: %v", s)
			break
		}
	}
}

// InitConfig 加载配置文件
func InitConfig(filePath string) {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("[init] 缺少配置文件: %s \n", err.Error()))
	}

	// 监听配置文件变化并热加载程序
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("[init] 重载配置文件:%s Op:%s\n", e.Name, e.Op)
		if err := viper.Unmarshal(&config.Conf); err != nil {
			fmt.Printf("[init] 重载配置文件失败：%s \n", err.Error())
		} else {
			fmt.Printf("[init] 重载配置文件成功：%v \n", config.Conf)
		}
	})

	// 加载配置
	if err := viper.Unmarshal(&config.Conf); err != nil {
		panic(fmt.Sprintf("[init] 载入配置文件失败： %s \n", err.Error()))
	}
	fmt.Printf("[init] 载入配置文件成功！配置文件路径：%s \n", filePath)
}

//	@tag.name			login
//	@tag.description	登录相关

//	@tag.name			user
//	@tag.description	用户信息相关

//	@tag.name			info
//	@tag.description	信息配置相关
