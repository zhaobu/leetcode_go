package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log       *zap.SugaredLogger //printf风格
	tlog      *Logger            //structured 风格
	LogFormat = "2006/01/02 15:04:05.000"
)

type Config struct {
	Logdir     string `mapstructure:"log_dir" json:"log_dir"`
	LogName    string `mapstructure:"log_name" json:"log_name"`
	Loglevel   string `mapstructure:"log_level" json:"log_level"`
	Debug      bool   `mapstructure:"debug" json:"debug"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size"`       //每个日志文件保存的最大尺寸 单位：M
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups"` //最多保留备份个数
	MaxAge     int    `mapstructure:"max_age" json:"max_age"`         //文件最多保存多少天
	Compress   bool   `mapstructure:"compress" json:"compress"`       //是否压缩
	AddCaller  bool   `mapstructure:"add_caller" json:"add_caller"`   //是否打印输出位置
}

//增加init,防止go test时日志报错
func init() {
	tlog = NewLogger(&Config{Debug: true, Loglevel: "debug", Compress: true, AddCaller: true}, 1)
	log = tlog.Sugar()
}

func InitLog(conf *Config) *Logger {
	tlog := NewLogger(conf, 1)
	log = tlog.Sugar()
	return tlog
}

func GetLogLevelStr() string {
	return tlog.Config.Loglevel
}

func GetLogLevel() zapcore.Level {
	return tlog.Level
}

//fmt.Sprint to construct and log a message
func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func DPanic(args ...interface{}) {
	log.DPanic(args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

//fmt.Sprintf to log a templated message
func Debugf(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	log.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	log.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	log.Fatalf(template, args...)
}

//key value形式打印
func Debugw(msg string, keysAndValues ...interface{}) {
	log.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	log.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	log.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	log.Errorw(msg, keysAndValues...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	log.DPanicw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	log.Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	log.Fatalw(msg, keysAndValues...)
}

//structured 风格打印
func TDebug(msg string, fields ...zap.Field) {
	tlog.Debug(msg, fields...)
}

func TInfo(msg string, fields ...zap.Field) {
	tlog.Info(msg, fields...)
}

func TWarn(msg string, fields ...zap.Field) {
	tlog.Warn(msg, fields...)
}

func TError(msg string, fields ...zap.Field) {
	tlog.Error(msg, fields...)
}

func TDPanic(msg string, fields ...zap.Field) {
	tlog.DPanic(msg, fields...)
}

func TPanic(msg string, fields ...zap.Field) {
	tlog.Panic(msg, fields...)
}

func TFatal(msg string, fields ...zap.Field) {
	tlog.Fatal(msg, fields...)
}
