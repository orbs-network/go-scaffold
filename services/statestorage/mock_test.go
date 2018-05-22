package statestorage

import (
	"testing"
	"github.com/orbs-network/go-scaffold/types/services/statestorage"
)

func TestMock(t *testing.T) {
	m := &MockService{}

	m.When("Start",nil).Return()
	m.Start(nil)

	m.When("Stop").Return()
	m.Stop()

	m.When("IsStarted").Return()
	m.IsStarted()

	m.When("WriteKey", nil).Return(&statestorage.WriteKeyOutput{})
	m.WriteKey(nil)

	m.When("ReadKey", nil).Return(&statestorage.ReadKeyOutput{})
	m.ReadKey(nil)
}