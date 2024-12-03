package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	reports := parse()
	part1(reports)
}

func parse() (reports [][]int) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		var report []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return reports
}

func part1(reports [][]int) {
	var safeReportCount int
	for _, report := range reports {
		if isSafe(report) {
			fmt.Printf("%v → safe \n", report)
			safeReportCount++
		} else {
			fmt.Printf("%v → not safe \n", report)
		}
	}

	fmt.Printf("(Part 1) %d reports are safe.\n", safeReportCount)
}

func isSafe(report []int) bool {
	last := report[0]
	increasing := true
	decreasing := true
	for _, level := range report[1:] {
		if last > level {
			increasing = false
		} else {
			decreasing = false
		}
		diff := last - level
		if diff == 0 || diff > 3 || diff < -3 {
			return false
		}
		last = level
	}
	return increasing || decreasing
}
