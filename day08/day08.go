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
	part2(antennas, mapSize)
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
	for _, antennas := range antennas {
		antinodes := findAntinodes(antennas, mapSize, calcFirstAntinodes)
		for _, antinode := range antinodes {
			locationsWithAntinodes[antinode] = true
		}
	}

	fmt.Printf("(Part 1) %d unique locations with antinodes\n", len(locationsWithAntinodes))
}

func part2(antennas map[rune][]v.Vector, mapSize v.Vector) {
	locatonsWithAntinodes := make(map[v.Vector]bool)
	for _, antennas := range antennas {
		antinodes := findAntinodes(antennas, mapSize, calcAllAntinodes)
		for _, antinode := range antinodes {
			locatonsWithAntinodes[antinode] = true
		}
	}

	fmt.Printf("(Part 2) %d unique locations with antinodes\n", len(locatonsWithAntinodes))
}

func findAntinodes(antennas []v.Vector, mapSize v.Vector, antinodeFunc func(v.Vector, v.Vector, v.Vector) []v.Vector) []v.Vector {
	var antinodes []v.Vector
	for _, antenna1 := range antennas {
		for _, antenna2 := range antennas {
			if antenna1 == antenna2 {
				continue
			}
			antinodes = append(antinodes, antinodeFunc(antenna1, antenna2, mapSize)...)
		}
	}
	return antinodes
}

func calcFirstAntinodes(antenna1, antenna2, mapSize v.Vector) (antinodes []v.Vector) {
	diff := antenna1.Sub(antenna2)
	antinode := antenna1.Add(diff)
	if isInMap(antinode, mapSize) {
		antinodes = append(antinodes, antinode)
	}
	antinode = antenna2.Sub(diff)
	if isInMap(antinode, mapSize) {
		antinodes = append(antinodes, antinode)
	}
	return antinodes
}

func calcAllAntinodes(antenna1, antenna2, mapSize v.Vector) (antinodes []v.Vector) {
	diff := antenna1.Sub(antenna2)
	antinodes = append(antinodes, antenna1, antenna2)
	antinode := antenna1.Add(diff)
	for isInMap(antinode, mapSize) {
		antinodes = append(antinodes, antinode)
		antinode = antinode.Add(diff)
	}
	antinode = antenna2.Sub(diff)
	for isInMap(antinode, mapSize) {
		antinodes = append(antinodes, antinode)
		antinode = antinode.Sub(diff)
	}
	return antinodes
}

func isInMap(vec v.Vector, mapSize v.Vector) bool {
	if vec.X < mapSize.X && vec.X >= 0 && vec.Y < mapSize.Y && vec.Y >= 0 {
		return true
	}
	return false
}
