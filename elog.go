package elog

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//ElogOutput output type interface
type ElogOutput interface {
	getElogString() string
}

//JsonOutput json log format
type jsonOutput struct {
	output string
}

func (s *jsonOutput) getElogString() string {
	return s.output
}

//ConsoleOutput base log format
type consoleOutput struct {
	output string
}

func (s *consoleOutput) getElogString() string {
	return s.output
}

//JsonOutput: use this variable as parameter to InitLogger or InitFileLogger function
var JsonOutput = &jsonOutput{
	output: "json",
}

//ConsoleOutput: use this variable as parameter to InitLogger or InitFileLogger function
var ConsoleOutput = &consoleOutput{
	output: "console",
}

//InitLogger initialize logger with file log, output = elog.ConsoleOutput or elog.JsonOutput. logFilePath - path to log file. If an empty string is entered, default logFilePath = "./logs/logs.log"
func InitFileLogger(t ElogOutput, logFilePath string) *zap.SugaredLogger {
	var sl *zap.SugaredLogger
	encoder := getEncoder(t.getElogString())
	fileSyncer := getLogWriter(logFilePath)
	ws := zapcore.NewMultiWriteSyncer(os.Stdout, fileSyncer)
	core := zapcore.NewCore(encoder, ws, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	sl = logger.Sugar()
	return sl
}

//InitLogger initialize logger witout file log, output = elog.ConsoleOutput or elog.JsonOutput
func InitLogger(t ElogOutput) *zap.SugaredLogger {
	var sl *zap.SugaredLogger
	encoder := getEncoder(t.getElogString())
	ws := zapcore.NewMultiWriteSyncer(os.Stdout)
	core := zapcore.NewCore(encoder, ws, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	sl = logger.Sugar()
	return sl
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("02 Jan 06 15:04 -0700"))
}

func getEncoder(t string) zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = syslogTimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.FunctionKey = "func"
	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.CallerKey = "caller"
	encoderConfig.MessageKey = "message"
	if t == "console" {
		return zapcore.NewConsoleEncoder(encoderConfig)
	} else if t == "json" {
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logFilePath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/logs.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	if logFilePath == "" {
		return zapcore.AddSync(lumberJackLogger)
	}
	lumberJackLogger.Filename = logFilePath
	return zapcore.AddSync(lumberJackLogger)
}
