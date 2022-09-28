package transaction

type Direction int

const (
	IN = iota
	OUT
)

func (d Direction) String() string {
	switch d {
	case IN:
		return "IN"
	case OUT:
		return "OUT"
	}
	panic(any("direction: underlying datatype out of range"))
}
