package publicapi

import (
	"net/http"
	"context"
	"time"
	"github.com/orbs-network/go-scaffold/utils/logger"
)

type Server interface {
	Start(service Service, stop *chan error)
	Stop()
	IsStarted() bool
}

type server struct {
	logger logger.Interface
	stop *chan error
	service Service
	httpServer *http.Server
}

func NewServer(logger logger.Interface) Server {
	return &server{
		logger: logger,
	}
}

func (s *server) Start(service Service, stop *chan error) {
	if service == nil || stop == nil {
		panic("required arguments not given")
	}
	if s.stop == nil {
		s.stop = stop
		s.service = service
		s.httpServer = &http.Server{Addr: ":8080", Handler: s.createRouter()}
		go s.startHttp()
		s.logger.Info("PublicApi server started")
	}
}

func (s *server) Stop() {
	if s.stop != nil && s.httpServer != nil {
		s.logger.Info("PublicApi server stopping")
		ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
		*s.stop <- s.httpServer.Shutdown(ctx)
		s.stop = nil
	}
}

func (s *server) IsStarted() bool {
	return s.stop != nil
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