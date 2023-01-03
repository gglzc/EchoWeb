package util

import (
	"github.com/brpaz/echozap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	logger, _ := zap.Config{
		Encoding: "json",
		Level: zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{  
			MessageKey: "message",  
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,// INFO
			
			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,
	
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()
	zapLogger,_:=echozap.ZapLogger(logger)
	
}

