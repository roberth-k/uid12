package uid12

import (
	"math/rand"
)

type globalRand int64

const GlobalRand globalRand = 0

func (globalRand) Int63() int64 {
	return rand.Int63()
}

func (globalRand) Seed(seed int64) {
	rand.Seed(seed)
}
