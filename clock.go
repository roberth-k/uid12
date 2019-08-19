package uid12

import (
	"time"
)

type Clock interface {
	Now() time.Time
}

type globalClock int64

const GlobalClock globalClock = 0

func (globalClock) Now() time.Time {
	return time.Now()
}
