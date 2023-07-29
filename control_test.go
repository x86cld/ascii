package ascii_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/x86cld/ascii"
)

var _ = Describe("byte code", func() {
	It("defines NUL correctly", func() {
		Expect(ascii.NUL).To(Equal(byte(0x00)))
	})
	It("defines SOH correctly", func() {
		Expect(ascii.SOH).To(Equal(byte(0x01)))
	})
	It("defines STX correctly", func() {
		Expect(ascii.STX).To(Equal(byte(0x02)))
	})
	It("defines ETX correctly", func() {
		Expect(ascii.ETX).To(Equal(byte(0x03)))
	})
	It("defines EOT correctly", func() {
		Expect(ascii.EOT).To(Equal(byte(0x04)))
	})
	It("defines ENQ correctly", func() {
		Expect(ascii.ENQ).To(Equal(byte(0x05)))
	})
	It("defines ACK correctly", func() {
		Expect(ascii.ACK).To(Equal(byte(0x06)))
	})
	It("defines BEL correctly", func() {
		Expect(ascii.BEL).To(Equal(byte(0x07)))
	})
	It("defines BS correctly", func() {
		Expect(ascii.BS).To(Equal(byte(0x08)))
	})
	It("defines HT correctly", func() {
		Expect(ascii.HT).To(Equal(byte(0x09)))
	})
	It("defines LF correctly", func() {
		Expect(ascii.LF).To(Equal(byte(0x0A)))
	})
	It("defines VT correctly", func() {
		Expect(ascii.VT).To(Equal(byte(0x0B)))
	})
	It("defines FF correctly", func() {
		Expect(ascii.FF).To(Equal(byte(0x0C)))
	})
	It("defines CR correctly", func() {
		Expect(ascii.CR).To(Equal(byte(0x0D)))
	})
	It("defines SO correctly", func() {
		Expect(ascii.SO).To(Equal(byte(0x0E)))
	})
	It("defines SI correctly", func() {
		Expect(ascii.SI).To(Equal(byte(0x0F)))
	})
	It("defines DLE correctly", func() {
		Expect(ascii.DLE).To(Equal(byte(0x10)))
	})
	It("defines DC1 correctly", func() {
		Expect(ascii.DC1).To(Equal(byte(0x11)))
	})
	It("defines DC2 correctly", func() {
		Expect(ascii.DC2).To(Equal(byte(0x12)))
	})
	It("defines DC3 correctly", func() {
		Expect(ascii.DC3).To(Equal(byte(0x13)))
	})
	It("defines DC4 correctly", func() {
		Expect(ascii.DC4).To(Equal(byte(0x14)))
	})
	It("defines NAK correctly", func() {
		Expect(ascii.NAK).To(Equal(byte(0x15)))
	})
	It("defines SYN correctly", func() {
		Expect(ascii.SYN).To(Equal(byte(0x16)))
	})
	It("defines ETB correctly", func() {
		Expect(ascii.ETB).To(Equal(byte(0x17)))
	})
	It("defines CAN correctly", func() {
		Expect(ascii.CAN).To(Equal(byte(0x18)))
	})
	It("defines EM correctly", func() {
		Expect(ascii.EM).To(Equal(byte(0x19)))
	})
	It("defines SUB correctly", func() {
		Expect(ascii.SUB).To(Equal(byte(0x1A)))
	})
	It("defines ESC correctly", func() {
		Expect(ascii.ESC).To(Equal(byte(0x1B)))
	})
	It("defines FS correctly", func() {
		Expect(ascii.FS).To(Equal(byte(0x1C)))
	})
	It("defines GS correctly", func() {
		Expect(ascii.GS).To(Equal(byte(0x1D)))
	})
	It("defines RS correctly", func() {
		Expect(ascii.RS).To(Equal(byte(0x1E)))
	})
	It("defines US correctly", func() {
		Expect(ascii.US).To(Equal(byte(0x1F)))
	})
})

