/*
* CODE GENERATED AUTOMATICALLY WITH GOREST (github.com/jsaund/gorest)
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package go500px

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
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
	baseUrl            string
	pathSubstitutions  map[string]string
	queryParams        url.Values
	postFormParams     url.Values
	postBody           interface{}
	postMultiPartParam map[string][]byte
	headerParams       map[string]string
}

func NewGetPhotoRequestBuilder(baseUrl string) GetPhotoRequestBuilder {
	return &GetPhotoRequestBuilderImpl{
		baseUrl:            baseUrl,
		pathSubstitutions:  make(map[string]string),
		queryParams:        url.Values{},
		postFormParams:     url.Values{},
		postMultiPartParam: make(map[string][]byte),
		headerParams:       make(map[string]string),
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

func (b *GetPhotoRequestBuilderImpl) build() (req *http.Request, err error) {
	url := b.baseUrl + b.applyPathSubstituions("/photos/{id}")
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
