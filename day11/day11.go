package main

import (
	_ "embed"
	"fmt"
	"iter"
	"log"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	stones := parse()
	part1(slices.Clone(stones))
	part2(slices.Clone(stones))
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
	rounds := 25
	for r := range rounds {
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
	fmt.Printf("(Part 1) Number of stones after %d rounds: %d\n", rounds, len(stones))
}

func part2(stones []int) {
	rounds := 75

	stoneCountsA := make(map[int]int)
	stoneCountsB := make(map[int]int)
	for _, stone := range stones {
		stoneCountsA[stone]++
	}

	source := &stoneCountsA
	target := &stoneCountsB
	fmt.Println(source)
	for round := range rounds {
		fmt.Printf("Round %2d: ", round)
		clear(*target)
		for stone, count := range *source {
			digits := int(math.Log10(float64(stone))) + 1
			switch {
			case stone == 0:
				(*target)[1] += count
			case digits%2 == 0 && stone != 1:
				splitFactor := int(math.Pow10(digits / 2))
				(*target)[stone/splitFactor] += count
				(*target)[stone%splitFactor] += count
			default:
				(*target)[stone*2024] += count
			}
		}
		fmt.Printf("%18d \n", sum(maps.Values(*target)))
		source, target = target, source
	}
	fmt.Printf("(Part 2) Number of stones after %d rounds: %d\n", rounds, sum(maps.Values(*source)))
}

func sum(seek iter.Seq[int]) (sum int) {
	for value := range seek {
		sum += value
	}
	return
}
