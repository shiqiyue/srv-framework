package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shinedone/srv-framework/pkg/context"
)

func TraceIdSetter(traceIdHeaderName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context.SetTraceId(ctx, ctx.GetHeader(traceIdHeaderName))
		ctx.Next()
	}
}
