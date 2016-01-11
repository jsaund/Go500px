package go500px

//go:generate $GOPATH/src/github.com/jsaund/gorest/gorest -input get_photos_request_builder.go -output get_photos_request_builder_impl.go -pkg go500px

// @GET("/v1/photos")
type GetPhotosRequestBuilder interface {
	// @QUERY("feature")
	Feature(feature string) GetPhotosRequestBuilder

	// @QUERY("only")
	Only(only string) GetPhotosRequestBuilder

	// @QUERY("exclude")
	Exclude(exclude string) GetPhotosRequestBuilder

	// @QUERY("sort")
	Sort(sort string) GetPhotosRequestBuilder

	// @QUERY("sort_direction")
	SortDirection(direction string) GetPhotosRequestBuilder

	// @QUERY("image_size")
	ImageSize(size string) GetPhotosRequestBuilder

	// @QUERY("tags")
	Tags(tags int8) GetPhotosRequestBuilder

	// @QUERY("consumer_key")
	ConsumerKey(consumerKey string) GetPhotosRequestBuilder

	// @SYNC("GetPhotosResponse")
	Run() (GetPhotosResponse, error)

	// @ASYNC("GetPhotosCallback")
	RunAsync(callback GetPhotosCallback)
}
