package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var input string

func main() {
	reports := parse()
	count := solve(reports, false)
	fmt.Printf("(Part 1) %d reports are safe.\n", count)
	count = solve(reports, true)
	fmt.Printf("(Part 2) %d reports are safe after dampening.\n", count)
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

func solve(reports [][]int, dampen bool) int {
	var safeReportCount int
	for _, report := range reports {
		if isSafe(report, dampen) {
			safeReportCount++
		}
	}
	return safeReportCount
}

func isSafe(report []int, dampen bool) bool {
	monotonic := isMonotonic(report)
	maxDiff := isMaxDiff(report)
	if monotonic && maxDiff {
		// fmt.Println(report, "→ safe")
		return true
	}
	if dampen {
		// fmt.Println(report, "→ unsafe (dampen)")
		for i := range report {
			newReport := slices.Clone(report)
			newReport = slices.Delete(newReport, i, i+1)
			// fmt.Printf("(%d) ", i)
			if isSafe(newReport, false) {
				return true
			}
		}
	}
	// fmt.Println(report, "→ unsafe")
	return false
}

const (
	UNKNOWN = iota
	INCREASING
	DECREASING
)

func isMonotonic(report []int) bool {
	monotonic := UNKNOWN
	var last int
	for i, level := range report {
		if i == 0 {
			last = level
			continue
		}
		if level > last {
			if monotonic == DECREASING {
				return false
			}
			monotonic = INCREASING
		}
		if level < last {
			if monotonic == INCREASING {
				return false
			}
			monotonic = DECREASING
		}
		last = level
	}
	return true
}

func isMaxDiff(report []int) bool {
	var last int
	for i, level := range report {
		if i == 0 {
			last = level
			continue
		}
		diff := last - level
		if diff < -3 || diff == 0 || diff > 3 {
			return false
		}
		last = level
	}
	return true
}
