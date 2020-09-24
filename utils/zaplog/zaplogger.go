package zaplog

import (
	"fmt"
	"os"
	"time"

	"github.com/gogf/gf/os/gfile"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Level  zapcore.Level
	Config *Config
	*zap.Logger
}

func NewLogger(conf *Config, skip int) *Logger {
	var allCore []zapcore.Core

	//设置默认值
	if conf.Loglevel == "" {
		conf.Loglevel = "debug"
	}
	if conf.MaxAge <= 0 {
		conf.MaxAge = 7
	}
	if conf.MaxBackups <= 0 {
		conf.MaxBackups = 30
	}
	if conf.MaxSize <= 0 {
		conf.MaxSize = 30
	}
	hook := lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s-%s", conf.Logdir, time.Now().Format("20060102"), conf.LogName), //日志文件路径
		MaxSize:    conf.MaxSize,                                                                      //每个日志文件保存的最大尺寸 单位：M
		MaxBackups: conf.MaxBackups,                                                                   //最多保留备份个数
		MaxAge:     conf.MaxAge,                                                                       //文件最多保存多少天
		Compress:   conf.Compress,                                                                     //是否压缩 disabled by default
	}
	var (
		level      zapcore.Level
		fileWriter zapcore.WriteSyncer
	)
	switch conf.Loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	if conf.Logdir != "" {
		fileWriter = zapcore.AddSync(&hook)
	}

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)

	// for human operators.
	var encoderConfig zapcore.EncoderConfig
	timeFormat := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(LogFormat))
	}
	//如果是debug模式,同时输出到到终端
	if conf.Debug {
		if conf.Logdir != "" && gfile.Exists(conf.Logdir) {
			//重新生成文件
			_, err := os.OpenFile(hook.Filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
			if err != nil {
				fmt.Printf("err os.OpenFile()")
			}
		}
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}

	encoderConfig.EncodeTime = timeFormat
	if nil != fileWriter {
		allCore = append(allCore, zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), fileWriter, level))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if conf.Debug {
		allCore = append(allCore, zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleDebugging, level))
		// allCore = append(allCore, zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), consoleDebugging, level))
	}
	core := zapcore.NewTee(allCore...)

	// From a zapcore.Core, it's easy to construct a Logger.
	if conf.AddCaller {
		return &Logger{Logger: zap.New(core).WithOptions(zap.AddCaller(), zap.AddCallerSkip(skip)), Config: conf, Level: level}
	} else {
		return &Logger{Logger: zap.New(core), Config: conf, Level: level}
	}
}
