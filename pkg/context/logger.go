package context

import (
	"context"
	kitLog "github.com/go-kit/kit/log"
)

var (
	KEY_LOGGER = "c_logger"
)

// 设置日志变量到上下文中
func SetLogger(c context.Context, kLog kitLog.Logger) context.Context {
	return context.WithValue(c, KEY_LOGGER, kLog)
}

// 从上下文中获取日志变量
func GetLogger(c context.Context) kitLog.Logger {
	v, ok := c.Value(KEY_LOGGER).(kitLog.Logger)
	if !ok {
		return nil
	}
	return v
}
