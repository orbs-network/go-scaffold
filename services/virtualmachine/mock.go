package virtualmachine

import (
	"github.com/maraino/go-mock"
	"github.com/orbs-network/go-experiment/services/statestorage"
	"github.com/orbs-network/go-experiment/types/services/virtualmachine"
)

type MockService struct {
	mock.Mock
}

func (s *MockService) Start(stateStorage statestorage.Service, stop *chan error) {
	s.Called(stop)
}

func (s *MockService) Stop() {
	s.Called()
}

func (s *MockService) ProcessTransaction(input *virtualmachine.ProcessTransactionInput) (*virtualmachine.ProcessTransactionOutput, error) {
	ret := s.Called(input)
	return ret.Get(0).(*virtualmachine.ProcessTransactionOutput), ret.Error(1)
}

func (s *MockService) CallContract(input *virtualmachine.CallContractInput) (*virtualmachine.CallContractOutput, error) {
	ret := s.Called(input)
	return ret.Get(0).(*virtualmachine.CallContractOutput), ret.Error(1)
}