package publicapi

import (
	"github.com/orbs-network/go-scaffold/services/virtualmachine"
	"github.com/orbs-network/go-scaffold/types/services/publicapi"
	"errors"
	"github.com/orbs-network/go-scaffold/utils/logger"
)

type Service interface {
	publicapi.Methods
	Start(virtualMachine virtualmachine.Service, stop *chan error)
	Stop()
	IsStarted() bool
}

type service struct {
	logger logger.Interface
	stop *chan error
	virtualMachine virtualmachine.Service
}

func NewService(logger logger.Interface) Service {
	return &service{
		logger: logger,
	}
}

func (s *service) Start(virtualMachine virtualmachine.Service, stop *chan error) {
	if virtualMachine == nil || stop == nil {
		panic("required arguments not given")
	}
	if s.stop == nil {
		s.stop = stop
		s.virtualMachine = virtualMachine
		s.logger.Info("PublicApi service started")
	}
}

func (s *service) Stop() {
	if s.stop != nil {
		s.logger.Info("PublicApi service stopping")
		*s.stop <- errors.New("PublicApi service stopped")
		s.stop = nil
	}
}

func (s *service) IsStarted() bool {
	return s.stop != nil
}