package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shinedone/srv-framework/pkg/contexts"
)

func GlobalTranscationIdSetter(globalTranscationIdHeaderName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		contexts.SetGlobalTranscationId(ctx, ctx.GetHeader(globalTranscationIdHeaderName))
		ctx.Next()
	}
}
