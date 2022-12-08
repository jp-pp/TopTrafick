package placementsController

import (
	httpController "TopTrafficTest/service/placementsController/internal/delivery/http"
	placementsUseCase "TopTrafficTest/service/placementsController/internal/domain/useCase/placements"
	httpServerConfig "TopTrafficTest/service/placementsController/internal/tools/config/httpServer"
	httpServerTool "TopTrafficTest/service/placementsController/internal/tools/httpServer"
)

type app struct {
	httpServer
}

type httpServer interface {
	StartHttpServer() (err	error)
}

func (a *app) NewApp()  {

	var (
		httpServerConf  		httpServerTool.Config
		httpServerHandlers 		httpServerTool.Handlers
		httpControllerUseCase	httpController.PlacementsUseCase
		placementsRepositories	[]placementsUseCase.Repository
	)

	placementsRepositories = make([]placementsUseCase.Repository, 0)

	httpControllerUseCase = placementsUseCase.NewPlacements(placementsRepositories)

	httpServerConf = httpServerConfig.NewConfig("127.0.0.1", 6000)
	httpServerHandlers = httpController.NewPlacementsController(httpControllerUseCase)


	a.httpServer = httpServerTool.NewHttpServer(httpServerConf, httpServerHandlers)
}

func (a *app) Start()  {

	var (
		err error
	)

	err = a.httpServer.StartHttpServer()
	if err != nil {
		panic(err)
	}
}