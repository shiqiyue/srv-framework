package https

import (
	"context"
	ferror "github.com/shinedone/srv-framework/pkg/errors"
	"github.com/shinedone/srv-framework/pkg/jsons"
	"net/http"
)

// 处理错误
func EncodeError(ctx context.Context, w http.ResponseWriter, err error) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	appError, ok := err.(ferror.AppError)
	if !ok {
		appError = ferror.New(err.Error())
	}
	_ = jsons.NewEncoder(w).Encode(appError)
}
