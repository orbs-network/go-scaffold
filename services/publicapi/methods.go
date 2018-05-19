package publicapi

import "github.com/orbs-network/go-experiment/types/services/publicapi"

func (s *service) Transfer(*publicapi.TransferInput) (*publicapi.TransferOutput, error) {
	return &publicapi.TransferOutput{Success: "ok", Result: 17}, nil
}

func (s *service) GetBalance(*publicapi.GetBalanceInput) (*publicapi.GetBalanceOutput, error) {
	return &publicapi.GetBalanceOutput{Success: "ok", Result: 19}, nil
}