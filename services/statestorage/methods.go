package statestorage

import (
	"github.com/orbs-network/go-experiment/types/services/statestorage"
	"github.com/orbs-network/go-experiment/utils/errors"
)

func (s *service) WriteKey(input *statestorage.WriteKeyInput) (*statestorage.WriteKeyOutput, error) {
	if input == nil {
		return nil, &errors.ErrInvalidInput{Argument: "WriteKeyInput", Method: "StateStorage.WriteKey"}
	}
	if (input.Key == "") {
		return nil, &errors.ErrInvalidInput{Argument: "Key", Method: "StateStorage.WriteKey"}
	}
	s.db[input.Key] = input.Value
	return &statestorage.WriteKeyOutput{}, nil
}

func (s *service) ReadKey(input *statestorage.ReadKeyInput) (*statestorage.ReadKeyOutput, error) {
	if input == nil {
		return nil, &errors.ErrInvalidInput{Argument: "ReadKeyInput", Method: "StateStorage.ReadKey"}
	}
	if (input.Key == "") {
		return nil, &errors.ErrInvalidInput{Argument: "Key", Method: "StateStorage.ReadKey"}
	}
	value, ok := s.db[input.Key]
	if (!ok) {
		value = 0
	}
	return &statestorage.ReadKeyOutput{Value: value}, nil
}