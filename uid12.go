package uid12

import (
	"fmt"
)

var globalSource = NewSource()

func Generate() Value {
	return globalSource.Generate()
}

func Parse(s string) (Value, error) {
	v := rfc4648Decode(s)

	if v == Zero {
		return Zero, fmt.Errorf("invalid uid12.Value: %s", s)
	}

	if !v.Check() {
		return Zero, fmt.Errorf("invalid uid12.Value: %s (bad checksum)", s)
	}

	return v, nil
}

func MustParse(s string) Value {
	v, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return v
}
