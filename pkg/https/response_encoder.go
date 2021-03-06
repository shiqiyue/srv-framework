package https

import (
	"context"
	"github.com/shinedone/srv-framework/pkg/jsons"
	"net/http"
)

// 返回结果编码成json
func EncodeJsonResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	err, ok := response.(error)
	if ok && err != nil {
		EncodeError(ctx, w, err)
		return nil
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return jsons.NewEncoder(w).Encode(response)
}
