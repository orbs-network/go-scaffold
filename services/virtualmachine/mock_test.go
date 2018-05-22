package virtualmachine

import (
	"testing"
	"github.com/orbs-network/go-scaffold/types/services/virtualmachine"
)

func TestMock(t *testing.T) {
	m := &MockService{}

	m.When("Start", nil, nil).Return()
	m.Start(nil, nil)

	m.When("Stop").Return()
	m.Stop()

	m.When("IsStarted").Return()
	m.IsStarted()

	m.When("ProcessTransaction", nil).Return(&virtualmachine.ProcessTransactionOutput{})
	m.ProcessTransaction(nil)

	m.When("CallContract", nil).Return(&virtualmachine.CallContractOutput{})
	m.CallContract(nil)
}