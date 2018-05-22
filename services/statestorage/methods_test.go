package statestorage

import (
	. "github.com/onsi/gomega"
	"testing"
	"github.com/orbs-network/go-scaffold/types/services/statestorage"
	"github.com/orbs-network/go-scaffold/utils/logger"
)

var writeTable = []struct{
	input *statestorage.WriteKeyInput
	output *statestorage.WriteKeyOutput
	errs bool
}{
	{nil, nil, true},
	{&statestorage.WriteKeyInput{Key: ""},nil, true},
	{&statestorage.WriteKeyInput{Key: "user1", Value: 17}, &statestorage.WriteKeyOutput{}, false},
}

func TestWriteKey(t *testing.T) {
	Ω := NewGomegaWithT(t)
	for _, tt := range writeTable {
		stop := make(chan error, 10)
		s := NewService(&logger.StubLogger{})
		s.Start(&stop)

		output, err := s.WriteKey(tt.input)
		if tt.errs {
			Ω.Expect(err).To(HaveOccurred())
		} else {
			Ω.Expect(err).ToNot(HaveOccurred())
			Ω.Expect(output).To(Equal(tt.output))
		}
	}
}

var readTable = []struct{
	input *statestorage.ReadKeyInput
	output *statestorage.ReadKeyOutput
	errs bool
}{
	{nil, nil, true},
	{&statestorage.ReadKeyInput{Key: ""},nil, true},
	{&statestorage.ReadKeyInput{Key: "user1"}, &statestorage.ReadKeyOutput{Value: 0}, false},
}

func TestReadKey(t *testing.T) {
	Ω := NewGomegaWithT(t)
	for _, tt := range readTable {
		stop := make(chan error, 10)
		s := NewService(&logger.StubLogger{})
		s.Start(&stop)

		output, err := s.ReadKey(tt.input)
		if tt.errs {
			Ω.Expect(err).To(HaveOccurred())
		} else {
			Ω.Expect(err).ToNot(HaveOccurred())
			Ω.Expect(output).To(Equal(tt.output))
		}
	}
}