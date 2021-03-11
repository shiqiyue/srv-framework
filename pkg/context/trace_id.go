package context

import (
	"context"
)

var (
	KEY_TRACE_ID = "c_trace_id"
)

// 设置链路ID变量到上下文中
func SetTraceId(c context.Context, traceId string) context.Context {
	return context.WithValue(c, KEY_TRACE_ID, traceId)
}

// 从上下文中获取链路ID
func GetTraceId(c context.Context) string {
	v, ok := c.Value(KEY_TRACE_ID).(string)
	if !ok {
		return ""
	}
	return v
}
