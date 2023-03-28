package initialize

import (
	"fmt"
	"gjm/global"
	"os"
	"path"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 初始化zap，并添加到全局
func InitLogger() {
	cores := GetZapCores()
	global.LOG = zap.New(zapcore.NewTee(cores...), zap.AddCaller())
	global.LOG.Debug("打印debug信息")
	global.LOG.Info("打印info信息")
	global.LOG.Error("打印error信息")
	// 生产模式，是日志文件
}

func GetZapCores() []zapcore.Core {
	var cores []zapcore.Core
	if global.ENV == "development" {
		fmt.Println("构造开发环境zap logger")
		cores = append(cores, GetConsoleCore())
		cores = append(cores, GetFileCores()...)

	} else if global.ENV == "production" {
		fmt.Println("构造生产环境zap logger")
		cores = append(cores, GetFileCores()...)
	}
	return cores
}

// 控制台日志
func GetConsoleCore() zapcore.Core {
	// encoder
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderCfg)
	// WriteSyncer
	writer := os.Stdout
	// LevelEnabler
	// 打印全部
	levelEnabler := zap.DebugLevel
	core := zapcore.NewCore(encoder, writer, levelEnabler)
	return core
}

// 输出文件日志
func GetFileCores() []zapcore.Core {
	var cores []zapcore.Core
	for level := zapcore.InfoLevel; level <= zapcore.FatalLevel; level++ {
		// encoder
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
		encoder := zapcore.NewJSONEncoder(encoderCfg)
		// WriteSyncer
		writer, err := GetWriteSyncer(level.String())
		if err != nil {
			fmt.Println(err)
			return cores
		}
		// LevelEnabler
		// 打印全部
		levelEnabler := GetSingleLevelEnablerFunc(level)
		core := zapcore.NewCore(encoder, writer, levelEnabler)
		cores = append(cores, core)

	}
	return cores
}

// 按日期与level进行文件分割
func GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	now := time.Now()
	date := now.Format("2006-01-02")
	fileName := fmt.Sprintf("%s.log", level)
	rootPath := "log"
	// 创建目录
	// err := os.MkdirAll(path.Join(rootPath, date), os.ModePerm)
	// if err != nil {
	// 	return nil, err
	// }
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(rootPath, date, fileName),
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	})
	return writeSyncer, nil
}

// 是用来生成每一个level的LevelEnabler的
func GetSingleLevelEnablerFunc(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}
