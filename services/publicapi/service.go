package publicapi

import (
	"github.com/orbs-network/go-scaffold/services/virtualmachine"
	"github.com/orbs-network/go-scaffold/types/services/publicapi"
	"errors"
	"log"
)

type Service interface {
	publicapi.Methods
	Start(virtualMachine virtualmachine.Service, stop *chan error)
	Stop()
}

type service struct {
	stop *chan error
	virtualMachine virtualmachine.Service
}

func NewService() Service {
	return &service{}
}

func (s *service) Start(virtualMachine virtualmachine.Service, stop *chan error) {
	if s.stop == nil {
		s.stop = stop
		s.virtualMachine = virtualMachine
		log.Print("PublicApi service started")
	}
}

func (s *service) Stop() {
	if s.stop != nil {
		log.Print("PublicApi service stopping")
		*s.stop <- errors.New("PublicApi service stopped")
		s.stop = nil
	}
}