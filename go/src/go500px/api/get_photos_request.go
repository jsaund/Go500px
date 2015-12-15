package api

import "net/http"

const (
	getPhotosApi string = "https://api.500px.com/v1/photos"
)

type getPhotos struct {
	request *http.Request
}

type getPhotosBuilder struct {
	*RequestBuilder
	FeatureQuery       string `api_param:"feature"`
	OnlyQuery          string `api_param:"only"`
	ExcludeQuery       string `api_param:"exclude"`
	SortQuery          string `api_param:"sort"`
	SortDirectionQuery string `api_param:"sort_direction"`
	ImageSizeQuery     string `api_param:"image_size"`
	TagsQuery          uint8  `api_param:"tags"`
}

func GetPhotosBuilder() *getPhotosBuilder {
	return &getPhotosBuilder{
		RequestBuilder: NewRequestBuilder(getPhotosApi),
	}
}

func (builder *getPhotosBuilder) Feature(feature string) *getPhotosBuilder {
	builder.FeatureQuery = feature
	return builder
}

func (builder *getPhotosBuilder) Only(only string) *getPhotosBuilder {
	builder.OnlyQuery = only
	return builder
}

func (builder *getPhotosBuilder) Exclude(exclude string) *getPhotosBuilder {
	builder.ExcludeQuery = exclude
	return builder
}

func (builder *getPhotosBuilder) Sort(sort string) *getPhotosBuilder {
	builder.SortQuery = sort
	return builder
}

func (builder *getPhotosBuilder) SortDirection(sortDirection string) *getPhotosBuilder {
	builder.SortDirectionQuery = sortDirection
	return builder
}

func (builder *getPhotosBuilder) ImageSize(imageSize string) *getPhotosBuilder {
	builder.ImageSizeQuery = imageSize
	return builder
}

func (builder *getPhotosBuilder) Tags(enabled bool) *getPhotosBuilder {
	if enabled {
		builder.TagsQuery = uint8(1)
	} else {
		builder.TagsQuery = uint8(0)
	}
	return builder
}

func (builder *getPhotosBuilder) Build() (ApiRequest, error) {
	var err error
	if req, err := builder.RequestBuilder.build(builder, "GET"); err == nil {
		getPhotosRequest := &getPhotos{
			request: req,
		}
		registerResponseHandler(getPhotosRequest, GetPhotosResponse)
		return getPhotosRequest, nil
	}
	return nil, err
}

func (api *getPhotos) AuthenticationRequired() bool {
	return true
}

func (api *getPhotos) Get() *http.Request {
	return api.request
}
