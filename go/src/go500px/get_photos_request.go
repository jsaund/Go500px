package go500px

import (
	"net/http"
	"net/url"
	"reflect"
)

const api string = "/v1/photos"

type GetPhotosBuilder struct {
	apiUrl             string
	FeatureQuery       string `api_param:"feature"`
	OnlyQuery          string `api_param:"only"`
	ExcludeQuery       string `api_param:"exclude"`
	SortQuery          string `api_param:"sort"`
	SortDirectionQuery string `api_param:"sort_direction"`
	ImageSizeQuery     string `api_param:"image_size"`
	TagsQuery          int8   `api_param:"tags"`
}

func NewGetPhotosBuilder(baseUrl string) *GetPhotosBuilder {
	return &GetPhotosBuilder{apiUrl: baseUrl + api}
}

func (builder *GetPhotosBuilder) Feature(feature string) *GetPhotosBuilder {
	builder.FeatureQuery = feature
	return builder
}

func (builder *GetPhotosBuilder) Only(only string) *GetPhotosBuilder {
	builder.OnlyQuery = only
	return builder
}

func (builder *GetPhotosBuilder) Exclude(exclude string) *GetPhotosBuilder {
	builder.ExcludeQuery = exclude
	return builder
}

func (builder *GetPhotosBuilder) Sort(sort string) *GetPhotosBuilder {
	builder.SortQuery = sort
	return builder
}

func (builder *GetPhotosBuilder) SortDirection(sortDirection string) *GetPhotosBuilder {
	builder.SortDirectionQuery = sortDirection
	return builder
}

func (builder *GetPhotosBuilder) ImageSize(imageSize string) *GetPhotosBuilder {
	builder.ImageSizeQuery = imageSize
	return builder
}

func (builder *GetPhotosBuilder) Tags(enabled bool) *GetPhotosBuilder {
	if enabled {
		builder.TagsQuery = int8(1)
	} else {
		builder.TagsQuery = int8(0)
	}
	return builder
}

func (builder *GetPhotosBuilder) build() (*http.Request, error) {
	queryParams := url.Values{}
	apiParams := reflect.ValueOf(builder).Elem()

	for i := 0; i < apiParams.NumField(); i++ {
		valueField := apiParams.Field(i)
		typeField := apiParams.Type().Field(i)
		tag := typeField.Tag

		if !valueField.CanInterface() {
			continue
		}

		key := tag.Get("api_param")
		if key == "" {
			continue
		}

		if valLiteral, ok := valueField.Interface().(string); ok {
			queryParams.Add(key, valLiteral)
		}
	}

	req, err := http.NewRequest("GET", builder.apiUrl, nil)
	if err != nil {
		return nil, err
	}

	if len(queryParams) > 0 {
		req.URL.RawQuery = queryParams.Encode()
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
}
