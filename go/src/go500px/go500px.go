package go500px

import (
	"fmt"
	"net/http"

	"github.com/dghubble/oauth1"
)

const (
	consumerKey      = "8C6ImXPi4dKEnOWC3YwPnKQO1QIYbqaystDCsijC"
	consumerSecret   = "FJMl9qrMv2b96B103TYZbxkJpH7DsHkzWinAfUIg"
	fiveHundredPxAPI = "https://api.500px.com/v1"
)

var client *http.Client
var fiveHundredPxConfig = oauth1.Config{
	ConsumerKey:    consumerKey,
	ConsumerSecret: consumerSecret,
	CallbackURL:    "",
	Endpoint: oauth1.Endpoint{
		RequestTokenURL: fiveHundredPxAPI + "/oauth/request_token",
		AccessTokenURL:  fiveHundredPxAPI + "/oauth/access_token",
		AuthorizeURL:    fiveHundredPxAPI + "/oauth/authorize",
	},
}

func init() {
	client = &http.Client{}
}

func getClient() *http.Client {
	return client
}

func Login(username, password string, callback LoginCallback) {
	login(username, password, fiveHundredPxConfig, callback)
}

func Session(token, secret string) error {
	if token == "" || secret == "" {
		return fmt.Errorf("Invalid token or secret")
	}
	clientToken := oauth1.NewToken(token, secret)
	client = fiveHundredPxConfig.Client(oauth1.NoContext, clientToken)
	return nil
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
