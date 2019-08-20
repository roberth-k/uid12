package uid12

import (
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func Test_rfc4648Encode(t *testing.T) {
	tests := []struct {
		s string
		v Value
	}{
		{"AAAAAAAAAAAA", Zero},
		{"AAAAAAAAAAAA", MinValue},
		{"77777777776U", MaxValue},
		{"", -1},
		{"", 0x4000000000000001},
		{"", 0x2000000000000001},
		{"", 0x1000000000000001},
		{"QAAAAAAAAAAB", 0x0800000000000001},
		{"CEIRCEIRCEIR", 0x0111111111111111},
	}

	for i, test := range tests {
		t.Run(
			fmt.Sprintf("%d: %s - 0x%016X", i, test.s, int64(test.v)),
			func(t *testing.T) {
				assert.Equal(t, test.s, rfc4648Encode(test.v))
				if test.s != "" && test.v != Zero {
					assert.Equal(t, stdEncode(test.v), test.s)
				}
			})
	}
}

func Test_rfc4648Decode(t *testing.T) {
	tests := []struct {
		v Value
		s string
	}{
		{Zero, ""},
		{MinValue, "AAAAAAAAAAAA"},
		{MaxValue, "77777777776U"},
		{0x0800000000000001, "QAAAAAAAAAAB"},
		{0x0111111111111111, "CEIRCEIRCEIR"},
		{0x0111111111111111, "CE1RCE1RCE1R"},
		{Zero, "@EIRCEIRCEIR"},
		{Zero, "CEIRCE@RCEIR"},
		{Zero, "CEIRCEIRCEI@"},
	}

	for i, test := range tests {
		t.Run(
			fmt.Sprintf("%d: 0x%016X - %s", i, int64(test.v), test.s),
			func(t *testing.T) {
				assert.Equal(t, test.v, rfc4648Decode(test.s))
				if test.s != "" && test.v != Zero {
					// the base32.StdEncoding doesn't support the aliases
					s := strings.ReplaceAll(test.s, "1", "I")
					s = strings.ReplaceAll(s, "0", "O")
					s = strings.ReplaceAll(s, "8", "B")
					s = strings.ReplaceAll(s, "9", "P")
					assert.Equal(t, stdDecode(s), test.v)
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
