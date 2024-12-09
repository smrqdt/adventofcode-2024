package vector

import (
	"fmt"
	"math"
)

var (
	North      = Vector{0, -1}
	EAST       = Vector{1, 0}
	SOUTH      = Vector{0, 1}
	WEST       = Vector{-1, 0}
	DIRECTIONS = []Vector{North, EAST, SOUTH, WEST}
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
	return fmt.Sprintf("Vec[%d|%d]", v.X, v.Y)
}
