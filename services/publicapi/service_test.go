package publicapi

import (
	. "github.com/onsi/gomega"
	"github.com/orbs-network/go-scaffold/utils/logger"
	"github.com/orbs-network/go-scaffold/services/virtualmachine"
	"testing"
)

func TestNewServiceRestart(t *testing.T) {
	Ω := NewGomegaWithT(t)
	stop := make(chan error, 10)
	s := NewService(&logger.StubLogger{})
	vm := virtualmachine.NewService(&logger.StubLogger{})
	Ω.Expect(s.IsStarted()).To(BeFalse())

	s.Start(vm, &stop)
	Ω.Expect(s.IsStarted()).To(BeTrue())
	Ω.Expect(stop).ToNot(Receive())

	s.Stop()
	Ω.Expect(s.IsStarted()).To(BeFalse())
	Ω.Expect(stop).To(Receive())

	s.Start(vm, &stop)
	Ω.Expect(s.IsStarted()).To(BeTrue())
	Ω.Expect(stop).ToNot(Receive())
}

func TestDoubleStart(t *testing.T) {
	Ω := NewGomegaWithT(t)
	stop := make(chan error, 10)
	stop2 := make(chan error, 10)
	s := NewService(&logger.StubLogger{})
	vm := virtualmachine.NewService(&logger.StubLogger{})

	s.Start(vm, &stop)
	s.Start(vm, &stop2)
	Ω.Expect(s.IsStarted()).To(BeTrue())
	Ω.Expect(stop).ToNot(Receive())
	Ω.Expect(stop2).ToNot(Receive())

	s.Stop()
	Ω.Expect(s.IsStarted()).To(BeFalse())
	Ω.Expect(stop).To(Receive())
	Ω.Expect(stop2).ToNot(Receive())
}

func TestDoubleStop(t *testing.T) {
	Ω := NewGomegaWithT(t)
	stop := make(chan error, 10)
	s := NewService(&logger.StubLogger{})
	vm := virtualmachine.NewService(&logger.StubLogger{})

	s.Start(vm, &stop)
	s.Stop()
	Ω.Expect(s.IsStarted()).To(BeFalse())
	Ω.Expect(stop).To(Receive())

	s.Stop()
	Ω.Expect(s.IsStarted()).To(BeFalse())
	Ω.Expect(stop).ToNot(Receive())
}

func TestNoChannel(t *testing.T) {
	Ω := NewGomegaWithT(t)
	s := NewService(&logger.StubLogger{})
	vm := virtualmachine.NewService(&logger.StubLogger{})

	Ω.Expect(func () {
		s.Start(vm, nil)
	}).To(Panic())
}