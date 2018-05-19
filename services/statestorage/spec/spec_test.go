package spec

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestStateStorageSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StateStorage Service Spec Suite")
}