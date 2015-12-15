package go500px

import "go500px/manager"

type Go500px struct {
	apiManager manager.ApiManager
}

func NewGo500px(consumerKey string) *Go500px {
	return &Go500px{
		apiManager: manager.NewApiManager(consumerKey),
	}
}

func (g *Go500px) GetApiManager() manager.ApiManager {
	return g.apiManager
}
