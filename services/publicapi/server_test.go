package publicapi

import (
	. "github.com/onsi/gomega"
	"github.com/orbs-network/go-scaffold/utils/logger"
	"testing"
)

func TestNewServerRestart(t *testing.T) {
	Ω := NewGomegaWithT(t)
	stop := make(chan error, 10)
	s := NewServer(&logger.StubLogger{})
	pa := NewService(&logger.StubLogger{})
	Ω.Expect(s.IsStarted()).To(BeFalse())

	s.Start(pa, &stop)
	Ω.Expect(s.IsStarted()).To(BeTrue())
	Ω.Expect(stop).ToNot(Receive())

	s.Stop()
	Ω.Expect(s.IsStarted()).To(BeFalse())
	Ω.Expect(stop).To(Receive())

	s.Start(pa, &stop)
	Ω.Expect(s.IsStarted()).To(BeTrue())
	Ω.Expect(stop).ToNot(Receive())
}

func TestServerDoubleStart(t *testing.T) {
	Ω := NewGomegaWithT(t)
	stop := make(chan error, 10)
	stop2 := make(chan error, 10)
	s := NewServer(&logger.StubLogger{})
	pa := NewService(&logger.StubLogger{})

	s.Start(pa, &stop)
	s.Start(pa, &stop2)
	Ω.Expect(s.IsStarted()).To(BeTrue())
	Ω.Expect(stop).ToNot(Receive())
	Ω.Expect(stop2).ToNot(Receive())

	s.Stop()
	Ω.Expect(s.IsStarted()).To(BeFalse())
	Ω.Expect(stop).To(Receive())
	Ω.Expect(stop2).ToNot(Receive())
}

func TestServerDoubleStop(t *testing.T) {
	Ω := NewGomegaWithT(t)
	stop := make(chan error, 10)
	s := NewServer(&logger.StubLogger{})
	pa := NewService(&logger.StubLogger{})

	s.Start(pa, &stop)
	s.Stop()
	Ω.Expect(s.IsStarted()).To(BeFalse())
	Ω.Expect(stop).To(Receive())

	s.Stop()
	Ω.Expect(s.IsStarted()).To(BeFalse())
	Ω.Expect(stop).ToNot(Receive())
}

func TestServerNoChannel(t *testing.T) {
	Ω := NewGomegaWithT(t)
	s := NewServer(&logger.StubLogger{})
	pa := NewService(&logger.StubLogger{})

	Ω.Expect(func () {
		s.Start(pa, nil)
	}).To(Panic())
}