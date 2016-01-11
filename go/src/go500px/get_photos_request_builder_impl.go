/*
* CODE GENERATED AUTOMATICALLY WITH Go500px
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package go500px

import (
	"net/http"
	"net/url"
)

type GetPhotosCallback interface {
	OnStart()
	OnError(reason string)
	OnSuccess(response GetPhotosResponse)
}

type GetPhotosRequestBuilderImpl struct {
	baseUrl     string
	queryParams url.Values
}

func NewGetPhotosRequestBuilder(baseUrl string) GetPhotosRequestBuilder {
	return &GetPhotosRequestBuilderImpl{
		baseUrl:     baseUrl,
		queryParams: url.Values{},
	}
}

func (b *GetPhotosRequestBuilderImpl) ConsumerKey(consumerKey string) GetPhotosRequestBuilder {
	b.queryParams.Add("consumer_key", string(consumerKey))
	return b
}

func (b *GetPhotosRequestBuilderImpl) Exclude(exclude string) GetPhotosRequestBuilder {
	b.queryParams.Add("exclude", string(exclude))
	return b
}

func (b *GetPhotosRequestBuilderImpl) Feature(feature string) GetPhotosRequestBuilder {
	b.queryParams.Add("feature", string(feature))
	return b
}

func (b *GetPhotosRequestBuilderImpl) ImageSize(size string) GetPhotosRequestBuilder {
	b.queryParams.Add("image_size", string(size))
	return b
}

func (b *GetPhotosRequestBuilderImpl) Only(only string) GetPhotosRequestBuilder {
	b.queryParams.Add("only", string(only))
	return b
}

func (b *GetPhotosRequestBuilderImpl) Sort(sort string) GetPhotosRequestBuilder {
	b.queryParams.Add("sort", string(sort))
	return b
}

func (b *GetPhotosRequestBuilderImpl) SortDirection(direction string) GetPhotosRequestBuilder {
	b.queryParams.Add("sort_direction", string(direction))
	return b
}

func (b *GetPhotosRequestBuilderImpl) Tags(tags int8) GetPhotosRequestBuilder {
	b.queryParams.Add("tags", string(tags))
	return b
}

func (b *GetPhotosRequestBuilderImpl) build() (*http.Request, error) {
	req, err := http.NewRequest("GET", b.baseUrl+"/v1/photos", nil)
	if err != nil {
		return nil, err
	}
	if len(b.queryParams) > 0 {
		req.URL.RawQuery = b.queryParams.Encode()
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (b *GetPhotosRequestBuilderImpl) Run() (GetPhotosResponse, error) {
	request, err := b.build()
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = request.URL.Query().Encode()

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	result, err := NewGetPhotosResponse(response.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (b *GetPhotosRequestBuilderImpl) RunAsync(callback GetPhotosCallback) {
	go func() {
		if callback != nil {
			callback.OnStart()
		}
		response, err := b.Run()

		if callback != nil {
			if err != nil {
				callback.OnError(err.Error())
			} else {
				callback.OnSuccess(response)
			}
		}
	}()
}
