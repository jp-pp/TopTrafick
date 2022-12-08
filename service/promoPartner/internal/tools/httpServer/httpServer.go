package httpServer

import (
	"log"
	"net/http"
	"os"
	"strconv"
)


type httpServer struct {
	Config
	Handlers
	infoLogger	*log.Logger
}

func NewHttpServer(config Config, handlers Handlers) *httpServer {
	return &httpServer{Config: config, Handlers: handlers}
}


type Config interface {
	Host() string
	Port() uint
}

type Handlers interface {
	Placements(w http.ResponseWriter, r *http.Request)
}

func (s *httpServer) StartHttpServer() {

	var (
		addr 	string
		err		error
	)

	addr = s.Config.Host()+":"+strconv.FormatUint(uint64(s.Port()), 10)

	s.infoLogger = log.New(os.Stdout, addr+" INFO: ", log.LstdFlags | log.Lshortfile)

	http.HandleFunc("/bid_request", s.Placements)

	s.infoLogger.Println("server start")

	err = http.ListenAndServe(addr,  nil)
	if err != nil {
		s.infoLogger.Panic(err)
	}
}


