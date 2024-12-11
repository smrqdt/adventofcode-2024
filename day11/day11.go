package main

import (
	_ "embed"
	"iter"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/log"

	"github.com/smrqdt/adventofcode-2024/pkg/helpers"
)

//go:embed input
var input string

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetTimeFormat(time.TimeOnly)

	stones := parse()
	part1(slices.Clone(stones))
	part2(slices.Clone(stones))
}

func parse() []int {
	defer helpers.TrackTime(time.Now(), "parse()")
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
	defer helpers.TrackTime(time.Now(), "part1()")

	rounds := 25
	for round := range rounds {
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
		log.Info("Round finished", "round", round, "stones", len(stones))
	}
	log.Warnf("(Part 1) Number of stones after %d rounds: %d", rounds, len(stones))
}

func part2(stones []int) {
	defer helpers.TrackTime(time.Now(), "part2()")

	rounds := 75

	stoneCountsA := make(map[int]int)
	stoneCountsB := make(map[int]int)
	for _, stone := range stones {
		stoneCountsA[stone]++
	}

	source := &stoneCountsA
	target := &stoneCountsB
	log.Debug(source)
	for round := range rounds {
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
		log.Info("Round finished", "round", round, "stones", sum(maps.Values(*target)))
		source, target = target, source
	}
	log.Printf("(Part 2) Number of %d different stones after %d rounds: %d", len(*source), rounds, sum(maps.Values(*source)))
}

func sum(seek iter.Seq[int]) (sum int) {
	for value := range seek {
		sum += value
	}
	return
}
