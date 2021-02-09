package services

import (
	"cmd/src/main.go/lib/services/apimanager"
	"cmd/src/main.go/lib/services/databases"
)

type AllServices struct {
	Databases  databases.Databases
	ApiManager apimanager.ApiManager
}

func StartAllServices() AllServices {
	var allServices AllServices

	allServices.Databases = databases.StartAllDatabase()
	allServices.ApiManager = apimanager.StartApiManager()

	return allServices
}
