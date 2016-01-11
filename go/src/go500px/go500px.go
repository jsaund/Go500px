package go500px

func GetPhotos(builder GetPhotosRequestBuilder, consumerKey string) (*Photos, error) {
	response, err := builder.Run()
	if err != nil {
		return nil, err
	}
	return response.GetPhotos(), nil
}

func GetPhotosAsync(builder GetPhotosRequestBuilder, consumerKey string, callback GetPhotosCallback) {
	builder.RunAsync(callback)
}
