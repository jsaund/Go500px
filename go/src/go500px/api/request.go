package api

import (
	"net/http"
	"net/url"
	"reflect"
)

type ApiRequest interface {
	Get() *http.Request
	AuthenticationRequired() bool
}

type RequestBuilder struct {
	api string
}

func NewRequestBuilder(api string) *RequestBuilder {
	return &RequestBuilder{api}
}

func (b *RequestBuilder) build(v interface{}, method string) (*http.Request, error) {
	apiParams := reflect.ValueOf(v).Elem()
	queryParams := url.Values{}

	for i := 0; i < apiParams.NumField(); i++ {
		if value, ok := apiParams.Field(i).Interface().(string); ok {
			queryParams.Add(apiParams.Type().Field(i).Tag.Get("api_param"), value)
		}
	}

	url := b.api
	if len(queryParams) > 0 {
		url = url + queryParams.Encode()
	}

	req, err := http.NewRequest(method, getPhotosApi, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
}
