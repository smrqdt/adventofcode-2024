package vector

import (
	"fmt"
	"math"
)

var (
	NORTH                     = Vector{0, -1}
	EAST                      = Vector{1, 0}
	SOUTH                     = Vector{0, 1}
	WEST                      = Vector{-1, 0}
	DIRECTIONS                = []Vector{NORTH, EAST, SOUTH, WEST}
	NORTH_EAST                = Vector{1, -1}
	SOUTH_EAST                = Vector{1, 1}
	SOUTH_WEST                = Vector{-1, 1}
	NORTH_WEST                = Vector{-1, -1}
	DIAGONALS                 = []Vector{NORTH_EAST, SOUTH_EAST, SOUTH_WEST, NORTH_WEST}
	DIRECTIONS_WITH_DIAGONALS = []Vector{NORTH, NORTH_EAST, EAST, SOUTH_EAST, SOUTH, SOUTH_WEST, WEST, NORTH_WEST}
)

type Vector struct {
	X, Y int
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.X + v2.X, v.Y + v2.Y}
}

func (v Vector) Sub(v2 Vector) Vector {
	return Vector{v.X - v2.X, v.Y - v2.Y}
}

func (v Vector) Scale(i float64) Vector {
	return Vector{int(float64(v.X) * i), int(float64(v.Y) * i)}
}

func (v Vector) RotateRight() Vector {
	return Vector{-v.Y, v.X}
}

func (v Vector) RotateLeft() Vector {
	return Vector{v.Y, -v.X}
}

func (v Vector) Reverse() Vector {
	return Vector{-v.Y, -v.X}
}

func (v Vector) Abs() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v Vector) String() string {
	return fmt.Sprintf("%c Vec(%d|%d)", v.Arrow(), v.X, v.Y)
}

func (v Vector) Arrow() rune {
	if v == (Vector{}) {
		return '•'
	}

	angle := v.Angle()
	octant := int(math.Round(8*angle/(2*math.Pi)+8)) % 8

	arrows := []rune{'⬆', '⬈', '➡', '⬊', '⬇', '⬋', '⬅', '⬉'}
	return arrows[octant]
}

func (v Vector) Angle() float64 {
	return math.Abs(math.Atan2(float64(v.X), float64(v.Y)) - math.Pi)
}
