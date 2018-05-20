package statestorage

import (
	"github.com/maraino/go-mock"
	"github.com/orbs-network/go-experiment/types/services/statestorage"
)

type MockService struct {
	mock.Mock
}

func (s *MockService) Start(stop *chan error) {
	s.Called(stop)
}

func (s *MockService) Stop() {
	s.Called()
}

func (s *MockService) WriteKey(input *statestorage.WriteKeyInput) (*statestorage.WriteKeyOutput, error) {
	ret := s.Called(input)
	return ret.Get(0).(*statestorage.WriteKeyOutput), ret.Error(1)
}

func (s *MockService) ReadKey(input *statestorage.ReadKeyInput) (*statestorage.ReadKeyOutput, error) {
	ret := s.Called(input)
	return ret.Get(0).(*statestorage.ReadKeyOutput), ret.Error(1)
}