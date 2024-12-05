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
	updates, rules := parse()
	solve(updates, rules)
}

func parse() ([][]int, []PageOrderRule) {
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

	return updates, rules
}

func solve(updates [][]int, rules []PageOrderRule) {
	var part1Sum, part2Sum int

	for _, update := range updates {
		if checkUpdate(update, rules) {
			part1Sum += update[len(update)/2]
		} else {
			// fmt.Println(update, "")
			fixedUpdate := fixUpdate(update, rules)
			if !checkUpdate(fixedUpdate, rules) {
				panic(fixedUpdate)
			}
			// fmt.Println(" → ", fixedUpdate)
			part2Sum += fixedUpdate[len(fixedUpdate)/2]
		}
	}

	fmt.Printf("(Part 1) Sum of middle page numbers from correctly-ordered updates: %d\n", part1Sum)
	fmt.Printf("(Part 2) Sum of middle page numbers from repaired updates: %d\n", part2Sum)
}

func checkUpdate(update []int, rules []PageOrderRule) bool {
	for _, rule := range rules {
		if !rule.Check(update) {
			return false
		}
	}
	return true
}

func fixUpdate(update []int, rules []PageOrderRule) []int {
	for !checkUpdate(update, rules) {
		for _, rule := range rules {
			if !rule.Check(update) {
				update = rule.Fix(update)
			}
		}
	}
	return update
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

func (r PageOrderRule) Fix(update []int) []int {
	beforeIdx := slices.Index(update, r.before)
	afterIdx := slices.Index(update, r.after)

	update = slices.Delete(update, beforeIdx, beforeIdx+1)
	update = slices.Insert(update, afterIdx, r.before)
	// fmt.Println(update, " Fixed: ", r.before, r.after)
	return update
}
