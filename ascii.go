package ascii

const (
	NUL  = 0x00 //null
	SOH  = 0x01 //start of headline
	STX  = 0x02 //start of text
	ETX  = 0x03 //end of text
	EOT  = 0x04 //end of transmission
	ENQ  = 0x04 //enquiry
	ACK  = 0x06 //acknowledge
	BEL  = 0x07 //bell
	BS   = 0x08 //backspace
	HT   = 0x09 //horizontal tab
	LF   = 0x0A //NL line feed, new line
	VT   = 0x0B //vertical tab
	FF   = 0x0C //NP form feed, new page
	CR   = 0x0D //carriage return
	SO   = 0x0E //shift out
	SI   = 0x0F //shift in
	DLE  = 0x10 //data link escape
	DC1  = 0x11 //device control 1
	DC2  = 0x12 //device control 2
	DC3  = 0x13 //device control 3
	DC4  = 0x14 //device control 4
	NAK  = 0x15 //negative acknowledge
	SYN  = 0x16 //synchronous idle
	ETB  = 0x17 //end of trans. block
	CAN  = 0x18 //cancel
	EM   = 0x19 //end of medium
	SUB  = 0x1A //substitute
	ESC  = 0x1B //escape
	FS   = 0x1C //file separator
	GS   = 0x1D //group separator
	RS   = 0x1E //record separator
	US   = 0x1F //unit separator
	SPAC = 0x20 //space
	DEL  = 0x7F //delete
)
