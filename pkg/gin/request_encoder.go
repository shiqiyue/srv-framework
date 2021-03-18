package gin

import (
	"context"
	kitHttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func EncodeJsonRequest(ctx context.Context, req *http.Request, requestModel interface{}) error {
	return kitHttp.EncodeJSONRequest(ctx, req, requestModel)
}
