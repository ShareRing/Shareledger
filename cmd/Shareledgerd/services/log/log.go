package log

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func Init() {
	logger, _ := zap.NewDevelopment()
	Log = logger.Sugar()
}
