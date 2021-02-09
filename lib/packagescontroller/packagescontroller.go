package packagescontroller

import "cmd/src/main.go/lib/services"

type PackagesController struct {
	Services services.AllServices
}

func StartPackagesController(services services.AllServices) {
	var packagesController PackagesController
	packagesController.Services = services
	packagesController.startRepositories()
	packagesController.StartDependecies()
}

func (pc PackagesController) startRepositories() {

}

func (pc PackagesController) StartDependecies() {

}
