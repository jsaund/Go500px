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
		fmt.Printf("Initialize map!\n")
		responseHandlerMap = make(map[reflect.Type]func(response *http.Response) (ApiResponse, error))
	}
	fmt.Printf("Register %s.\n", reflect.TypeOf(request))
	responseHandlerMap[reflect.TypeOf(request)] = response
}

func GetResponse(request interface{}, response *http.Response) (ApiResponse, error) {
	if f, ok := responseHandlerMap[reflect.TypeOf(request)]; ok {
		fmt.Printf("Run function to get response\n")
		return f(response)
	}
	return nil, fmt.Errorf("Unregistered response handler")
}
