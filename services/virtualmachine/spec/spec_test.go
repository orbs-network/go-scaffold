package spec

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestVirtualMachineSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VirtualMachine Service Spec Suite")
}