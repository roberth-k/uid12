package uid12

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValue_MarshalJSON(t *testing.T) {
	value := Generate()
	data, err := json.Marshal(value)
	require.NoError(t, err)
	require.Equal(t, fmt.Sprintf(`"%s"`, value), string(data))
}

func TestValue_UnmarshalJSON(t *testing.T) {
	value := Generate()
	data := fmt.Sprintf(`"%s"`, value)
	var value1 Value
	require.NoError(t, json.Unmarshal([]byte(data), &value1))
	require.Equal(t, value, value1)
}
