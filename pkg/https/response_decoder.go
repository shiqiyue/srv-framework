package https

import (
	"context"
	"github.com/shinedone/srv-framework/pkg/jsons"
	"net/http"
)

func NewSimpleJsonResponseDecoder(ctx context.Context, r *http.Response, requestModel interface{}) func(ctx context.Context, r *http.Response) (interface{}, error) {

	return func(ctx context.Context, r *http.Response) (interface{}, error) {
		err := jsons.NewDecoder(r.Body).Decode(requestModel)
		if err != nil {
			return nil, err
		}
		return requestModel, nil
	}
}
