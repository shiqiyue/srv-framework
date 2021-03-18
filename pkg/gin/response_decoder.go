package gin

import (
	"context"
	"github.com/shinedone/srv-framework/pkg/json"
	"net/http"
)

func NewSimpleJsonResponseDecoder(ctx context.Context, r *http.Response, requestModel interface{}) func(ctx context.Context, r *http.Response) (interface{}, error) {

	return func(ctx context.Context, r *http.Response) (interface{}, error) {
		err := json.NewDecoder(r.Body).Decode(requestModel)
		if err != nil {
			return nil, err
		}
		return requestModel, nil
	}
}
