package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shinedone/srv-framework/pkg/context"
)

func GlobalTranscationIdSetter(globalTranscationIdHeaderName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context.SetGlobalTranscationId(ctx, ctx.GetHeader(globalTranscationIdHeaderName))
		ctx.Next()
	}
}
