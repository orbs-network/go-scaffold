package publicapi

import (
	"errors"
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
}

func NewServer() Server {
	return &server{}
}

func (s *server) Start(service *Service, stop *chan error) {
	if s.stop == nil {
		s.stop = stop
		s.service = service
		go s.startHttp()
		log.Print("PublicApi server started")
	}
}

func (s *server) Stop() {
	if s.stop != nil {
		*s.stop <- errors.New("PublicApi server stopped")
		s.stop = nil
	}
}

func (s *server) startHttp() {
	http.HandleFunc("/transfer/", s.transferHandler)
	http.HandleFunc("/balance/", s.balanceHandler)
	*s.stop <- http.ListenAndServe(":8000", nil)
}