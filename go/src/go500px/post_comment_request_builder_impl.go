/*
* CODE GENERATED AUTOMATICALLY WITH GOREST (github.com/jsaund/gorest)
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package go500px

import (
	"net/http"
	"net/url"
	"strings"
)

type PostCommentCallback interface {
	OnStart()
	OnError(reason string)
	OnSuccess(response PostCommentResponse)
}

type PostCommentRequestBuilderImpl struct {
	baseUrl           string
	pathSubstitutions map[string]string
	queryParams       url.Values
	postFormParams    url.Values
}

func NewPostCommentRequestBuilder(baseUrl string) PostCommentRequestBuilder {
	return &PostCommentRequestBuilderImpl{
		baseUrl:           baseUrl,
		pathSubstitutions: make(map[string]string),
		queryParams:       url.Values{},
		postFormParams:    url.Values{},
	}
}

func (b *PostCommentRequestBuilderImpl) PhotoID(id string) PostCommentRequestBuilder {
	b.pathSubstitutions["id"] = string(id)
	return b
}

func (b *PostCommentRequestBuilderImpl) Body(body string) PostCommentRequestBuilder {
	b.postFormParams.Add("body", string(body))
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

func (b *PostCommentRequestBuilderImpl) build() (*http.Request, error) {
	req, err := http.NewRequest("POST", b.baseUrl+b.applyPathSubstituions("/photos/{id}/comment"), nil)
	if err != nil {
		return nil, err
	}
	if len(b.queryParams) > 0 {
		req.URL.RawQuery = b.queryParams.Encode()
	}
	if len(b.postFormParams) > 0 {
		req.URL.RawQuery = b.postFormParams.Encode()
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (b *PostCommentRequestBuilderImpl) Run() (PostCommentResponse, error) {
	request, err := b.build()
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = request.URL.Query().Encode()

	response, err := getClient().Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
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
