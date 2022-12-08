package promoPartner

import (
	"TopTrafficTest/service/promoPartner/internal/data/repository"
	"TopTrafficTest/service/promoPartner/internal/delivery/http"
	"TopTrafficTest/service/promoPartner/internal/tools/httpServer"
	httpServerConfig "TopTrafficTest/tools/config/httpServer"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type app struct {
	host 					string
	port					uint
	dataFileName			string
}

func NewApp(host string, port uint, dataFileName string) *app {
	return &app{host: host, port: port, dataFileName: dataFileName}
}


type server interface {
	StartHttpServer()
}

func (a *app) Start() {

	var (
		placementsRepository	http.PlacementsRepository
		httpServerHandlers 		httpServer.Handlers
		httpServerConf			httpServer.Config
		errorLogger				*log.Logger
		httpServ 				server
	)

	flags := log.LstdFlags | log.Lshortfile


	addr := a.host+":"+strconv.FormatUint(uint64(a.port), 10)

	errorLogger = log.New(os.Stdout, addr+" ERROR: ", flags)

	dp, _ := filepath.Abs("../../service/promoPartner/internal/data/repository/"+a.dataFileName)

	placementsRepository = repository.NewPlacements(dp)

	httpServerHandlers = http.NewPlacements(errorLogger, placementsRepository)

	httpServerConf = httpServerConfig.NewConfig(a.host, a.port)

	httpServ = httpServer.NewHttpServer(httpServerConf, httpServerHandlers)
	httpServ.StartHttpServer()
}
