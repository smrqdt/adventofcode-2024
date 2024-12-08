package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"

	v "github.com/smrqdt/adventofcode-2024/pkg/vector"
)

//go:embed input
var input string

func main() {
	antennas, mapSize := parse()
	part1(antennas, mapSize)
}

func parse() (map[rune][]v.Vector, v.Vector) {
	antennas := make(map[rune][]v.Vector)
	var mapSize v.Vector

	scanner := bufio.NewScanner(strings.NewReader(input))
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		mapSize = v.Vector{X: len(line), Y: y + 1}
		for x, char := range line {
			if char == '.' {
				continue
			}
			antennas[char] = append(antennas[char], v.Vector{X: x, Y: y})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return antennas, mapSize
}

func part1(antennas map[rune][]v.Vector, mapSize v.Vector) {
	locationsWithAntinodes := make(map[v.Vector]bool)
	for freq, antennas := range antennas {
		antinodes := findAntinodes(freq, antennas, mapSize)
		for _, antinode := range antinodes {
			locationsWithAntinodes[antinode] = true
		}
	}

	fmt.Printf("(Part 1) %d unique locations with antinodes\n", len(locationsWithAntinodes))
}

func findAntinodes(freq rune, antennas []v.Vector, mapSize v.Vector) []v.Vector {
	var antinodes []v.Vector
	for _, antenna1 := range antennas {
		for _, antenna2 := range antennas {
			if antenna1 == antenna2 {
				continue
			}
			antinodeCandidates := calcAntinodes(antenna1, antenna2)
			for _, aC := range antinodeCandidates {
				if isInMap(aC, mapSize) {
					antinodes = append(antinodes, aC)
				}
			}
		}
	}
	return antinodes
}

func calcAntinodes(antenna1, antenna2 v.Vector) []v.Vector {
	diff := antenna1.Sub(antenna2)
	antinode1 := antenna1.Add(diff)
	antinode2 := antenna2.Sub(diff)
	return []v.Vector{antinode1, antinode2}
}

func isInMap(vec v.Vector, mapSize v.Vector) bool {
	if vec.X < mapSize.X && vec.X >= 0 && vec.Y < mapSize.Y && vec.Y >= 0 {
		return true
	}
	return false
}
