package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func main() {
	guard, obstructions := parse()
	part1(guard, obstructions)
}

type Vector struct {
	x, y int
}

func (v *Vector) Add(v2 Vector) {
	v.x += v2.x
	v.y += v2.y
}

func (v *Vector) Sub(v2 Vector) {
	v.x -= v2.x
	v.y -= v2.y
}

var (
	UP    = Vector{0, -1}
	RIGHT = Vector{1, 0}
	DOWN  = Vector{0, 1}
	LEFT  = Vector{-1, 0}
)

func parse() (Guard, [][]bool) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var obstructions Map
	var guard Guard

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]bool, len(line))
		for col, char := range line {
			switch char {
			case '#':
				row[col] = true
			case '^':
				guard = Guard{position: Vector{col, len(obstructions)}, direction: UP}
			}
		}
		obstructions = append(obstructions, row)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// fmt.Println(obstructions)

	return guard, obstructions
}

func part1(guard Guard, obstructions [][]bool) {
	visited := make(Map, len(obstructions))
	for i := range visited {
		visited[i] = make([]bool, len(obstructions[0]))
	}

	visited[guard.position.y][guard.position.x] = true
	for guard.step(obstructions) {
		visited[guard.position.y][guard.position.x] = true
	}

	fmt.Println(visited)
	var visitedCount int
	for _, row := range visited {
		for _, value := range row {
			if value {
				visitedCount++
			}
		}
	}

	fmt.Printf("(Part 1) The guard visited %d distinct locations \n", visitedCount)
}

type Guard struct {
	position  Vector
	direction Vector
}

func (g *Guard) step(obstructions [][]bool) bool {
	// fmt.Println(g)
	g.position.Add(g.direction)
	if g.position.y < 0 || g.position.y >= len(obstructions) ||
		g.position.x < 0 || g.position.x >= len(obstructions[0]) {
		return false
	}
	if obstructions[g.position.y][g.position.x] {
		g.position.Sub(g.direction)
		g.turn()
		g.step(obstructions)
	}
	return true
}

func (g *Guard) turn() {
	switch g.direction {
	case UP:
		g.direction = RIGHT
	case RIGHT:
		g.direction = DOWN
	case DOWN:
		g.direction = LEFT
	case LEFT:
		g.direction = UP
	}
}

func (g Guard) String() string {
	var dirSym string
	switch g.direction {
	case UP:
		dirSym = "↑"
	case DOWN:
		dirSym = "↓"
	case LEFT:
		dirSym = "←"
	case RIGHT:
		dirSym = "→"
	default:
		dirSym = fmt.Sprint(g.direction)
	}
	return fmt.Sprintf("Guard{%d, %d, %v}", g.position.x, g.position.y, dirSym)
}

type Map [][]bool

func (m Map) String() string {
	var str string
	for _, row := range m {
		for _, value := range row {
			if value {
				str += "#"
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}
