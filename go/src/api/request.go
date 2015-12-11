package api

import "net/http"

type ApiRequest interface {
	Get() *http.Request
	AuthenticationRequired() bool
}
