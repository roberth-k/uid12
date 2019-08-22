package uid12

import (
	"time"
)

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

func (value Value) Check() bool {
	checksum := value & 0x3f
	v := value & 0x7fffffffffffffc0
	return v%37 == checksum
}

func (value Value) Time() time.Time {
	ts := (int64(value) >> 29) & 0x7fffffff
	return time.Unix(ts, 0)
}

func (value Value) Nonce() int64 {
	return int64(value) & 0xffffffc0
}

func (value Value) Checksum() int64 {
	return int64(value) & 0x3f
}
