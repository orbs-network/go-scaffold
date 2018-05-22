package virtualmachine

import (
	. "github.com/onsi/gomega"
	"testing"
	"github.com/orbs-network/go-scaffold/utils/logger"
	_statestorage "github.com/orbs-network/go-scaffold/services/statestorage"
	"github.com/orbs-network/go-scaffold/types/services/statestorage"
	"errors"
)

var transferTable = []struct{
	to string
	read1Err error
	read2Err error
	write1Err error
	write2Err error
	output int32
	errs bool
}{
	{"user2", errors.New("a"), nil, nil, nil, 0, true},
	{"user2", nil, errors.New("a"), nil, nil, 0, true},
	{"user2", nil, nil, errors.New("a"), nil, 0, true},
	{"user2", nil, nil, nil, errors.New("a"), 0, true},
	{"user2", nil, nil, nil, nil, 90, false},
	{"user1", nil, nil, nil, nil, 100,false},
}

func TestTransfer(t *testing.T) {
	Ω := NewGomegaWithT(t)
	for _, tt := range transferTable {
		stop := make(chan error, 10)
		s := NewService(&logger.StubLogger{})
		ss := &_statestorage.MockService{}
		s.Start(ss, &stop)

		ss.When("ReadKey", &statestorage.ReadKeyInput{Key: "user1"}).Return(&statestorage.ReadKeyOutput{Value: 100}, tt.read1Err)
		ss.When("ReadKey", &statestorage.ReadKeyInput{Key: "user2"}).Return(&statestorage.ReadKeyOutput{Value: 50}, tt.read2Err)
		ss.When("WriteKey", &statestorage.WriteKeyInput{Key: "user1", Value: 90}).Return(&statestorage.WriteKeyOutput{}, tt.write1Err)
		ss.When("WriteKey", &statestorage.WriteKeyInput{Key: "user2", Value: 60}).Return(&statestorage.WriteKeyOutput{}, tt.write2Err)
		output, err := s.(*service).processTransfer("user1", tt.to, 10)
		if tt.errs {
			Ω.Expect(err).To(HaveOccurred())
		} else {
			Ω.Expect(err).ToNot(HaveOccurred())
			Ω.Expect(output).To(BeEquivalentTo(tt.output))
		}
	}
}

var balanceTable = []struct{
	read1Err error
	output int32
	errs bool
}{
	{errors.New("a"), 0, true},
	{nil, 100, false},
}

func TestGetBalance(t *testing.T) {
	Ω := NewGomegaWithT(t)
	for _, tt := range balanceTable {
		stop := make(chan error, 10)
		s := NewService(&logger.StubLogger{})
		ss := &_statestorage.MockService{}
		s.Start(ss, &stop)

		ss.When("ReadKey", &statestorage.ReadKeyInput{Key: "user1"}).Return(&statestorage.ReadKeyOutput{Value: 100}, tt.read1Err)
		output, err := s.(*service).processGetBalance("user1")
		if tt.errs {
			Ω.Expect(err).To(HaveOccurred())
		} else {
			Ω.Expect(err).ToNot(HaveOccurred())
			Ω.Expect(output).To(BeEquivalentTo(tt.output))
		}
	}
}