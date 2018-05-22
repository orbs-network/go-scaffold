package publicapi

import (
	"testing"
	"github.com/orbs-network/go-scaffold/types/services/publicapi"
)

func TestMock(t *testing.T) {
	m := &MockService{}

	m.When("Start", nil, nil).Return()
	m.Start(nil, nil)

	m.When("Stop").Return()
	m.Stop()

	m.When("IsStarted").Return()
	m.IsStarted()

	m.When("Transfer", nil).Return(&publicapi.TransferOutput{})
	m.Transfer(nil)

	m.When("GetBalance", nil).Return(&publicapi.GetBalanceOutput{})
	m.GetBalance(nil)
}