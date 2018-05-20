package spec

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/orbs-network/go-experiment/utils/testing"
	uut "github.com/orbs-network/go-experiment/services/virtualmachine"
	_statestorage "github.com/orbs-network/go-experiment/services/statestorage"
	"github.com/orbs-network/go-experiment/types/services/virtualmachine"
	"github.com/orbs-network/go-experiment/types/protocol"
	"github.com/orbs-network/go-experiment/types/services/statestorage"
)

var _ = Describe("Transactions", func() {

	var (
		service uut.Service
		stateStorage *_statestorage.MockService
		stop chan error
	)

	BeforeEach(func() {
		stop = make(chan error, 10)
		service = uut.NewService()
		stateStorage = &_statestorage.MockService{}
		service.Start(stateStorage, &stop)
	})

	AfterEach(func() {
		service.Stop()
	})

	It("should return error on unknown method", func() {
		_, err := service.ProcessTransaction(&virtualmachine.ProcessTransactionInput{Method: "unknown111"})
		Expect(err).To(HaveOccurred())
	})

	It("should support 'Transfer' transaction method", func() {
		stateStorage.When("ReadKey", &statestorage.ReadKeyInput{Key: "user1"}).Return(&statestorage.ReadKeyOutput{Value: 100}, nil).Times(1)
		stateStorage.When("ReadKey", &statestorage.ReadKeyInput{Key: "user2"}).Return(&statestorage.ReadKeyOutput{Value: 50}, nil).Times(1)
		stateStorage.When("WriteKey", &statestorage.WriteKeyInput{Key: "user1", Value: 90}).Return(&statestorage.WriteKeyOutput{}, nil).Times(1)
		stateStorage.When("WriteKey", &statestorage.WriteKeyInput{Key: "user2", Value: 60}).Return(&statestorage.WriteKeyOutput{}, nil).Times(1)
		t := protocol.Transaction{From: &protocol.Address{Username: "user1"}, To: &protocol.Address{Username: "user2"}, Amount: 10}
		out, err := service.ProcessTransaction(&virtualmachine.ProcessTransactionInput{Method: "Transfer", Arg: &t})
		Expect(err).ToNot(HaveOccurred())
		Expect(out.Result).To(BeEquivalentTo(90))
		Expect(stateStorage).To(ExecuteAsPlanned())
	})

})