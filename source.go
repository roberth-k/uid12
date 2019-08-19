package uid12

import (
	"math"
	"math/rand"
)

var globalSource = NewSource()

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
	ts := source.clock.Now().Unix() & math.MaxInt32
	nonce := (source.random.Int63() << 6) & 0xffffffc0
	value := (ts << 32) | nonce
	checksum := value % 37
	value = value | (checksum & 0x3f)
	return Value(value)
}
