package middleware

import (
	"github.com/gin-gonic/gin"
	kitLog "github.com/go-kit/kit/log"
	"github.com/shinedone/srv-framework/pkg/context"
)

func LoggerSetter(logger kitLog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context.SetLogger(ctx, logger)
		ctx.Next()
	}
}
