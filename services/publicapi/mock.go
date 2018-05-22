package publicapi

import (
	"github.com/maraino/go-mock"
	"github.com/orbs-network/go-scaffold/services/virtualmachine"
	"github.com/orbs-network/go-scaffold/types/services/publicapi"
)

type MockService struct {
	mock.Mock
}

func (s *MockService) Start(virtualMachine virtualmachine.Service, stop *chan error) {
	s.Called(virtualMachine, stop)
}

func (s *MockService) Stop() {
	s.Called()
}

func (s *MockService) IsStarted() bool {
	return s.Called().Bool(0)
}

func (s *MockService) Transfer(input *publicapi.TransferInput) (*publicapi.TransferOutput, error) {
	ret := s.Called(input)
	return ret.Get(0).(*publicapi.TransferOutput), ret.Error(1)
}

func (s *MockService) GetBalance(input *publicapi.GetBalanceInput) (*publicapi.GetBalanceOutput, error) {
	ret := s.Called(input)
	return ret.Get(0).(*publicapi.GetBalanceOutput), ret.Error(1)
}