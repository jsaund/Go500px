package api

import (
	"fmt"
	"net/http"
	"reflect"
)

type ApiResponse interface {
	Get() interface{}
}

var responseHandlerMap map[reflect.Type]func(response *http.Response) (ApiResponse, error)

func registerResponseHandler(request interface{}, response func(response *http.Response) (ApiResponse, error)) {
	if responseHandlerMap == nil {
		responseHandlerMap = make(map[reflect.Type]func(response *http.Response) (ApiResponse, error))
	}
	responseHandlerMap[reflect.TypeOf(request)] = response
}

func GetResponse(request interface{}, response *http.Response) (ApiResponse, error) {
	if f, ok := responseHandlerMap[reflect.TypeOf(request)]; ok {
		return f(response)
	}
	return nil, fmt.Errorf("Unregistered response handler")
}
