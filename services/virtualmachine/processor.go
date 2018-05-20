package virtualmachine

import "github.com/orbs-network/go-experiment/types/services/statestorage"

func (s *service) processTransfer(fromUsername string, toUsername string, amount int32) (int32, error) {
	fromBalance, err := s.stateStorage.ReadKey(&statestorage.ReadKeyInput{Key: fromUsername})
	if err != nil {
		return 0, err
	}
	toBalance, err := s.stateStorage.ReadKey(&statestorage.ReadKeyInput{Key: toUsername})
	if err != nil {
		return 0, err
	}
	_, err = s.stateStorage.WriteKey(&statestorage.WriteKeyInput{Key: fromUsername, Value: fromBalance.Value - amount})
	if err != nil {
		return 0, err
	}
	_, err = s.stateStorage.WriteKey(&statestorage.WriteKeyInput{Key: toUsername, Value: toBalance.Value + amount})
	if err != nil {
		return 0, err
	}
	return fromBalance.Value - amount, nil
}

func (s *service) processGetBalance(username string) (int32, error) {
	balance, err := s.stateStorage.ReadKey(&statestorage.ReadKeyInput{Key: username})
	if err != nil {
		return 0, err
	}
	return balance.Value, nil
}