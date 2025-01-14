package logger

import (
	"fmt"
	"gitee.com/wappyer/golang-backend-template/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Client struct {
	CtxValue []string // 需记录日志的上下文参数
}

var client = &Client{}

func GetClientIns() *Client {
	return client
}

// Initialize 初始化 log
func Initialize(conf config.LogConfigs, ctxValue []string) {
	client = &Client{CtxValue: ctxValue}

	// 日志打印级别
	logLevel := map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dPanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	if _, ok := logLevel[conf.LogLevel]; !ok {
		panic("错误的日志打印级别配置：%v，请配置debug/info/warn/error/dPanic/panic/fatal其一")
	}

	var coreArr []zapcore.Core
	for k, v := range logLevel {
		// 低于配置级别的日志不打印
		if v < logLevel[conf.LogLevel] {
			continue
		}
		coreArr = append(coreArr, getLogCore(conf, k, v))
	}
	core := zapcore.NewTee(coreArr...)

	// zap.Addcaller() 输出日志打印文件和行数如：logger/logger_test.go:33
	// zap.AddStacktrace(zapcore.WarnLevel) 警告日志以上 显示完整的调用栈
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.WarnLevel))

	// 1. zap.ReplaceGlobals 函数将当前初始化的 logger 替换到全局的 logger,
	// 2. 使用 logger 的时候 直接通过 zap.S().Debugf("xxx") or zap.L().Debug("xxx")
	// 3. 使用 zap.S() 和 zap.L() 提供全局锁，保证一个全局的安全访问logger的方式
	zap.ReplaceGlobals(logger)
	return
}

func getLogCore(conf config.LogConfigs, levelStr string, logLevel zapcore.Level) zapcore.Core {
	// 获取日志输出编码
	encoder := getEncoder(conf.LogFormat)

	ext := path.Ext(conf.LogFileName)
	errorLogFilename := strings.TrimRight(conf.LogFileName, ext) + "-" + levelStr + ext
	writeSyncer := getLogWriter(conf, errorLogFilename) // 日志文件输出配置(文件位置和切割)

	levelFunc := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == logLevel
	})
	return zapcore.NewCore(encoder, writeSyncer, levelFunc)
}

// getEncoder 编码器(如何写入日志)
func getEncoder(logFormat string) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // log 时间格式 例如: 2021-09-11t20:05:54.852+0800
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // 输出level序列化为全大写字符串，如 INFO DEBUG ERROR
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if logFormat == "json" {
		return zapcore.NewJSONEncoder(encoderConfig) // 以json格式写入
	}
	return zapcore.NewConsoleEncoder(encoderConfig) // 以logfmt格式写入
}

// getLogWriter 获取日志输出方式  日志文件 控制台
func getLogWriter(conf config.LogConfigs, fileName string) zapcore.WriteSyncer {
	if conf.LogPath == "" {
		panic("请配置日志文件夹路径")
	}
	// 判断日志路径是否存在，如果不存在就创建
	if exist := IsExist(conf.LogPath); !exist {
		if err := os.MkdirAll(conf.LogPath, os.ModePerm); err != nil {
			panic(fmt.Sprintf("日志输出文件夹'%v'新建失败：%v", conf.LogPath, err))
		}
	}

	// 日志文件 与 日志切割 配置
	if fileName == "" {
		fileName = conf.LogFileName
	}
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(conf.LogPath, fileName), // 日志文件路径
		MaxSize:    conf.LogFileMaxSize,                   // 单个日志文件最大多少 mb
		MaxBackups: conf.LogFileMaxBackups,                // 日志备份数量
		MaxAge:     conf.LogMaxAge,                        // 日志最长保留时间
		Compress:   conf.LogCompress,                      // 是否压缩日志
		LocalTime:  true,                                  // 是否使用本地时间
	}
	if conf.LogStdout {
		// 日志同时输出到控制台和日志文件中
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))
	} else {
		// 日志只输出到日志文件
		return zapcore.AddSync(lumberJackLogger)
	}
}

// IsExist 判断文件或者目录是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
