package uid12

import (
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_rfc4648Encode(t *testing.T) {
	tests := []struct {
		expect string
		input  Value
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
			fmt.Sprintf("0x%016X", int64(test.input)),
			func(t *testing.T) {
				assert.Equal(t, test.expect, rfc4648Encode(test.input))
				if test.expect != "" {
					assert.Equal(t, stdEncode(test.input), test.expect)
				}
			})
	}
}

var benchValue = Value(time.Now().UnixNano()) & MaxValue

func Benchmark_rfc4648Encode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rfc4648Encode(benchValue)
	}
}

func stdEncode(value Value) string {
	var data [8]byte
	// move 4 bits to the left to work around 8-byte padding behaviour
	binary.BigEndian.PutUint64(data[:], uint64(value)<<4)
	str := base32.StdEncoding.EncodeToString(data[:])[0:12]
	return str
}

func Benchmark_base32StdEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stdEncode(benchValue)
	}
}
