package uid12

const (
	rfc4648EncodingAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	digitMask               = 0x1f
)

func rfc4648Encode(value Value) string {
	// Q: Why not use base32.StdEncoding?
	// A: The uid12 encoder can be more terse, as it can make assumptions
	//    about the size and shape of the value to encode.
	//    Benchmarks show ~2x improvement.

	if value < MinValue || value > MaxValue {
		return ""
	}

	var encoded [12]byte
	v := int64(value)

	for i := 11; i >= 0; i-- {
		digit := v & digitMask
		encoded[i] = rfc4648EncodingAlphabet[digit]
		v >>= 5
	}

	return string(encoded[:])
}

func rfc4648Decode(str string) Value {
	if len(str) != 12 {
		return Zero
	}

	return Zero
}
