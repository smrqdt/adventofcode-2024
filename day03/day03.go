package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input
var input string

func main() {
	part1()
}

func part1() {

	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := pattern.FindAllStringSubmatch(input, -1)

	var resultSum int

	for _, match := range matches {
		var nums []int
		for _, value := range match[1:] {
			num, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		resultSum += nums[0] * nums[1]
	}

	fmt.Printf("(Part 1) Sum of all multiplications: %d\n", resultSum)
}
