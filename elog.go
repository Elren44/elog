package elog

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() *zap.SugaredLogger {
	var sl *zap.SugaredLogger
	encoder := getEncoder()
	fileSyncer := getLogWriter()
	ws := zapcore.NewMultiWriteSyncer(os.Stdout, fileSyncer)
	core := zapcore.NewCore(encoder, ws, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sl = logger.Sugar()
	return sl
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("02 Jan 06 15:04 -0700"))
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = SyslogTimeEncoder

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.FunctionKey = "func"
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "message"
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
