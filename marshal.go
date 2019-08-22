package uid12

import (
	"encoding/json"
	"fmt"
)

func (value Value) MarshalJSON() ([]byte, error) {
	if !value.Check() {
		return nil, fmt.Errorf("invalid uid12.Value: %s (bad checksum)", value)
	}

	encoded := make([]byte, 14)
	encoded[0] = '"'
	encoded[13] = '"'
	rfc4648EncodeTo(value, encoded[1:])

	return encoded, nil
}

func (value *Value) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	v, err := Parse(s)
	if err != nil {
		return err
	}

	*value = v
	return nil
}
