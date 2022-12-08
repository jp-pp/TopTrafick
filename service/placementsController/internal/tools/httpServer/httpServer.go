package httpServer

import (
	"net/http"
	"strconv"
)


type httpServer struct {
	Config
	Handlers
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

func (s *httpServer) StartHttpServer() (err	error) {

	var (
		addr string
	)

	addr = s.Config.Host()+":"+strconv.FormatUint(uint64(s.Port()), 10)

	http.HandleFunc("/placements/request", s.Placements)
	err = http.ListenAndServe(addr,  nil)
	if err != nil {
		return
	}

	return

}


