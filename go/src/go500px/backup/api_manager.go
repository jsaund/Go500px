package go500px

import (
	"net/http"
)

type ApiManager interface {
	Do(apiRequest ApiRequest) (*ApiResponse, error)
}

type apiManagerImpl struct {
	client      *http.Client
	consumerKey string
}

func NewApiManager(consumerKey string) *apiManagerImpl {
	return &apiManagerImpl{
		client:      &http.Client{},
		consumerKey: consumerKey,
	}
}

func (a *apiManagerImpl) Do(apiRequest ApiRequest) (*ApiResponse, error) {
	request := apiRequest.Get()
	queryParams := request.URL.Query()
	if apiRequest.AuthenticationRequired() {
		queryParams.Add("consumer_key", a.consumerKey)
	}
	request.URL.RawQuery = queryParams.Encode()
	response, err := a.client.Do(request)

	if err != nil {
		return nil, err
	}

	return GetResponse(apiRequest, response)
}
