/*
* CODE GENERATED AUTOMATICALLY WITH GOREST (github.com/jsaund/gorest)
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package go500px

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/jsaund/gorest/restclient"
)

type GetPhotosCallback interface {
	OnStart()
	OnError(reason string)
	OnSuccess(response GetPhotosResponse)
}

type GetPhotosRequestBuilderImpl struct {
	pathSubstitutions  map[string]string
	queryParams        url.Values
	postFormParams     url.Values
	postBody           interface{}
	postMultiPartParam map[string][]byte
	headerParams       map[string]string
}

func NewGetPhotosRequestBuilder() GetPhotosRequestBuilder {
	return &GetPhotosRequestBuilderImpl{
		pathSubstitutions:  make(map[string]string),
		queryParams:        url.Values{},
		postFormParams:     url.Values{},
		postMultiPartParam: make(map[string][]byte),
		headerParams:       make(map[string]string),
	}
}

func (b *GetPhotosRequestBuilderImpl) ConsumerKey(consumerKey string) GetPhotosRequestBuilder {
	b.queryParams.Add("consumer_key", fmt.Sprintf("%v", consumerKey))
	return b
}

func (b *GetPhotosRequestBuilderImpl) Exclude(exclude string) GetPhotosRequestBuilder {
	b.queryParams.Add("exclude", fmt.Sprintf("%v", exclude))
	return b
}

func (b *GetPhotosRequestBuilderImpl) Feature(feature string) GetPhotosRequestBuilder {
	b.queryParams.Add("feature", fmt.Sprintf("%v", feature))
	return b
}

func (b *GetPhotosRequestBuilderImpl) ImageSize(size string) GetPhotosRequestBuilder {
	b.queryParams.Add("image_size", fmt.Sprintf("%v", size))
	return b
}

func (b *GetPhotosRequestBuilderImpl) Only(only string) GetPhotosRequestBuilder {
	b.queryParams.Add("only", fmt.Sprintf("%v", only))
	return b
}

func (b *GetPhotosRequestBuilderImpl) Sort(sort string) GetPhotosRequestBuilder {
	b.queryParams.Add("sort", fmt.Sprintf("%v", sort))
	return b
}

func (b *GetPhotosRequestBuilderImpl) SortDirection(direction string) GetPhotosRequestBuilder {
	b.queryParams.Add("sort_direction", fmt.Sprintf("%v", direction))
	return b
}

func (b *GetPhotosRequestBuilderImpl) Tags(tags int8) GetPhotosRequestBuilder {
	b.queryParams.Add("tags", fmt.Sprintf("%v", tags))
	return b
}

func (b *GetPhotosRequestBuilderImpl) applyPathSubstituions(api string) string {
	if len(b.pathSubstitutions) == 0 {
		return api
	}

	for key, value := range b.pathSubstitutions {
		api = strings.Replace(api, "{"+key+"}", value, -1)
	}

	return api
}

func (b *GetPhotosRequestBuilderImpl) build() (req *http.Request, err error) {
	restClient := restclient.GetClient()
	if restClient == nil {
		return nil, fmt.Errorf("A rest client has not been registered yet. You must call client.RegisterClient first")
	}
	url := restClient.BaseURL() + b.applyPathSubstituions("/photos")
	httpMethod := "GET"
	switch httpMethod {
	case "POST", "PUT":
		if b.postBody != nil {
			// Assume the body is to be marshalled to JSON
			contentBody, err := json.Marshal(b.postBody)
			if err != nil {
				return nil, err
			}
			contentReader := bytes.NewReader(contentBody)
			req, err = http.NewRequest(httpMethod, url, contentReader)
			if err != nil {
				return nil, err
			}
			req.Header.Set("Content-Type", "application/json")
		} else if len(b.postFormParams) > 0 {
			contentForm := b.postFormParams.Encode()
			contentReader := strings.NewReader(contentForm)
			if req, err = http.NewRequest(httpMethod, url, contentReader); err != nil {
				return nil, err
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		} else if len(b.postMultiPartParam) > 0 {
			contentBody := &bytes.Buffer{}
			writer := multipart.NewWriter(contentBody)
			for key, value := range b.postMultiPartParam {
				if err := writer.WriteField(key, string(value)); err != nil {
					return nil, err
				}
			}
			if err = writer.Close(); err != nil {
				return nil, err
			}
			if req, err = http.NewRequest(httpMethod, url, contentBody); err != nil {
				return nil, err
			}
			req.Header.Set("Content-Type", "multipart/form-data")
		}
	case "GET", "DELETE":
		req, err = http.NewRequest(httpMethod, url, nil)
		if err != nil {
			return nil, err
		}
		if len(b.queryParams) > 0 {
			req.URL.RawQuery = b.queryParams.Encode()
		}
	}
	req.Header.Set("Accept", "application/json")
	for key, value := range b.headerParams {
		req.Header.Set(key, value)
	}
	return req, nil
}

func (b *GetPhotosRequestBuilderImpl) Run() (GetPhotosResponse, error) {
	request, err := b.build()
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = request.URL.Query().Encode()

	restClient := restclient.GetClient()
	if restClient == nil {
		return nil, fmt.Errorf("A rest client has not been registered yet. You must call client.RegisterClient first")
	}

	if restClient.Debug() {
		restclient.DebugRequest(request)
	}

	response, err := restClient.HttpClient().Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if restClient.Debug() {
		restclient.DebugResponse(response)
	}

	return NewGetPhotosResponse(response.Body)
}

func (b *GetPhotosRequestBuilderImpl) RunAsync(callback GetPhotosCallback) {
	if callback != nil {
		callback.OnStart()
	}

	go func(b *GetPhotosRequestBuilderImpl) {
		response, err := b.Run()

		if callback != nil {
			if err != nil {
				callback.OnError(err.Error())
			} else {
				callback.OnSuccess(response)
			}
		}
	}(b)
}
