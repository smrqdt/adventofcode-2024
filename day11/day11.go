package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	stones := parse()
	part1(stones)
}

func parse() []int {
	var stones []int
	fields := strings.Fields(input)
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			log.Fatal(err)
		}
		stones = append(stones, num)
	}
	return stones
}

func part1(stones []int) {
	for r := range 25 {
		fmt.Printf("Round %3d: ", r)
		var inserts int
		for i := range len(stones) {
			stone := stones[i+inserts]
			digits := int(math.Log10(float64(stone))) + 1
			switch {
			case stones[i+inserts] == 0:
				stones[i+inserts] = 1
			case digits%2 == 0 && stone != 1:
				splitFactor := int(math.Pow10(digits / 2))
				stones[i+inserts] = stone / splitFactor
				inserts++
				stones = slices.Insert(stones, i+inserts, stone%splitFactor)
			default:
				stones[i+inserts] *= 2024
			}
		}
		fmt.Printf("%6d\n", len(stones))
	}
	fmt.Printf("(Part 1) %d \n", len(stones))
}
