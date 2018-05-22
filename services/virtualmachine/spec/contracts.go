package spec

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/orbs-network/go-scaffold/utils/testing"
	uut "github.com/orbs-network/go-scaffold/services/virtualmachine"
	_statestorage "github.com/orbs-network/go-scaffold/services/statestorage"
	"github.com/orbs-network/go-scaffold/types/services/virtualmachine"
	"github.com/orbs-network/go-scaffold/types/protocol"
	"github.com/orbs-network/go-scaffold/types/services/statestorage"
)

var _ = Describe("Contracts", func() {

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
		_, err := service.CallContract(&virtualmachine.CallContractInput{Method: "unknown111"})
		Expect(err).To(HaveOccurred())
	})

	It("should support 'GetBalance' contract method", func() {
		stateStorage.When("ReadKey", &statestorage.ReadKeyInput{Key: "user1"}).Return(&statestorage.ReadKeyOutput{Value: 100}, nil).Times(1)
		addr := protocol.Address{Username: "user1"}
		out, err := service.CallContract(&virtualmachine.CallContractInput{Method: "GetBalance", Arg: &addr})
		Expect(err).ToNot(HaveOccurred())
		Expect(out.Result).To(BeEquivalentTo(100))
		Expect(stateStorage).To(ExecuteAsPlanned())
	})

})