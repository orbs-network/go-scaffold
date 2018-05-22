package virtualmachine

import (
	. "github.com/onsi/gomega"
	"testing"
	"github.com/orbs-network/go-scaffold/utils/logger"
	"github.com/orbs-network/go-scaffold/types/services/virtualmachine"
	"github.com/orbs-network/go-scaffold/types/protocol"
	"github.com/orbs-network/go-scaffold/services/statestorage"
)

var transactionTable = []struct{
	input *virtualmachine.ProcessTransactionInput
	output *virtualmachine.ProcessTransactionOutput
	errs bool
}{
	{nil, nil, true},
	{&virtualmachine.ProcessTransactionInput{Method: "Transfer", Arg: nil},nil, true},
	{&virtualmachine.ProcessTransactionInput{Method: "Transfer", Arg: &protocol.Transaction{From: nil, To: &protocol.Address{Username: "user"}}},nil, true},
	{&virtualmachine.ProcessTransactionInput{Method: "Transfer", Arg: &protocol.Transaction{From: &protocol.Address{Username: "user"}, To: nil}},nil, true},
	{&virtualmachine.ProcessTransactionInput{Method: "Other", Arg: &protocol.Transaction{From: &protocol.Address{Username: "user1"}, To: &protocol.Address{Username: "user2"}}}, nil, true},
	{&virtualmachine.ProcessTransactionInput{Method: "Transfer", Arg: &protocol.Transaction{From: &protocol.Address{Username: "user1"}, To: &protocol.Address{Username: "user2"}, Amount: 10}}, &virtualmachine.ProcessTransactionOutput{Result: -10}, false},
}

func TestProcessTransaction(t *testing.T) {
	Ω := NewGomegaWithT(t)
	for _, tt := range transactionTable {
		stop := make(chan error, 10)
		s := NewService(&logger.StubLogger{})
		ss := statestorage.NewService(&logger.StubLogger{})
		s.Start(ss, &stop)

		output, err := s.ProcessTransaction(tt.input)
		if tt.errs {
			Ω.Expect(err).To(HaveOccurred())
		} else {
			Ω.Expect(err).ToNot(HaveOccurred())
			Ω.Expect(output).To(Equal(tt.output))
		}
	}
}

var contractTable = []struct{
	input *virtualmachine.CallContractInput
	output *virtualmachine.CallContractOutput
	errs bool
}{
	{nil, nil, true},
	{&virtualmachine.CallContractInput{Method: "GetBalance", Arg: nil},nil, true},
	{&virtualmachine.CallContractInput{Method: "Other", Arg: &protocol.Address{Username: "user1"}}, nil, true},
	{&virtualmachine.CallContractInput{Method: "GetBalance", Arg: &protocol.Address{Username: "user1"}}, &virtualmachine.CallContractOutput{Result: 0}, false},
}

func TestCallContract(t *testing.T) {
	Ω := NewGomegaWithT(t)
	for _, tt := range contractTable {
		stop := make(chan error, 10)
		s := NewService(&logger.StubLogger{})
		ss := statestorage.NewService(&logger.StubLogger{})
		s.Start(ss, &stop)

		output, err := s.CallContract(tt.input)
		if tt.errs {
			Ω.Expect(err).To(HaveOccurred())
		} else {
			Ω.Expect(err).ToNot(HaveOccurred())
			Ω.Expect(output).To(Equal(tt.output))
		}
	}
}
