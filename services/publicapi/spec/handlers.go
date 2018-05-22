package spec

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/orbs-network/go-scaffold/utils/testing"
	uut "github.com/orbs-network/go-scaffold/services/publicapi"
	_virtualmachine "github.com/orbs-network/go-scaffold/services/virtualmachine"
	"github.com/orbs-network/go-scaffold/types/services/virtualmachine"
	"github.com/orbs-network/go-scaffold/types/protocol"
	"github.com/orbs-network/go-scaffold/utils/logger"
	"net/http"
)

var _ = Describe("HTTP", func() {

	var (
		service uut.Service
		server uut.Server
		virtualMachine *_virtualmachine.MockService
		stop chan error
	)

	BeforeEach(func() {
		stop = make(chan error, 10)
		service = uut.NewService(&logger.StubLogger{})
		server = uut.NewServer(&logger.StubLogger{})
		virtualMachine = &_virtualmachine.MockService{}
		service.Start(virtualMachine, &stop)
		server.Start(service, &stop)
	})

	AfterEach(func() {
		server.Stop()
		service.Stop()
	})

	It("should respond to GET /api/transfer", func() {
		t := protocol.Transaction{From: &protocol.Address{Username: "user1"}, To: &protocol.Address{Username: "user2"}, Amount: 10}
		virtualMachine.When("ProcessTransaction", &virtualmachine.ProcessTransactionInput{Method: "Transfer", Arg: &t}).Return(&virtualmachine.ProcessTransactionOutput{Result: 100}, nil).Times(1)
		resp, err := http.Get("http://localhost:8080/api/transfer?from=user1&to=user2&amount=10")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(ResponseBodyAsString(resp)).To(Equal("100"))
		Expect(virtualMachine).To(ExecuteAsPlanned())
	})

	It("should respond to GET /api/balance", func() {
		addr := protocol.Address{Username: "user1"}
		virtualMachine.When("CallContract", &virtualmachine.CallContractInput{Method: "GetBalance", Arg: &addr}).Return(&virtualmachine.CallContractOutput{Result: 100}, nil).Times(1)
		resp, err := http.Get("http://localhost:8080/api/balance?from=user1")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(ResponseBodyAsString(resp)).To(Equal("100"))
		Expect(virtualMachine).To(ExecuteAsPlanned())
	})

})