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
	fmt.Printf("\n(Part 1) %d reports are safe.\n\n", count)
	count = solve(reports, true)
	fmt.Printf("\n(Part 2) %d reports are safe after dampening.\n\n", count)
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
	fmt.Print(report, " → ")
	monotonic, failedAt := isMonotonic(report)
	var maxDiff bool
	if monotonic {
		fmt.Print("(M) ")
		maxDiff, failedAt = isMaxDiff(report)
		if maxDiff {
			fmt.Print("(D) ")
		}
	}
	if !monotonic || !maxDiff {
		if dampen {
			fmt.Printf("→ damped bei removing %d (%d) → ", failedAt, report[failedAt])
			report = slices.Delete(report, failedAt, failedAt+1)
			return isSafe(report, false)
		}
		fmt.Println("→ unsafe")
		return false
	}
	fmt.Println("→ safe")
	return true
}

const (
	UNKNOWN = iota
	INCREASING
	DECREASING
)

func isMonotonic(report []int) (bool, int) {
	monotonic := UNKNOWN
	var last int
	for i, level := range report {
		if i == 0 {
			last = level
			continue
		}
		if level > last {
			if monotonic == DECREASING {
				return false, i
			}
			monotonic = INCREASING
		}
		if level < last {
			if monotonic == INCREASING {
				return false, i
			}
			monotonic = DECREASING
		}
		last = level
	}
	return true, 0
}

func isMaxDiff(report []int) (bool, int) {
	var last int
	for i, level := range report {
		if i == 0 {
			last = level
			continue
		}
		diff := last - level
		if diff < -3 || diff == 0 || diff > 3 {
			return false, i
		}
		last = level
	}
	return true, 0
}
