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

type PostCommentCallback interface {
	OnStart()
	OnError(reason string)
	OnSuccess(response PostCommentResponse)
}

type PostCommentRequestBuilderImpl struct {
	pathSubstitutions  map[string]string
	queryParams        url.Values
	postFormParams     url.Values
	postBody           interface{}
	postMultiPartParam map[string][]byte
	headerParams       map[string]string
}

func NewPostCommentRequestBuilder() PostCommentRequestBuilder {
	return &PostCommentRequestBuilderImpl{
		pathSubstitutions:  make(map[string]string),
		queryParams:        url.Values{},
		postFormParams:     url.Values{},
		postMultiPartParam: make(map[string][]byte),
		headerParams:       make(map[string]string),
	}
}

func (b *PostCommentRequestBuilderImpl) PhotoID(id string) PostCommentRequestBuilder {
	b.pathSubstitutions["id"] = fmt.Sprintf("%v", id)
	return b
}

func (b *PostCommentRequestBuilderImpl) Body(body string) PostCommentRequestBuilder {
	b.postFormParams.Add("body", fmt.Sprintf("%v", body))
	return b
}

func (b *PostCommentRequestBuilderImpl) applyPathSubstituions(api string) string {
	if len(b.pathSubstitutions) == 0 {
		return api
	}

	for key, value := range b.pathSubstitutions {
		api = strings.Replace(api, "{"+key+"}", value, -1)
	}

	return api
}

func (b *PostCommentRequestBuilderImpl) build() (req *http.Request, err error) {
	restClient := restclient.GetClient()
	if restClient == nil {
		return nil, fmt.Errorf("A rest client has not been registered yet. You must call client.RegisterClient first")
	}
	url := restClient.BaseURL() + b.applyPathSubstituions("/photos/{id}/comments")
	httpMethod := "POST"
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

func (b *PostCommentRequestBuilderImpl) Run() (PostCommentResponse, error) {
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

	result, err := NewPostCommentResponse(response.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *PostCommentRequestBuilderImpl) RunAsync(callback PostCommentCallback) {
	if callback != nil {
		callback.OnStart()
	}

	go func(b *PostCommentRequestBuilderImpl) {
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
