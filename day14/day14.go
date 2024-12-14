package main

import (
	"bufio"
	_ "embed"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	g "github.com/smrqdt/adventofcode-2024/pkg/grid"
	"github.com/smrqdt/adventofcode-2024/pkg/helpers"
	v "github.com/smrqdt/adventofcode-2024/pkg/vector"
)

//go:embed input
var input string

func main() {
	log.SetLevel(log.DebugLevel)

	robots := parse(input)
	_ = part1(robots)
	_ = part2(robots)
}

type Robot struct {
	Pos v.Vector
	Vel v.Vector
}

func (r *Robot) Move(seconds int, max v.Vector) {
	r.Pos = r.Pos.Add(r.Vel.Scale(float64(seconds)))
	r.Pos = v.Vector{X: r.Pos.X % max.X, Y: r.Pos.Y % max.Y}
	if r.Pos.X < 0 {
		r.Pos.X += max.X
	}
	if r.Pos.Y < 0 {
		r.Pos.Y += max.Y
	}
}

func (r *Robot) Quadrant(max v.Vector) int {
	west := r.Pos.X < max.X/2
	east := r.Pos.X >= int(math.Ceil(float64(max.X)/2))
	north := r.Pos.Y < max.Y/2
	south := r.Pos.Y >= int(math.Ceil(float64(max.Y)/2))

	switch {
	case north && east:
		return 1
	case north && west:
		return 2
	case south && west:
		return 3
	case south && east:
		return 4
	}
	return 0
}

func parse(input string) (robots []Robot) {
	defer helpers.TrackTime(time.Now(), "parse()")
	scanner := bufio.NewScanner(strings.NewReader(input))
	pattern := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := pattern.FindStringSubmatch(line)
		var nums [4]int
		var err error
		for i, match := range matches[1:] {
			nums[i], err = strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
		}
		robots = append(robots, Robot{
			Pos: v.Vector{X: nums[0], Y: nums[1]},
			Vel: v.Vector{X: nums[2], Y: nums[3]},
		})
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return robots
}

func part1(robots []Robot) int {
	defer helpers.TrackTime(time.Now(), "part1()")

	max := v.Vector{X: 101, Y: 103}
	quadrants := make(map[int]int)

	for _, robot := range robots {
		robot.Move(100, max)
		quadrants[robot.Quadrant(max)]++
		log.Info("Robot", "robot", robot, "quadrant", robot.Quadrant(max))
	}

	safetyFactor := quadrants[1] * quadrants[2] * quadrants[3] * quadrants[4]
	log.Info("Quadrant", "quadrants", quadrants)

	log.Warnf("(Part 1) Safety Factor: %d\n", safetyFactor)
	return safetyFactor
}

func part2(robots []Robot) int {
	defer helpers.TrackTime(time.Now(), "part2()")

	max := v.Vector{X: 101, Y: 103}

	var safetyFactors []int

	for seconds := range 100000 {
		quadrants := make(map[int]int)
		for _, robot := range robots {
			robot.Move(seconds, max)
			quadrants[robot.Quadrant(max)]++
		}
		safetyFactors = append(safetyFactors, quadrants[1]*quadrants[2]*quadrants[3]*quadrants[4])
	}

	minSF := math.MaxInt
	for _, value := range safetyFactors {
		minSF = min(minSF, value)
	}
	// log.Info("", "SFs", safetyFactors, "minSF", minSF)
	seconds := slices.Index(safetyFactors, minSF)
	drawGrid(robots, max, seconds, minSF)

	log.Warnf("(Part 2) \n")
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func drawGrid(robots []Robot, max v.Vector, seconds int, safetyFactor int) {
	grid := g.New[*Robot](max)
	log.Info("Grid", "max", grid.Count())
	for _, robot := range robots {
		robot.Move(seconds, max)
		grid.SetValue(robot.Pos, &robot)
	}
	log.Info("", "grid", grid, "seconds", seconds, "safetyFactor", safetyFactor)
}

func avgRobots(robots []Robot) (avg v.Vector) {
	for _, robot := range robots {
		avg = avg.Add(robot.Pos)
	}
	return avg.Scale(1. / float64(len(robots)))
}

func distance(v1, v2 v.Vector) float64 {
	return v1.Sub(v2).Abs()
}
