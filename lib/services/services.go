package services

import "cmd/src/main.go/lib/services/databases"

type AllServices struct {
	Databases databases.Databases
}

func StartAllServices() AllServices {
	var allServices AllServices

	allServices.Databases = databases.StartAllDatabase()

	return allServices
}
