package middleware

import (
	"github.com/gin-gonic/gin"
	kitLog "github.com/go-kit/kit/log"
	"github.com/shinedone/srv-framework/pkg/contexts"
)

func LoggerSetter(logger kitLog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		contexts.SetLogger(ctx, logger)
		ctx.Next()
	}
}
