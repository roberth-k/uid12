package uid12

type Value int64

const (
	Zero     Value = 0
	MinValue Value = 1
	MaxValue Value = 0x0fffffffffffffff // 60 bits for 12 base32 digits
)

func (value Value) String() string {
	return rfc4648Encode(value)
}
