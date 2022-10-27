package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger(ModuleName string) (sugarLogger *zap.SugaredLogger) {
	writeSyncer := getLogWriter(ModuleName)
	encoder := getEncoder()

	//pe := zap.NewProductionEncoderConfig()
	//consoleEncoder := zapcore.NewConsoleEncoder(pe)

	//core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
	)
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
	return
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func getLogWriter(ModuleName string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("log/%s.log", ModuleName),
		MaxSize:    2,
		MaxBackups: 500,
		MaxAge:     30,
		Compress:   false,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}
