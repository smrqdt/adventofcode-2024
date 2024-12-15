package main

import (
	"bufio"
	_ "embed"
	"fmt"
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
	log.SetTimeFormat("")

	objects, robot, movements := parse(input)
	_ = part1(maps.Clone(objects), robot, movements)
	_ = part2(maps.Clone(objects), robot, movements)
}

func parse(input string) (objects map[v.Vector]rune, robot v.Vector, movements []v.Vector) {
	defer helpers.TrackTime(time.Now(), "parse()")

	segments := strings.Split(input, "\n\n")
	objects, robot = parseMap(segments[0])
	movements = parseMovements(segments[1])
	log.Info("Parsing finished", "objects", len(objects), "robot", robot, "movements", len(movements))
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
	printMap(objects, v.Vector{}, false)
	log.Warnf("(Part 1) Sum of all Coordinates: %d \n", coordinatesSum)
	return coordinatesSum
}

func part2(objects map[v.Vector]rune, robot v.Vector, movements []v.Vector) int {
	defer helpers.TrackTime(time.Now(), "part2()")

	scaledObjects := make(map[v.Vector]rune)
	var maxX, maxY int
	for vec, rune := range objects {
		newVec := v.Vector{X: vec.X * 2, Y: vec.Y}
		scaledObjects[newVec] = rune
		maxX = max(maxX, newVec.X)
		maxY = max(maxY, newVec.Y)
	}
	robot.X *= 2

	// printMap(scaledObjects, v.Vector{}, true)
	fmt.Print("\033[2J")
	for i, movement := range movements {
		newRobot, success, execMove := moveWideObject(robot, '@', movement, scaledObjects)
		if success {
			execMove()
			robot = newRobot
		}
		if log.GetLevel() <= log.DebugLevel {
			str := printMap(scaledObjects, movement, true)
			fmt.Print("\033[H")
			log.Debug("Move", "map", str, "i", i, "position", robot, "movement", movement)
			// time.Sleep(10 * time.Millisecond)
			// fmt.Scanln()
		}
	}

	var coordinatesSum int
	for position, object := range scaledObjects {
		if object == BOX {
			coordinatesSum += 100*position.Y + position.X
		}
	}

	// printMap(scaledObjects, v.Vector{}, true)
	log.Warnf("(Part 2) Sum of all Coordinates: %d", coordinatesSum)
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

func moveWideObject(position v.Vector, object rune, movement v.Vector, objects map[v.Vector]rune) (v.Vector, bool, func()) {
	if object == WALL {
		// Walls don’t move
		return position, false, func() {}
	}

	nextPos := position.Add(movement)
	obstaclesPos := []v.Vector{nextPos}

	switch movement {
	case v.NORTH, v.SOUTH:
		obstaclesPos = append(obstaclesPos, nextPos.Add(v.WEST))

		if object == BOX {
			obstaclesPos = append(obstaclesPos, nextPos.Add(v.EAST))
		}
	case v.EAST:
		if object == BOX {
			obstaclesPos = []v.Vector{nextPos.Add(v.EAST)}
		}
	case v.WEST:
		obstaclesPos = append(obstaclesPos, nextPos.Add(v.WEST))
	}

	allSuccess := true
	var executeFuncs []func()

	for _, pos := range obstaclesPos {
		if obstacle, ok := objects[pos]; ok {
			var execute func()
			var success bool
			_, success, execute = moveWideObject(pos, obstacle, movement, objects)
			allSuccess = allSuccess && success
			executeFuncs = append(executeFuncs, execute)
		}
	}
	if !allSuccess {
		return position, false, func() {}
	}

	returnFunc := func() {
		for _, f := range executeFuncs {
			f()
		}
		delete(objects, position)
		objects[nextPos] = object
	}
	return nextPos, true, returnFunc
}

func printMap(objects map[v.Vector]rune, movement v.Vector, doubleWidth bool) string {
	var maxX, maxY int

	for vec := range objects {
		maxX = max(maxX, vec.X)
		maxY = max(maxY, vec.Y)
	}

	overflow := 1
	if doubleWidth {
		overflow = 2
	}

	objectMap := grid.New[rune](v.Vector{X: maxX + overflow, Y: maxY + 1})

	for vec, value := range objects {
		objectMap.SetValue(vec, value)
		if doubleWidth {
			switch value {
			case 'O':
				objectMap.SetValue(vec, '[')
				objectMap.SetValue(vec.Add(v.EAST), ']')
			case '@':
				objectMap.SetValue(vec, movement.Arrow())
			case '#':
				objectMap.SetValue(vec.Add(v.EAST), '#')
			}
		}
	}

	return objectMap.String()
}
