package ascii_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestASCII(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ASCII Suite")
}
