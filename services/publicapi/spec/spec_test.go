package spec

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestPublicApiSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PublicApi Service Spec Suite")
}