package uid12

const digitMask = 0x1f

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

	v := int64(0)

	ch := str[0]
	digit := rfc4648DecodingAlphabet[ch]
	v |= digit
	if digit == 0 && ch != 'A' && ch != 'a' {
		return Zero
	}

	for i := 1; i < 12; i++ {
		ch = str[i]
		digit = rfc4648DecodingAlphabet[ch]
		v = (v << 5) | digit

		if digit == 0 && ch != 'A' && ch != 'a' {
			return Zero
		}
	}

	return Value(v)
}

const rfc4648EncodingAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

var rfc4648DecodingAlphabet = [256]int64{
	'0': 14, // alias of O
	'1': 8,  // alias of I
	'2': 26,
	'3': 27,
	'4': 28,
	'5': 29,
	'6': 30,
	'7': 31,
	'8': 1,  // alias of B
	'9': 15, // alias of P
	'A': 0,
	'B': 1,
	'C': 2,
	'D': 3,
	'E': 4,
	'F': 5,
	'G': 6,
	'H': 7,
	'I': 8,
	'J': 9,
	'K': 10,
	'L': 11,
	'M': 12,
	'N': 13,
	'O': 14,
	'P': 15,
	'Q': 16,
	'R': 17,
	'S': 18,
	'T': 19,
	'U': 20,
	'V': 21,
	'W': 22,
	'X': 23,
	'Y': 24,
	'Z': 25,
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
	'i': 8,
	'j': 9,
	'k': 10,
	'l': 11,
	'm': 12,
	'n': 13,
	'o': 14,
	'p': 15,
	'q': 16,
	'r': 17,
	's': 18,
	't': 19,
	'u': 20,
	'v': 21,
	'w': 22,
	'x': 23,
	'y': 24,
	'z': 25,
}
