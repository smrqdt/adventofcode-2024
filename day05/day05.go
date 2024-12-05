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
	rules, updates := parse()
	part1(rules, updates)
}

func parse() ([]PageOrderRule, [][]int) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var rules []PageOrderRule

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		values := strings.Split(line, "|")
		before, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		after, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		rules = append(rules, PageOrderRule{before, after})
	}

	var updates [][]int

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ",")
		var update []int
		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			update = append(update, num)
		}
		updates = append(updates, update)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return rules, updates
}

func part1(rules []PageOrderRule, updates [][]int) {
	var middlePageNumerSum int

	for _, update := range updates {
		if checkUpdate(rules, update) {
			middlePageNumerSum += update[len(update)/2]
		}
	}

	fmt.Printf("(Part 1) Sum of middle page numbers from correctly-ordered updates: %d\n", middlePageNumerSum)
}

func checkUpdate(rules []PageOrderRule, update []int) bool {
	for _, rule := range rules {
		if !rule.Check(update) {
			return false
		}
	}
	return true
}

type PageOrderRule struct {
	before, after int
}

func (r PageOrderRule) Check(update []int) bool {
	beforeIdx := slices.Index(update, r.before)
	afterIdx := slices.Index(update, r.after)
	if beforeIdx == -1 || afterIdx == -1 {
		return true
	}

	return beforeIdx < afterIdx
}