var _ = Describe("string code", func() {
	It("defines StringNUL correctly", func() {
		Expect(ascii.StringNUL).To(Equal("\x00"))
	})
	It("defines StringSOH correctly", func() {
		Expect(ascii.StringSOH).To(Equal("\x01"))
	})
	It("defines StringSTX correctly", func() {
		Expect(ascii.StringSTX).To(Equal("\x02"))
	})
	It("defines StringETX correctly", func() {
		Expect(ascii.StringETX).To(Equal("\x03"))
	})
	It("defines StringEOT correctly", func() {
		Expect(ascii.StringEOT).To(Equal("\x04"))
	})
	It("defines StringENQ correctly", func() {
		Expect(ascii.StringENQ).To(Equal("\x05"))
	})
	It("defines StringACK correctly", func() {
		Expect(ascii.StringACK).To(Equal("\x06"))
	})
	It("defines StringBEL correctly", func() {
		Expect(ascii.StringBEL).To(Equal("\x07"))
	})
	It("defines StringBS correctly", func() {
		Expect(ascii.StringBS).To(Equal("\x08"))
	})
	It("defines StringHT correctly", func() {
		Expect(ascii.StringHT).To(Equal("\x09"))
	})
	It("defines StringLF correctly", func() {
		Expect(ascii.StringLF).To(Equal("\x0A"))
	})
	It("defines StringVT correctly", func() {
		Expect(ascii.StringVT).To(Equal("\x0B"))
	})
	It("defines StringFF correctly", func() {
		Expect(ascii.StringFF).To(Equal("\x0C"))
	})
	It("defines StringCR correctly", func() {
		Expect(ascii.StringCR).To(Equal("\x0D"))
	})
	It("defines StringSO correctly", func() {
		Expect(ascii.StringSO).To(Equal("\x0E"))
	})
	It("defines StringSI correctly", func() {
		Expect(ascii.StringSI).To(Equal("\x0F"))
	})
	It("defines StringDLE correctly", func() {
		Expect(ascii.StringDLE).To(Equal("\x10"))
	})
	It("defines StringDC1 correctly", func() {
		Expect(ascii.StringDC1).To(Equal("\x11"))
	})
	It("defines StringDC2 correctly", func() {
		Expect(ascii.StringDC2).To(Equal("\x12"))
	})
	It("defines StringDC3 correctly", func() {
		Expect(ascii.StringDC3).To(Equal("\x13"))
	})
	It("defines StringDC4 correctly", func() {
		Expect(ascii.StringDC4).To(Equal("\x14"))
	})
	It("defines StringNAK correctly", func() {
		Expect(ascii.StringNAK).To(Equal("\x15"))
	})
	It("defines StringSYN correctly", func() {
		Expect(ascii.StringSYN).To(Equal("\x16"))
	})
	It("defines StringETB correctly", func() {
		Expect(ascii.StringETB).To(Equal("\x17"))
	})
	It("defines StringCAN correctly", func() {
		Expect(ascii.StringCAN).To(Equal("\x18"))
	})
	It("defines StringEM correctly", func() {
		Expect(ascii.StringEM).To(Equal("\x19"))
	})
	It("defines StringSUB correctly", func() {
		Expect(ascii.StringSUB).To(Equal("\x1A"))
	})
	It("defines StringESC correctly", func() {
		Expect(ascii.StringESC).To(Equal("\x1B"))
	})
	It("defines StringFS correctly", func() {
		Expect(ascii.StringFS).To(Equal("\x1C"))
	})
	It("defines StringGS correctly", func() {
		Expect(ascii.StringGS).To(Equal("\x1D"))
	})
	It("defines StringRS correctly", func() {
		Expect(ascii.StringRS).To(Equal("\x1E"))
	})
	It("defines StringUS correctly", func() {
		Expect(ascii.StringUS).To(Equal("\x1F"))
	})
})
