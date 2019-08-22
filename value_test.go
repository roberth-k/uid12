package uid12

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const marker = Value(0x123456789abcdef0)

func TestConstants(t *testing.T) {
	assert.Equal(t, int64(0), int64(Zero))
	assert.Equal(t, Zero, MinValue)
	mask := int64(0x0fffffffffffffc0)
	assert.Equal(t, mask+(mask%37), int64(MaxValue))
}

func TestValue_String(t *testing.T) {
	assert.Equal(t, "AAAAAAAAAAAA", Zero.String())
}

func TestValue_Check(t *testing.T) {
	assert.True(t, Zero.Check())
}

func TestValue_Time(t *testing.T) {
	assert.Equal(t, time.Unix(0, 0), Zero.Time())
	// shift right to bring back the sign bit that was shifted off.
	assert.Equal(t, time.Unix(0x23456789>>1, 0), marker.Time())
}

func TestValue_Nonce(t *testing.T) {
	assert.Equal(t, int64(0), Zero.Nonce())
	assert.Equal(t, int64(0x9abcdec0), marker.Nonce())
}

func TestValue_Checksum(t *testing.T) {
	assert.Equal(t, int64(0), Zero.Checksum())
	assert.Equal(t, int64(0x30), marker.Checksum())
}
