package go500px

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type getPhotosResponse struct {
	Photos []*photo `json:"photos"`
}

func GetPhotos(consumerKey string) (*Photos, error) {
	queryParams := url.Values{}
	queryParams.Add("consumer_key", consumerKey)
	queryParams.Add("feature", "fresh_today")
	queryParams.Add("sort", "highest_rating")

	api := "https://api.500px.com/v1/photos"
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.URL.RawQuery = queryParams.Encode()

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	photosResponse := &getPhotosResponse{}
	if err = json.NewDecoder(response.Body).Decode(photosResponse); err != nil {
		return nil, err
	}

	photos := NewPhotos()
	photos.photos = photosResponse.Photos

	return photos, nil
}
