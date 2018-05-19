package publicapi

import (
	"net/http"
	"log"
)

type Server interface {
	Start(service *Service, stop *chan error)
	Stop()
}

type server struct {
	stop *chan error
	service *Service
	httpServer *http.Server
}

func NewServer() Server {
	return &server{}
}

func (s *server) Start(service *Service, stop *chan error) {
	if s.stop == nil {
		s.stop = stop
		s.service = service
		s.httpServer = &http.Server{Addr: ":8080", Handler: s.createRouter()}
		go s.startHttp()
		log.Print("PublicApi server started")
	}
}

func (s *server) Stop() {
	if s.stop != nil && s.httpServer != nil {
		*s.stop <- s.httpServer.Shutdown(nil)
		s.stop = nil
	}
}

func (s *server) createRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/api/transfer", s.transferHandler)
	router.HandleFunc("/api/balance", s.balanceHandler)
	return router
}

func (s *server) startHttp() {
	*s.stop <- s.httpServer.ListenAndServe()
}