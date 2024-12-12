package grid

import (
	"bufio"
	"errors"
	"fmt"
	"iter"
	"strings"

	"github.com/smrqdt/adventofcode-2024/pkg/convert"
	"github.com/smrqdt/adventofcode-2024/pkg/vector"
)

var OutOfBoundsError error = fmt.Errorf("Coordinates outside of grid")

type Grid[T comparable] [][]T

func New[T comparable](v vector.Vector) Grid[T] {
	m := make(Grid[T], v.Y)
	for i := range m {
		m[i] = make([]T, v.X)
	}
	return m
}

func NewFromInput[T comparable](input string, mapFunc convert.ConvertFunc[rune, T]) (grid Grid[T], err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, make([]T, len(line)))
		for i, char := range line {
			grid[len(grid)-1][i], err = mapFunc(char)
			if err != nil {
				return nil, err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error during scanning input: %w", err)
	}
	return grid, nil
}

func (g Grid[T]) String() string {
	var str string
	for _, row := range g {
		for _, value := range row {
			switch v := any(value).(type) {
			case rune, byte:
				str = fmt.Sprintf("%s%c", str, v)
			case int:
				str = fmt.Sprintf("%s %3d", str, v)
			default:
				if value != *new(T) {
					str += "+"
				} else {
					str += "."
				}
			}
		}
		str += "\n"
	}
	return str
}

func (g Grid[T]) Count() vector.Vector {
	return vector.Vector{X: len(g[0]), Y: len(g)}
}

func (g Grid[T]) Value(v vector.Vector) (T, error) {
	if !g.IsValid(v) {
		return *new(T), fmt.Errorf("Vector %v is not in Grid %v: %w", v, g, OutOfBoundsError)
	}
	return g[v.Y][v.X], nil
}

func (g Grid[T]) SetValue(v vector.Vector, value T) error {
	if !g.IsValid(v) {
		return fmt.Errorf("Vector %v is not in Grid %v: %w", v, g, OutOfBoundsError)
	}
	g[v.Y][v.X] = value
	return nil
}

func (g Grid[T]) Values(vs []vector.Vector) (values []T, err error) {
	for _, v := range vs {
		value, err := g.Value(v)
		if errors.Is(err, OutOfBoundsError) {
			continue
		}
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

func (g Grid[T]) Column(col int) []T {
	column := make([]T, len(g))
	for _, row := range g {
		column = append(column, row[col])
	}
	return column
}

func (g Grid[T]) All() iter.Seq2[vector.Vector, T] {
	return func(yield func(vector.Vector, T) bool) {
		for y, row := range g {
			for x, cell := range row {
				if !yield(vector.Vector{X: x, Y: y}, cell) {
					return
				}
			}
		}
	}
}

func (g Grid[T]) IsValid(pos vector.Vector) bool {
	if pos.X >= len(g[0]) || pos.Y >= len(g) || pos.X < 0 || pos.Y < 0 {
		return false
	}
	return true
}

func (g Grid[T]) FindAll(toFind T) (found []vector.Vector) {
	for y, row := range g {
		for x, cell := range row {
			if cell == toFind {
				found = append(found, vector.Vector{X: x, Y: y})
			}
		}
	}
	return
}

func (g Grid[T]) GetNeighbour(v, dir vector.Vector) (neigh vector.Vector, exists bool, err error) {
	if !g.IsValid(v) {
		return vector.Vector{}, false, fmt.Errorf("Vector %v is not in Grid %#v: %w", v, g, OutOfBoundsError)
	}
	neigh = v.Add(dir)
	if !g.IsValid(neigh) {
		return neigh, false, nil
	}
	return neigh, true, nil
}

func (g Grid[T]) GetNeighbourValue(v, dir vector.Vector) (nVec vector.Vector, value T, exists bool, err error) {
	nVec, exists, err = g.GetNeighbour(v, dir)
	if err != nil || !exists {
		return nVec, *new(T), exists, err
	}
	value, err = g.Value(nVec)
	if err != nil {
		return nVec, value, exists, err
	}
	return
}

func (g Grid[T]) GetNeighbours(v vector.Vector) (neighbours []vector.Vector, err error) {
	for _, dir := range vector.DIRECTIONS {
		neigh, exists, err := g.GetNeighbour(v, dir)
		if err != nil {
			return nil, err
		}
		if exists {
			neighbours = append(neighbours, neigh)
		}
	}
	return neighbours, nil
}

func (g Grid[T]) GetNeighbourValues(v vector.Vector) (neighbours []T, err error) {
	neighVectors, err := g.GetNeighbours(v)
	if err != nil {
		return nil, err
	}
	neighbours, err = g.Values(neighVectors)
	if err != nil {
		return nil, err
	}

	return neighbours, nil
}
