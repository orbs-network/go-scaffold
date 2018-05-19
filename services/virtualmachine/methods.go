package virtualmachine

import "github.com/orbs-network/go-experiment/types/services/virtualmachine"

func (s *service) ProcessTransaction(*virtualmachine.ProcessTransactionInput) (*virtualmachine.ProcessTransactionOutput, error) {
	return &virtualmachine.ProcessTransactionOutput{Result: 34}, nil
}

func (s *service) CallContract(*virtualmachine.CallContractInput) (*virtualmachine.CallContractOutput, error) {
	return &virtualmachine.CallContractOutput{Result: 88}, nil
}