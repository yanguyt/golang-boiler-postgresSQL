package apimanager

import "cmd/src/main.go/lib/services/apimanager/ginhttp"

type ApiManager struct {
	Gin ginhttp.GinFamily
}

func StartApiManager() ApiManager {
	var apiManager ApiManager

	apiManager.Gin = ginhttp.StartGinEngine()

	return apiManager
}
