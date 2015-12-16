package go500px

import "net/http"

func GetPhotos(builder *GetPhotosBuilder, consumerKey string) (*Photos, error) {
	request, err := builder.build()
	if err != nil {
		return nil, err
	}

	queryParams := request.URL.Query()
	queryParams.Add("consumer_key", consumerKey)
	request.URL.RawQuery = queryParams.Encode()

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	photosResponse, err := NewGetPhotosResponse(response.Body)
	if err != nil {
		return nil, err
	}

	return photosResponse.GetPhotos(), nil
}
