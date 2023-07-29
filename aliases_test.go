package ascii_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/x86cld/ascii"
)

var _ = Describe("aliases", func() {
	It("define SoftEOF correctly", func() {
		Expect(ascii.SoftEOF).To(Equal(ascii.SUB))
	})
	It("define StringSoftEOF correctly", func() {
		Expect(ascii.StringSoftEOF).To(Equal(ascii.StringSUB))
	})
})
