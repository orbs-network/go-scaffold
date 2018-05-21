package publicapi

import (
	"net/http"
	"log"
	"context"
	"time"
)

type Server interface {
	Start(service Service, stop *chan error)
	Stop()
}

type server struct {
	stop *chan error
	service Service
	httpServer *http.Server
}

func NewServer() Server {
	return &server{}
}

func (s *server) Start(service Service, stop *chan error) {
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
		log.Print("PublicApi server stopping")
		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
		*s.stop <- s.httpServer.Shutdown(ctx)
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
	err := s.httpServer.ListenAndServe()
	if err != nil && s.stop != nil {
		*s.stop <- err
	}
}