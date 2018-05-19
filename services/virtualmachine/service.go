package virtualmachine

import (
	"github.com/orbs-network/go-experiment/services/statestorage"
	"github.com/orbs-network/go-experiment/types/services/virtualmachine"
	"errors"
	"log"
)

type Service interface {
	virtualmachine.Methods
	Start(stateStorage *statestorage.Service, stop *chan error)
	Stop()
}

type service struct {
	stop *chan error
	stateStorage *statestorage.Service
}

func NewService() Service {
	return &service{}
}

func (s *service) Start(stateStorage *statestorage.Service, stop *chan error) {
	if s.stop == nil {
		s.stop = stop
		s.stateStorage = stateStorage
		log.Print("VirtualMachine service started")
	}
}

func (s *service) Stop() {
	if s.stop != nil {
		*s.stop <- errors.New("VirtualMachine service stopped")
		s.stop = nil
	}
}