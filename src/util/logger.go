package util

import(
	"go.uber.org/zap"
)

// 性能高，用于中间件
var Logger *zap.Logger

// 普通打印
var Sugar *zap.SugaredLogger

func init() {
	Logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	Sugar = Logger.Sugar()
}
