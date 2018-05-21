package stress

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestE2EStress(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2E Stress Suite")
}