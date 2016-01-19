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

type GetPhotoCallback interface {
	OnStart()
	OnError(reason string)
	OnSuccess(response GetPhotoResponse)
}

type GetPhotoRequestBuilderImpl struct {
	baseUrl           string
	pathSubstitutions map[string]string
	queryParams       url.Values
}

func NewGetPhotoRequestBuilder(baseUrl string) GetPhotoRequestBuilder {
	return &GetPhotoRequestBuilderImpl{
		baseUrl:           baseUrl,
		pathSubstitutions: make(map[string]string),
		queryParams:       url.Values{},
	}
}

func (b *GetPhotoRequestBuilderImpl) PhotoID(id string) GetPhotoRequestBuilder {
	b.pathSubstitutions["id"] = string(id)
	return b
}

func (b *GetPhotoRequestBuilderImpl) Comments(include int8) GetPhotoRequestBuilder {
	b.queryParams.Add("comments", string(include))
	return b
}

func (b *GetPhotoRequestBuilderImpl) CommentsPage(page int) GetPhotoRequestBuilder {
	b.queryParams.Add("comments_page", string(page))
	return b
}

func (b *GetPhotoRequestBuilderImpl) ImageSize(size int) GetPhotoRequestBuilder {
	b.queryParams.Add("image_size", string(size))
	return b
}

func (b *GetPhotoRequestBuilderImpl) Tags(tags int8) GetPhotoRequestBuilder {
	b.queryParams.Add("tags", string(tags))
	return b
}

func (b *GetPhotoRequestBuilderImpl) applyPathSubstituions(api string) string {
	if len(b.pathSubstitutions) == 0 {
		return api
	}

	for key, value := range b.pathSubstitutions {
		api = strings.Replace(api, "{"+key+"}", value, -1)
	}

	return api
}

func (b *GetPhotoRequestBuilderImpl) build() (*http.Request, error) {
	req, err := http.NewRequest("GET", b.baseUrl+b.applyPathSubstituions("/photos/{id}"), nil)
	if err != nil {
		return nil, err
	}
	if len(b.queryParams) > 0 {
		req.URL.RawQuery = b.queryParams.Encode()
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (b *GetPhotoRequestBuilderImpl) Run() (GetPhotoResponse, error) {
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
	result, err := NewGetPhotoResponse(response.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *GetPhotoRequestBuilderImpl) RunAsync(callback GetPhotoCallback) {
	if callback != nil {
		callback.OnStart()
	}

	go func(b *GetPhotoRequestBuilderImpl) {
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
