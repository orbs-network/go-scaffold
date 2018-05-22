package statestorage

import (
	"github.com/orbs-network/go-scaffold/types/services/statestorage"
	"github.com/orbs-network/go-scaffold/utils/logger"
	"errors"
)

type Service interface {
	statestorage.Methods
	Start(stop *chan error)
	Stop()
	IsStarted() bool
}

type service struct {
	logger logger.Interface
	stop *chan error
	db map[string]int32
}

func NewService(logger logger.Interface) Service {
	return &service{
		logger: logger,
		db: make(map[string]int32),
	}
}

func (s *service) Start(stop *chan error) {
	if stop == nil {
		panic("stop channel not given")
	}
	if s.stop == nil {
		s.stop = stop
		s.logger.Info("StateStorage service started")
	}
}

func (s *service) Stop() {
	if s.stop != nil {
		s.logger.Info("StateStorage service stopping")
		*s.stop <- errors.New("StateStorage service stopped")
		s.stop = nil
	}
}

func (s *service) IsStarted() bool {
	return s.stop != nil
}