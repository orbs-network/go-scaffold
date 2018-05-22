package publicapi

import (
	"github.com/orbs-network/go-scaffold/types/services/publicapi"
	"github.com/orbs-network/go-scaffold/utils/errors"
	"github.com/orbs-network/go-scaffold/types/services/virtualmachine"
)

func (s *service) Transfer(input *publicapi.TransferInput) (*publicapi.TransferOutput, error) {
	if input == nil {
		return nil, &errors.ErrInvalidInput{Argument: "TransferInput", Method: "PublicApi.Transfer"}
	}
	if input.Transaction == nil || input.Transaction.From == nil || input.Transaction.To == nil {
		return nil, &errors.ErrInvalidInput{Argument: "Transaction", Method: "PublicApi.Transfer"}
	}
	result, err := s.virtualMachine.ProcessTransaction(&virtualmachine.ProcessTransactionInput{Method: "Transfer", Arg: input.Transaction})
	if err != nil {
		return nil, err
	}
	return &publicapi.TransferOutput{Success: "ok", Result: result.Result}, nil
}

func (s *service) GetBalance(input *publicapi.GetBalanceInput) (*publicapi.GetBalanceOutput, error) {
	if input == nil {
		return nil, &errors.ErrInvalidInput{Argument: "GetBalanceInput", Method: "PublicApi.GetBalance"}
	}
	if input.From == nil {
		return nil, &errors.ErrInvalidInput{Argument: "From", Method: "PublicApi.GetBalance"}
	}
	result, err := s.virtualMachine.CallContract(&virtualmachine.CallContractInput{Method: "GetBalance", Arg: input.From})
	if err != nil {
		return nil, err
	}
	return &publicapi.GetBalanceOutput{Success: "ok", Result: result.Result}, nil
}