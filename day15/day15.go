package main

import (
	"bufio"
	_ "embed"
	"maps"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/smrqdt/adventofcode-2024/pkg/grid"
	"github.com/smrqdt/adventofcode-2024/pkg/helpers"
	v "github.com/smrqdt/adventofcode-2024/pkg/vector"
)

//go:embed input
var input string

const (
	WALL  = '#'
	BOX   = 'O'
	ROBOT = '@'
)

func main() {
	log.SetLevel(log.DebugLevel)

	objects, robot, movements := parse(input)
	_ = part1(maps.Clone(objects), robot, movements)
}

func parse(input string) (objects map[v.Vector]rune, robot v.Vector, movements []v.Vector) {
	defer helpers.TrackTime(time.Now(), "parse()")

	segments := strings.Split(input, "\n\n")
	objects, robot = parseMap(segments[0])
	movements = parseMovements(segments[1])
	return
}

func parseMap(input string) (objects map[v.Vector]rune, robot v.Vector) {
	defer helpers.TrackTime(time.Now(), "parseMap()")
	objects = make(map[v.Vector]rune)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		for x, char := range line {
			switch char {
			case '.':
				// do nothing
			case '@':
				robot = v.Vector{X: x, Y: y}
				fallthrough
			case WALL, BOX:
				objects[v.Vector{X: x, Y: y}] = char
			default:
				panic("Unknown char")
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return objects, robot
}

func parseMovements(input string) (movements []v.Vector) {
	defer helpers.TrackTime(time.Now(), "parseMovements()")
	movements = make([]v.Vector, 0, len(input))
	for _, char := range input {
		switch char {
		case '^':
			movements = append(movements, v.NORTH)
		case '>':
			movements = append(movements, v.EAST)
		case 'v':
			movements = append(movements, v.SOUTH)
		case '<':
			movements = append(movements, v.WEST)
		case '\n':
			// skip
		default:
			panic("Unknown char")
		}
	}
	return movements
}

func part1(objects map[v.Vector]rune, robot v.Vector, movements []v.Vector) int {
	defer helpers.TrackTime(time.Now(), "part1()")

	for _, movement := range movements {
		robot, _ = moveObject(robot, '@', movement, objects)
	}

	var coordinatesSum int
	for position, object := range objects {
		if object == BOX {
			coordinatesSum += 100*position.Y + position.X
		}
	}
	printMap(objects)
	log.Warnf("(Part 1) Sum of all Coordinates: %d \n", coordinatesSum)
	return coordinatesSum
}

func moveObject(position v.Vector, object rune, movement v.Vector, objects map[v.Vector]rune) (v.Vector, bool) {
	if object == WALL {
		// Walls don’t move
		return position, false
	}

	newPos := position.Add(movement)
	success := true
	if obstacle, ok := objects[newPos]; ok {
		_, success = moveObject(newPos, obstacle, movement, objects)
	}
	if !success {
		return position, false
	}
	delete(objects, position)
	objects[newPos] = object
	return newPos, true
}

func printMap(objects map[v.Vector]rune) {
	var maxX, maxY int

	for vec := range objects {
		maxX = max(maxX, vec.X)
		maxY = max(maxY, vec.Y)
	}

	objectMap := grid.New[rune](v.Vector{X: maxX + 1, Y: maxY + 1})

	for vec, value := range objects {
		objectMap.SetValue(vec, value)
	}

	log.Info("Current Map", "map", objectMap)
}
