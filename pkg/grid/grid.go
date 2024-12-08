package grid

import "github.com/smrqdt/adventofcode-2024/pkg/vector"

type Grid[T comparable] [][]T

func NewGrid[T comparable](x, y int) *Grid[T] {
	m := make(Grid[T], y)
	for i := range m {
		m[i] = make([]T, x)
	}
	return &m
}

func (g Grid[T]) String() string {
	var str string
	for _, row := range g {
		for _, value := range row {
			if value != *new(T) {
				str += "+"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}

func (g Grid[T]) Column(col int) []T {
	column := make([]T, len(g))
	for _, row := range g {
		column = append(column, row[col])
	}
	return column
}

func (g Grid[T]) IsValid(pos vector.Vector) bool {
	if pos.X >= len(g[0]) || pos.Y >= len(g) || pos.X < 0 || pos.Y < 0 {
		return false
	}
	return true
}
