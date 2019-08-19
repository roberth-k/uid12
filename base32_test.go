package uid12

import (
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_rfc4648(t *testing.T) {
	assert.Equal(t, "", rfc4648Encode(Zero))
	assert.Equal(t, Zero, rfc4648Decode(""))

	tests := []struct {
		s string
		v Value
	}{
		{"", Zero},
		{"AAAAAAAAAAAB", MinValue},
		{"777777777777", MaxValue},
		{"", -1},
		{"", 0x4000000000000001},
		{"", 0x2000000000000001},
		{"", 0x1000000000000001},
		{"QAAAAAAAAAAB", 0x0800000000000001},
		{"CEIRCEIRCEIR", 0x0111111111111111},
	}

	for _, test := range tests {
		t.Run(
			fmt.Sprintf("%s 0x%016X", test.s, int64(test.v)),
			func(t *testing.T) {
				assert.Equal(t, test.s, rfc4648Encode(test.v))

				if test.s != "" {
					assert.Equal(t, test.v, rfc4648Decode(test.s))
				}

				if test.s != "" && test.v != Zero {
					assert.Equal(t, stdEncode(test.v), test.s)
					assert.Equal(t, stdDecode(test.s), test.v)
				}
			})
	}
}

func stdEncode(value Value) string {
	var data [8]byte
	// move 4 bits to the left to work around 8-byte padding behaviour
	binary.BigEndian.PutUint64(data[:], uint64(value)<<4)
	str := base32.StdEncoding.EncodeToString(data[:])
	return str[0:12]
}

func stdDecode(str string) Value {
	// todo: is the string append slowing the benchmark?
	decoded, _ := base32.StdEncoding.DecodeString(str + "A===")
	return Value(binary.BigEndian.Uint64(decoded[0:8]) >> 4)
}

var (
	encodeBenchValue = Value(time.Now().UnixNano()) & MaxValue
	decodeBenchValue = stdEncode(encodeBenchValue)
)

func Benchmark_rfc4648Encode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rfc4648Encode(encodeBenchValue)
	}
}

func Benchmark_base32StdEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stdEncode(encodeBenchValue)
	}
}

func Benchmark_rfc4648Decode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rfc4648Decode(decodeBenchValue)
	}
}

func Benchmark_base32StdDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stdDecode(decodeBenchValue)
	}
}
