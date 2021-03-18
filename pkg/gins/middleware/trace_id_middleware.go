package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shinedone/srv-framework/pkg/contexts"
)

func TraceIdSetter(traceIdHeaderName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		contexts.SetTraceId(ctx, ctx.GetHeader(traceIdHeaderName))
		ctx.Next()
	}
}
