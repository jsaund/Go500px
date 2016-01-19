package go500px

//go:generate $GOPATH/src/github.com/jsaund/gorest/gorest -input get_photo_request_builder.go -output get_photo_request_builder_impl.go -pkg go500px

// @GET("/photos/{id}")
type GetPhotoRequestBuilder interface {
	// @PATH("id")
	PhotoID(id string) GetPhotoRequestBuilder

	// @QUERY("image_size")
	ImageSize(size int) GetPhotoRequestBuilder

	// @QUERY("comments")
	Comments(include int8) GetPhotoRequestBuilder

	// @QUERY("comments_page")
	CommentsPage(page int) GetPhotoRequestBuilder

	// @QUERY("tags")
	Tags(tags int8) GetPhotoRequestBuilder

	// @SYNC("GetPhotoResponse")
	Run() (GetPhotoResponse, error)

	// @ASYNC("GetPhotoCallback")
	RunAsync(callback GetPhotoCallback)
}
