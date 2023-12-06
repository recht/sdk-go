package zap

import (
	"go.temporal.io/sdk/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New returns a Temporal logger backed by go.uber.org/zap.
func New(log *zap.Logger) log.Logger {
	return &zapAdapter{log: log}
}

type zapAdapter struct {
	log *zap.Logger
}

func (z *zapAdapter) Debug(msg string, keyvals ...interface{}) {
	z.log.Debug(msg, z.opts(keyvals)...)
}

func (z *zapAdapter) Info(msg string, keyvals ...interface{}) {
	z.log.Info(msg, z.opts(keyvals)...)
}

func (z *zapAdapter) Warn(msg string, keyvals ...interface{}) {
	z.log.Warn(msg, z.opts(keyvals)...)
}

func (z *zapAdapter) Error(msg string, keyvals ...interface{}) {
	z.log.Error(msg, z.opts(keyvals)...)
}

func (z *zapAdapter) opts(keyvals []any) []zapcore.Field {
	var res []zapcore.Field
	for i := 0; i < len(keyvals); i += 2 {
		res = append(res, zap.Any(keyvals[i].(string), keyvals[i+1]))
	}
	return res
}
