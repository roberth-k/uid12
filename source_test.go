package uid12

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

type fixedClock struct {
	t time.Time
}

func (clock fixedClock) Now() time.Time {
	return clock.t
}

type fixedRandom struct {
	v int64
}

func (random fixedRandom) Int63() int64 {
	return random.v
}

func (random fixedRandom) Seed(seed int64) {
	random.v = seed
}

const (
	all        = math.MaxInt64
	nochecksum = int64(0x7fffffffffffffc0)
	checksum   = int64(0x3f)
)

func TestSource_Generate(t *testing.T) {
	tests := []struct {
		e int64     // expect
		t time.Time // time
		r int64     // random
		m int64     // mask
	}{
		{0x0000000000000000, time.Unix(0, 0), 0, all},
		{0x0000000000000040, time.Unix(0, 0), 0x40, nochecksum},
		{0x000000000000005b, time.Unix(0, 0), 0x40, all},
		{0x000000000000001b, time.Unix(0, 0), 0x40, checksum},
		{0x000000001fffffc0, time.Unix(0, 0), math.MaxInt64, nochecksum},
		{0x0fffffffe0000000, time.Unix(math.MaxInt32, 0), 0, nochecksum},
	}

	for i, test := range tests {
		t.Run(
			fmt.Sprintf("%d: %s (0x%016X)", i, Value(test.e), test.e),
			func(t *testing.T) {
				clock := fixedClock{test.t}
				random := fixedRandom{test.r}
				source := NewSourceWith(random, clock)

				v := source.Generate()
				assert.Equal(t, test.e, int64(v)&test.m, "got 0x%016X", int64(v)&test.m)
				assert.Equal(t, test.t, v.Time())
			})
	}
}
