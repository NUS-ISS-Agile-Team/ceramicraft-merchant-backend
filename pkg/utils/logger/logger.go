package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	sugaredLogger *zap.SugaredLogger
	once          sync.Once
)

// 默认日志等级DEBUG
// 默认日志输出路径: 根目录下test.log
func GetLogger() *zap.SugaredLogger {
	once.Do(func() {
		writeSyncer := getLogWriter()
		encoder := getEncoder()
		core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
		sugaredLogger = zap.New(core, zap.AddCaller()).Sugar()
	})
	return sugaredLogger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	cwd, _ := os.Getwd()
	logPath := filepath.Join(cwd, "test.log")
	fmt.Println(logPath)
	file, _ := os.Create(logPath)
	return zapcore.AddSync(file)
}
