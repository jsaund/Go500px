package api

import (
	"net/http"
	"net/url"
)

const (
	consumerKey         = "YOUR CONSUMER KEY HERE"
	getPhotosApi string = "https://api.500px.com/v1/photos?"
)

func GetPhotos() (*http.Request, error) {
	params := url.Values{}
	params.Add("feature", "fresh_today")
	params.Add("consumer_key", consumerKey)

	url := getPhotosApi + params.Encode()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	return req, nil
}
