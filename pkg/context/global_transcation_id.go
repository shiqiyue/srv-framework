package context

import (
	"context"
)

var (
	KEY_GLOBAL_TRANSCATION_ID = "c_global_transcation_id"
)

// 设置全局事务ID变量到上下文中
func SetGlobalTranscationId(c context.Context, traceId string) context.Context {
	return context.WithValue(c, KEY_GLOBAL_TRANSCATION_ID, traceId)
}

// 从上下文中获取全局事务ID
func GetGlobalTranscationId(c context.Context) string {
	v, ok := c.Value(KEY_GLOBAL_TRANSCATION_ID).(string)
	if !ok {
		return ""
	}
	return v
}
