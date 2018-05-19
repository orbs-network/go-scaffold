package spec

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestE2ESpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Spec Suite")
}