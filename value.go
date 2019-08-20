package uid12

type Value int64

const (
	Zero     Value = 0
	MinValue Value = 0

	// MaxValue is the maximal uid12 value. It contains exactly 60 bits
	// (5 bits per base32 digit). The true max value is 0x0fffffffffffffc0
	// plus its modulo 37 from the checksum.
	MaxValue Value = 0x0fffffffffffffd4
)

func (value Value) String() string {
	return rfc4648Encode(value)
}
