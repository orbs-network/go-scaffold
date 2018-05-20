package virtualmachine

import (
	"github.com/orbs-network/go-experiment/types/services/virtualmachine"
	"github.com/orbs-network/go-experiment/utils/errors"
)

func (s *service) ProcessTransaction(input *virtualmachine.ProcessTransactionInput) (*virtualmachine.ProcessTransactionOutput, error) {
	if input == nil {
		return nil, &errors.ErrInvalidInput{Argument: "ProcessTransactionInput", Method: "VirtualMachine.ProcessTransaction"}
	}
	if input.Arg == nil || input.Arg.From == nil || input.Arg.To == nil {
		return nil, &errors.ErrInvalidInput{Argument: "Arg", Method: "VirtualMachine.ProcessTransaction"}
	}
	if input.Method != "Transfer" {
		return nil, &errors.ErrInvalidInput{Argument: "Method", Method: "VirtualMachine.ProcessTransaction"}
	}
	result, err := s.processTransfer(input.Arg.From.Username, input.Arg.To.Username, input.Arg.Amount)
	return &virtualmachine.ProcessTransactionOutput{Result: result}, err
}

func (s *service) CallContract(input *virtualmachine.CallContractInput) (*virtualmachine.CallContractOutput, error) {
	if input == nil {
		return nil, &errors.ErrInvalidInput{Argument: "CallContractInput", Method: "VirtualMachine.CallContract"}
	}
	if input.Arg == nil {
		return nil, &errors.ErrInvalidInput{Argument: "Arg", Method: "VirtualMachine.CallContract"}
	}
	if input.Method != "GetBalance" {
		return nil, &errors.ErrInvalidInput{Argument: "Method", Method: "VirtualMachine.CallContract"}
	}
	result, err := s.processGetBalance(input.Arg.Username)
	return &virtualmachine.CallContractOutput{Result: result}, err
}
