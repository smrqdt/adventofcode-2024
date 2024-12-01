package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	listA, listB := parse()
	part1(listA, listB)
}

//go:embed input
var input string

func parse() ([]int, []int) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var listA, listB []int

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		valA, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		valB, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		listA = append(listA, valA)
		listB = append(listB, valB)
	}
	return listA, listB
}

func part1(listA, listB []int) {
	var resultSum int

	slices.Sort(listA)
	slices.Sort(listB)

	for i := range listA {
		diff := listA[i] - listB[i]
		if diff > 0 {
			resultSum += diff
		} else {
			resultSum -= diff
		}
	}

	fmt.Printf("(Part 1) Total Distance: %d \n", resultSum)
}
