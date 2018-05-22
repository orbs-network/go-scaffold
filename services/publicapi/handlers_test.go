package publicapi

import (
	. "github.com/onsi/gomega"
	"testing"
	"github.com/orbs-network/go-scaffold/utils/logger"
	"github.com/orbs-network/go-scaffold/types/protocol"
	"errors"
	"github.com/orbs-network/go-scaffold/types/services/publicapi"
	"net/http"
	"net/http/httptest"
)

var transferHandlerTable = []struct{
	url string
	transferErr error
	errs bool
}{
	{"/transfer?", nil, true},
	{"/transfer?from=", nil, true},
	{"/transfer?from=user1", nil, true},
	{"/transfer?from=user1&to=", nil, true},
	{"/transfer?from=user1&to=user2", nil, true},
	{"/transfer?from=user1&to=user2&amount=", nil, true},
	{"/transfer?from=user1&to=user2&amount=abc", nil, true},
	{"/transfer?from=user1&to=user2&amount=3.14", nil, true},
	{"/transfer?from=user1&to=user2&amount=1700", errors.New("a"), true},
	{"/transfer?from=user1&to=user2&amount=1700", nil, false},
}

func TestTransferHandler(t *testing.T) {
	Ω := NewGomegaWithT(t)
	for _, tt := range transferHandlerTable {
		stop := make(chan error, 10)
		s := NewServer(&logger.StubLogger{})
		pa := &MockService{}
		s.Start(pa, &stop)

		pa.When("Transfer", &publicapi.TransferInput{Transaction: &protocol.Transaction{From: &protocol.Address{Username: "user1"}, To: &protocol.Address{Username: "user2"}, Amount: 1700}}).Return(&publicapi.TransferOutput{Success: "ok", Result: 10}, tt.transferErr)
		req, _ := http.NewRequest("GET", tt.url, nil)
		rec := httptest.NewRecorder()
		s.(*server).transferHandler(rec, req)
		if tt.errs {
			Ω.Expect(rec.Body.String()).To(HavePrefix("error"))
		} else {
			Ω.Expect(rec.Body.String()).To(Equal("10"))
		}
	}
}

var balanceHandlerTable = []struct{
	url string
	balanceErr error
	errs bool
}{
	{"/balance?", nil, true},
	{"/transfer?from=", nil, true},
	{"/transfer?from=user1", errors.New("a"), true},
	{"/transfer?from=user1", nil, false},
}

func TestBalanceHandler(t *testing.T) {
	Ω := NewGomegaWithT(t)
	for _, tt := range balanceHandlerTable {
		stop := make(chan error, 10)
		s := NewServer(&logger.StubLogger{})
		pa := &MockService{}
		s.Start(pa, &stop)

		pa.When("GetBalance", &publicapi.GetBalanceInput{From: &protocol.Address{Username: "user1"}}).Return(&publicapi.GetBalanceOutput{Success: "ok", Result: 10}, tt.balanceErr)
		req, _ := http.NewRequest("GET", tt.url, nil)
		rec := httptest.NewRecorder()
		s.(*server).balanceHandler(rec, req)
		if tt.errs {
			Ω.Expect(rec.Body.String()).To(HavePrefix("error"))
		} else {
			Ω.Expect(rec.Body.String()).To(Equal("10"))
		}
	}
}