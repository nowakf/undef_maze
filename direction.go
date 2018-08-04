package main

type direction int

var directions = map[vector]direction{
	vector{0, 0}:   SELF,
	vector{0, 1}:   N,
	vector{0, -1}:  S,
	vector{1, 0}:   E,
	vector{-1, 0}:  W,
	vector{1, 1}:   NE,
	vector{1, -1}:  NW,
	vector{-1, 1}:  SE,
	vector{-1, -1}: SW,
}

func (d direction) String() string {
	switch d {
	case N:
		return "N"
	case S:
		return "S"
	case E:
		return "E"
	case W:
		return "W"
	case SELF:
		return "SELF"
	case Cardinal:
		return "Cardinal"
	case NE:
		return "NE"
	case NW:
		return "NW"
	case SE:
		return "SE"
	case SW:
		return "SW"
	default:
		return "Error"
	}
}
func (d direction) ToVec() vector {
	switch d {
	case SELF:
		return vector{0, 0}
	case N:
		return vector{0, 1}
	case S:
		return vector{0, -1}
	case E:
		return vector{1, 0}
	case W:
		return vector{-1, 0}
	case NE:
		return vector{1, 1}
	case NW:
		return vector{1, -1}
	case SE:
		return vector{-1, 1}
	case SW:
		return vector{-1, -1}
	default:
		panic("nonexistent direction")
	}
}

const (
	N direction = iota
	S
	E
	W
	SELF
	Cardinal
	NE
	NW
	SE
	SW
)
