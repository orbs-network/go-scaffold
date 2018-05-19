package statestorage

import (
	"github.com/orbs-network/go-experiment/types/services/statestorage"
)

func (s *service) WriteKey(*statestorage.WriteKeyInput) (*statestorage.WriteKeyOutput, error) {
	return &statestorage.WriteKeyOutput{}, nil
}

func (s *service) ReadKey(*statestorage.ReadKeyInput) (*statestorage.ReadKeyOutput, error) {
	return &statestorage.ReadKeyOutput{Value: 47}, nil
}