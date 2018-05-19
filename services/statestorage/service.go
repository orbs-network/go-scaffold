package statestorage

import (
	"github.com/orbs-network/go-experiment/types/services/statestorage"
	"errors"
	"log"
)

type Service interface {
	statestorage.Methods
	Start(stop *chan error)
	Stop()
}

type service struct {
	stop *chan error
	db map[string]int32
}

func NewService() Service {
	return &service{
		db: make(map[string]int32),
	}
}

func (s *service) Start(stop *chan error) {
	if s.stop == nil {
		s.stop = stop
		log.Print("StateStorage service started")
	}
}

func (s *service) Stop() {
	if s.stop != nil {
		log.Print("StateStorage service stopping")
		*s.stop <- errors.New("StateStorage service stopped")
		s.stop = nil
	}
}