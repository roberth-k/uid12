package uid12

import (
	"math"
	"math/rand"
)

func NewSource() *Source {
	return NewSourceWith(nil, nil)
}

func NewSourceSeeded() *Source {
	return NewSourceWith(
		rand.NewSource(GlobalClock.Now().UnixNano()),
		GlobalClock)
}

func NewSourceWith(random rand.Source, clock Clock) *Source {
	if clock == nil {
		clock = GlobalClock
	}

	if random == nil {
		random = GlobalRand
	}

	return &Source{random, clock}
}

type Source struct {
	random rand.Source
	clock  Clock
}

func (source Source) Generate() Value {
	// 31 bits of time
	ts := source.clock.Now().Unix() & math.MaxInt32

	// 23 bits of entropy
	nonce := source.random.Int63() & 0x1fffffc0

	// as the ts is signed, we gain one bit by shifting off the sign
	value := (ts << 29) | nonce

	// 6 bits of checksum
	checksum := value % 37
	value = value | (checksum & 0x3f)

	return Value(value)
}
