package go500px

import (
	"fmt"
	"net/http"

	"github.com/dghubble/oauth1"
	"github.com/jsaund/gorest/restclient"
)

const (
	consumerKey      = "8C6ImXPi4dKEnOWC3YwPnKQO1QIYbqaystDCsijC"
	consumerSecret   = "FJMl9qrMv2b96B103TYZbxkJpH7DsHkzWinAfUIg"
	fiveHundredPxAPI = "https://api.500px.com/v1"
	debugHttp        = true
)

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
	restClient := restclient.NewDefaultClient(fiveHundredPxAPI, debugHttp, &http.Client{})
	restclient.RegisterClient(restClient)
}

func Login(username, password string, callback LoginCallback) {
	login(username, password, fiveHundredPxConfig, callback)
}

func Session(token, secret string) error {
	if token == "" || secret == "" {
		return fmt.Errorf("Invalid token or secret")
	}
	clientToken := oauth1.NewToken(token, secret)
	httpClient := fiveHundredPxConfig.Client(oauth1.NoContext, clientToken)
	restClient := restclient.NewDefaultClient(fiveHundredPxAPI, debugHttp, httpClient)
	restclient.RegisterClient(restClient)
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

func GetPhoto(builder GetPhotoRequestBuilder) (GetPhotoResponse, error) {
	return builder.Run()
}

func GetPhotoAsync(builder GetPhotoRequestBuilder, callback GetPhotoCallback) {
	builder.RunAsync(callback)
}

func PostComment(builder PostCommentRequestBuilder) (PostCommentResponse, error) {
	return builder.Run()
}

func PostCommentAsync(builder PostCommentRequestBuilder, callback PostCommentCallback) {
	builder.RunAsync(callback)
}
