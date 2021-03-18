package https

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"github.com/shinedone/srv-framework/pkg/jsons"
	"net/http"
)

// 创建一般的post请求解码
func NewSimpleJsonRequestDecoder(ctx context.Context, r *http.Request, requestModel interface{}) func(ctx context.Context, r *http.Request) (interface{}, error) {

	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		err := jsons.NewDecoder(r.Body).Decode(requestModel)
		if err != nil {
			return nil, err
		}
		return requestModel, nil
	}
}

// 什么都不做的请求解码
func DoNothingRequestDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

// 创建一般的form请求解码
func NewSimpleFormRequestDecoder(ctx context.Context, r *http.Request, requestModel interface{}) func(ctx context.Context, r *http.Request) (interface{}, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		form := r.Form
		err := mapstructure.Decode(form, requestModel)
		if err != nil {
			return nil, err
		}
		return requestModel, nil
	}
}
