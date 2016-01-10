package go500px

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
}

// @CALLBACK("GetPhotoResponse")
type GetPhotosCallback interface {
	OnStart()
	OnError(reason string)
	OnSuccess(response GetPhotoResponse)
}
