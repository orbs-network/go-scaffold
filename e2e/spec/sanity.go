package spec

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/orbs-network/go-experiment/utils/testing"
	"github.com/orbs-network/go-experiment/services"
	"net/http"
)

var _ = Describe("Sanity", func() {

	var (
		node services.Node
		stop chan error
	)

	BeforeEach(func() {
		stop = make(chan error, 10)
		node = services.NewNode()
		node.Start(&stop)
	})

	AfterEach(func() {
		node.Stop()
	})

	It("should respond to GET /api/balance", func() {
		resp, err := http.Get("http://localhost:8080/api/balance")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(ResponseBodyAsString(resp)).To(Equal("balance ok"))
	})

	It("should respond to GET /api/transfer", func() {
		resp, err := http.Get("http://localhost:8080/api/transfer")
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(ResponseBodyAsString(resp)).To(Equal("transfer ok"))
	})

})