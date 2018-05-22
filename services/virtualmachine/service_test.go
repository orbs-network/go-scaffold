package virtualmachine

import (
	. "github.com/onsi/gomega"
	"github.com/orbs-network/go-scaffold/utils/logger"
	"testing"
	"github.com/orbs-network/go-scaffold/services/statestorage"
)

func TestNewServiceRestart(t *testing.T) {
	Ω := NewGomegaWithT(t)
	stop := make(chan error, 10)
	s := NewService(&logger.StubLogger{})
	ss := statestorage.NewService(&logger.StubLogger{})
	Ω.Expect(s.IsStarted()).To(BeFalse())

	s.Start(ss, &stop)
	Ω.Expect(s.IsStarted()).To(BeTrue())
	Ω.Expect(stop).ToNot(Receive())

	s.Stop()
	Ω.Expect(s.IsStarted()).To(BeFalse())
	Ω.Expect(stop).To(Receive())

	s.Start(ss, &stop)
	Ω.Expect(s.IsStarted()).To(BeTrue())
	Ω.Expect(stop).ToNot(Receive())
}

func TestDoubleStart(t *testing.T) {
	Ω := NewGomegaWithT(t)
	stop := make(chan error, 10)
	stop2 := make(chan error, 10)
	s := NewService(&logger.StubLogger{})
	ss := statestorage.NewService(&logger.StubLogger{})

	s.Start(ss, &stop)
	s.Start(ss, &stop2)
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
	ss := statestorage.NewService(&logger.StubLogger{})

	s.Start(ss, &stop)
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
	ss := statestorage.NewService(&logger.StubLogger{})

	Ω.Expect(func () {
		s.Start(ss, nil)
	}).To(Panic())
}