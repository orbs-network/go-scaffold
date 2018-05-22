package virtualmachine

import (
	"github.com/orbs-network/go-scaffold/services/statestorage"
	"github.com/orbs-network/go-scaffold/types/services/virtualmachine"
	"github.com/orbs-network/go-scaffold/utils/logger"
	"errors"
	"sync"
)

type Service interface {
	virtualmachine.Methods
	Start(stateStorage statestorage.Service, stop *chan error)
	Stop()
	IsStarted() bool
}

type service struct {
	logger logger.Interface
	stop *chan error
	transactionSync *sync.Mutex
	stateStorage statestorage.Service
}

func NewService(logger logger.Interface) Service {
	return &service{
		logger: logger,
		transactionSync: &sync.Mutex{},
	}
}

func (s *service) Start(stateStorage statestorage.Service, stop *chan error) {
	if stateStorage == nil || stop == nil {
		panic("required arguments not given")
	}
	if s.stop == nil {
		s.stop = stop
		s.stateStorage = stateStorage
		s.logger.Info("VirtualMachine service started")
	}
}

func (s *service) Stop() {
	if s.stop != nil {
		s.logger.Info("VirtualMachine service stopping")
		*s.stop <- errors.New("VirtualMachine service stopped")
		s.stop = nil
	}
}

func (s *service) IsStarted() bool {
	return s.stop != nil
}