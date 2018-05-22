package publicapi

import (
	. "github.com/onsi/gomega"
	"testing"
	"github.com/orbs-network/go-scaffold/utils/logger"
	"github.com/orbs-network/go-scaffold/types/services/publicapi"
	"github.com/orbs-network/go-scaffold/types/protocol"
	_virtualmachine "github.com/orbs-network/go-scaffold/services/virtualmachine"
	"github.com/orbs-network/go-scaffold/types/services/virtualmachine"
	"errors"
)

var transferTable = []struct{
	input *publicapi.TransferInput
	processErr error
	output *publicapi.TransferOutput
	errs bool
}{
	{nil,nil,nil, true},
	{&publicapi.TransferInput{Transaction: nil},nil,nil,true},
	{&publicapi.TransferInput{Transaction: &protocol.Transaction{From: nil, To: &protocol.Address{Username: "user2"}}},nil, nil,true},
	{&publicapi.TransferInput{Transaction: &protocol.Transaction{From: &protocol.Address{Username: "user1"}, To: nil}},nil, nil,true},
	{&publicapi.TransferInput{Transaction: &protocol.Transaction{From: &protocol.Address{Username: "user1"}, To: &protocol.Address{Username: "user2"}}},errors.New("a"), nil,true},
	{&publicapi.TransferInput{Transaction: &protocol.Transaction{From: &protocol.Address{Username: "user1"}, To: &protocol.Address{Username: "user2"}}},nil, &publicapi.TransferOutput{Success: "ok", Result: 10},false},
}

func TestTransfer(t *testing.T) {
	Ω := NewGomegaWithT(t)
	for _, tt := range transferTable {
		stop := make(chan error, 10)
		s := NewService(&logger.StubLogger{})
		vm := &_virtualmachine.MockService{}
		s.Start(vm, &stop)

		vm.When("ProcessTransaction", &virtualmachine.ProcessTransactionInput{Method: "Transfer", Arg: &protocol.Transaction{From: &protocol.Address{Username: "user1"}, To: &protocol.Address{Username: "user2"}}}).Return(&virtualmachine.ProcessTransactionOutput{Result: 10}, tt.processErr)
		output, err := s.Transfer(tt.input)
		if tt.errs {
			Ω.Expect(err).To(HaveOccurred())
		} else {
			Ω.Expect(err).ToNot(HaveOccurred())
			Ω.Expect(output).To(Equal(tt.output))
		}
	}
}

var balanceTable = []struct{
	input *publicapi.GetBalanceInput
	callErr error
	output *publicapi.GetBalanceOutput
	errs bool
}{
	{nil,nil,nil, true},
	{&publicapi.GetBalanceInput{From: nil},nil,nil,true},
	{&publicapi.GetBalanceInput{From: &protocol.Address{Username: "user1"}},errors.New("a"), nil,true},
	{&publicapi.GetBalanceInput{From: &protocol.Address{Username: "user1"}},nil, &publicapi.GetBalanceOutput{Success: "ok", Result: 10},false},
}

func TestGetBalance(t *testing.T) {
	Ω := NewGomegaWithT(t)
	for _, tt := range balanceTable {
		stop := make(chan error, 10)
		s := NewService(&logger.StubLogger{})
		vm := &_virtualmachine.MockService{}
		s.Start(vm, &stop)

		vm.When("CallContract", &virtualmachine.CallContractInput{Method: "GetBalance", Arg: &protocol.Address{Username: "user1"}}).Return(&virtualmachine.CallContractOutput{Result: 10}, tt.callErr)
		output, err := s.GetBalance(tt.input)
		if tt.errs {
			Ω.Expect(err).To(HaveOccurred())
		} else {
			Ω.Expect(err).ToNot(HaveOccurred())
			Ω.Expect(output).To(Equal(tt.output))
		}
	}
}