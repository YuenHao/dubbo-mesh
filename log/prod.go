// +build prod

//在生产部署的日志打印，通过条件编译直接将debug置空
package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logLevel levelType = InfoLevel
)

func defaultLogger() *zap.Logger {

	errSync := zapcore.AddSync(os.Stderr)
	outSync := zapcore.AddSync(os.Stdout)
	core := zapcore.NewTee(
		zapcore.NewCore(defaultConsoleEncoder(), errSync, zap.LevelEnablerFunc(zapErrEnable)),
		zapcore.NewCore(defaultConsoleEncoder(), outSync, zap.LevelEnablerFunc(zapOutEnable)),
	)
	return zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel))
}

func DebugEnable() bool {
	return false
}

func Debug(args ...interface{}) {}

func Debugf(template string, args ...interface{}) {}
