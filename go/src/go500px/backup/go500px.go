package go500px

import "net/http"

type Go500px struct {
	apiManager ApiManager
}

func NewGo500px(consumerKey string) *Go500px {
	return &Go500px{
		apiManager: NewApiManager(consumerKey),
	}
}

func (g *Go500px) GetApiManager() ApiManager {
	return g.apiManager
}
