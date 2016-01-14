package go500px

import "net/http"

var client *http.Client

func init() {
	client = &http.Client{}
}

func getClient() *http.Client {
	return client
}

func GetPhotos(builder GetPhotosRequestBuilder) (*Photos, error) {
	response, err := builder.Run()
	if err != nil {
		return nil, err
	}
	return response.GetPhotos(), nil
}

func GetPhotosAsync(builder GetPhotosRequestBuilder, callback GetPhotosCallback) {
	builder.RunAsync(callback)
}
