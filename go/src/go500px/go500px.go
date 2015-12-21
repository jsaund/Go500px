package go500px

import "net/http"

type GetPhotosListener interface {
	OnStart()
	OnError(reason string)
	OnSuccess(photos *Photos)
}

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

func GetPhotosAsync(builder *GetPhotosBuilder, consumerKey string, callback GetPhotosListener) {
	go func() {
		callback.OnStart()
		response, err := GetPhotos(builder, consumerKey)
		if err != nil {
			callback.OnError(err.Error())
		} else {
			callback.OnSuccess(response)
		}
	}()
}
