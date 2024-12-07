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
	numbers := parse()
	part1(numbers)
}

func parse() [][]int {
	var numbers [][]int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ": ")
		result, err := strconv.Atoi(splits[0])
		if err != nil {
			panic(err)
		}
		fields := strings.Fields(splits[1])
		operands := []int{result}
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			operands = append(operands, num)
		}
		numbers = append(numbers, operands)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return numbers
}

func part1(numbers [][]int) {
	var totalCalibration int
	var count int
	for _, row := range numbers {
		solvable, equation := findOperand(row[0], row[1], row[2:])
		fmt.Printf("%d = %d %s", row[0], row[1], equation)
		if solvable {
			fmt.Println(" ✅")
			count++
			totalCalibration += row[0]
		} else {
			fmt.Println(" ❌")
		}
	}
	fmt.Printf("(Part 1) total calibration result: %d\n", totalCalibration)
	fmt.Printf("         %d out of %d solvable", count, len(numbers))
}

func findOperand(target, result int, operands []int) (bool, string) {
	if len(operands) == 1 {
		if result+operands[0] == target {
			return true, fmt.Sprintf("+ %d", operands[0])
		}
		if result*operands[0] == target {
			return true, fmt.Sprintf("* %d", operands[0])
		}
		return false, fmt.Sprintf("_ %d", operands[0])
	}
	if solvable, equation := findOperand(target, result+operands[0], operands[1:]); solvable {
		return true, fmt.Sprintf("+ %d %s", operands[0], equation)
	}
	solvable, equation := findOperand(target, result*operands[0], operands[1:])
	if solvable {
		return true, fmt.Sprintf("* %d %s", operands[0], equation)
	}
	return false, fmt.Sprintf("_ %d %s", operands[0], equation)
}
