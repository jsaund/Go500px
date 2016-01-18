package go500px

import (
	"log"

	"github.com/dghubble/oauth1"
)

const (
	xauthUsername = "x_auth_username"
	xauthPassword = "x_auth_password"
	xauthMode     = "x_auth_mode"
	clientMode    = "client_auth"
)

type LoginCallback interface {
	OnStart()
	OnError(reason string)
	OnSuccess(token, secret string)
}

type loginTask struct {
	username string
	password string
	config   oauth1.Config
	callback LoginCallback
}

func login(username, password string, config oauth1.Config, callback LoginCallback) {
	task := loginTask{
		username,
		password,
		config,
		callback,
	}

	callback.OnStart()
	go func(t loginTask) {
		requestToken, requestSecret, err := config.RequestToken()
		if err != nil {
			log.Printf("Failed to request temporary token. Reason: %v", err)
			t.callback.OnError(err.Error())
			return
		}

		_, err = config.AuthorizationURL(requestToken)
		if err != nil {
			log.Printf("Failed to get authorization url. Reason: %v", err)
			t.callback.OnError(err.Error())
			return
		}

		extraOauthParams := map[string]string{
			xauthUsername: t.username,
			xauthPassword: t.password,
			xauthMode:     clientMode,
		}

		accessToken, accessSecret, err := config.AccessToken(requestToken, requestSecret, "", extraOauthParams)
		if err != nil {
			log.Printf("Failed to authorize user. Reason: %v", err)
			t.callback.OnError(err.Error())
			return
		}

		if err = Session(accessToken, accessSecret); err != nil {
			log.Printf("Failed to authorize user. Reason: %v", err)
			t.callback.OnError(err.Error())
			return
		}

		callback.OnSuccess(accessToken, accessSecret)
	}(task)
}
