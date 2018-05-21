package virtualmachine

import (
	"github.com/orbs-network/go-experiment/services/statestorage"
	"github.com/orbs-network/go-experiment/types/services/virtualmachine"
	"errors"
	"log"
	"sync"
)

type Service interface {
	virtualmachine.Methods
	Start(stateStorage statestorage.Service, stop *chan error)
	Stop()
}

type service struct {
	stop *chan error
	transactionSync *sync.Mutex
	stateStorage statestorage.Service
}

func NewService() Service {
	return &service{
		transactionSync: &sync.Mutex{},
	}
}

func (s *service) Start(stateStorage statestorage.Service, stop *chan error) {
	if s.stop == nil {
		s.stop = stop
		s.stateStorage = stateStorage
		log.Print("VirtualMachine service started")
	}
}

func (s *service) Stop() {
	if s.stop != nil {
		log.Print("VirtualMachine service stopping")
		*s.stop <- errors.New("VirtualMachine service stopped")
		s.stop = nil
	}
}