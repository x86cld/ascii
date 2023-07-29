package ascii

// Control codes in byte form.
const (
	NUL byte = iota
	SOH      // start of heading
	STX      // start of text
	ETX      // end of text
	EOT      // end of transmission
	ENQ      // enquiry
	ACK      // acknowledge
	BEL      // bell
	BS       // backspace
	HT       // horizontal tab
	LF       // line feed
	VT       // vertical tab
	FF       // form feed
	CR       // carriage return
	SO       // shift out
	SI       // shift in
	DLE      // data link escape
	DC1      // device control 1
	DC2      // device control 2
	DC3      // device control 3
	DC4      // device control 4
	NAK      // negative acknowledge
	SYN      // synchronous idle
	ETB      // end of transmission block
	CAN      // cancel
	EM       // end of medium
	SUB      // substitute
	ESC      // escape
	FS       // file separator
	GS       // group separator
	RS       // record separator
	US       // unit separator
)

// Control codes in string form.
const (
	StringNUL = string(NUL)
	StringSOH = string(SOH)
	StringSTX = string(STX)
	StringETX = string(ETX)
	StringEOT = string(EOT)
	StringENQ = string(ENQ)
	StringACK = string(ACK)
	StringBEL = string(BEL)
	StringBS  = string(BS)
	StringHT  = string(HT)
	StringLF  = string(LF)
	StringVT  = string(VT)
	StringFF  = string(FF)
	StringCR  = string(CR)
	StringSO  = string(SO)
	StringSI  = string(SI)
	StringDLE = string(DLE)
	StringDC1 = string(DC1)
	StringDC2 = string(DC2)
	StringDC3 = string(DC3)
	StringDC4 = string(DC4)
	StringNAK = string(NAK)
	StringSYN = string(SYN)
	StringETB = string(ETB)
	StringCAN = string(CAN)
	StringEM  = string(EM)
	StringSUB = string(SUB)
	StringESC = string(ESC)
	StringFS  = string(FS)
	StringGS  = string(GS)
	StringRS  = string(RS)
	StringUS  = string(US)
)
