package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/smrqdt/adventofcode-2024/pkg/helpers"
	v "github.com/smrqdt/adventofcode-2024/pkg/vector"
)

//go:embed input
var input string
var NotSolvableError error = fmt.Errorf("System of linear equations is not solvable")

func main() {
	log.SetLevel(log.DebugLevel)

	arcades := parse(input)
	_ = part1(arcades)
	_ = part2(arcades)
}

type Arcade struct {
	BtnA, BtnB v.Vector
	Prize      v.Vector
}

func parse(input string) []Arcade {
	defer helpers.TrackTime(time.Now(), "parse()")

	scanner := bufio.NewScanner(strings.NewReader(input))

	btnRe := regexp.MustCompile(`Button (A|B): X\+(\d+), Y\+(\d+)`)
	prizeRe := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	var arcades []Arcade
	for scanner.Scan() {
		arcade := Arcade{}
		matches := btnRe.FindStringSubmatch(scanner.Text())
		num, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		arcade.BtnA.X = num

		num, err = strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}
		arcade.BtnA.Y = num

		if !scanner.Scan() {
			panic("Scanner exhausted")
		}
		matches = btnRe.FindStringSubmatch(scanner.Text())
		num, err = strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		arcade.BtnB.X = num

		num, err = strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}
		arcade.BtnB.Y = num

		if !scanner.Scan() {
			panic("Scanner exhausted")
		}
		matches = prizeRe.FindStringSubmatch(scanner.Text())
		num, err = strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		arcade.Prize.X = num

		num, err = strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		arcade.Prize.Y = num

		arcades = append(arcades, arcade)

		scanner.Scan()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return arcades
}

func part1(arcades []Arcade) int {
	defer helpers.TrackTime(time.Now(), "part1()")

	var tokensSpent int
	for _, arcade := range arcades {
		result, err := gauss(
			[][]float64{
				{float64(arcade.BtnA.X), float64(arcade.BtnB.X)},
				{float64(arcade.BtnA.Y), float64(arcade.BtnB.Y)},
			},
			[]float64{float64(arcade.Prize.X), float64(arcade.Prize.Y)},
		)
		if errors.Is(err, NotSolvableError) {
			continue
		}
		if err != nil {
			panic(err)
		}
		tokensSpent += result[0]*3 + result[1]
	}

	log.Warnf("(Part 1) %d Tokens spent to win all Prizes \n", tokensSpent)
	return tokensSpent
}

func part2(arcades []Arcade) int {
	defer helpers.TrackTime(time.Now(), "part2()")

	var tokensSpent int
Loop:
	for _, arcade := range arcades {
		A := [2][2]float64{
			{float64(arcade.BtnA.X), float64(arcade.BtnB.X)},
			{float64(arcade.BtnA.Y), float64(arcade.BtnB.Y)},
		}
		b := [2]float64{float64(arcade.Prize.X + 10000000000000), float64(arcade.Prize.Y + 10000000000000)}
		result, err := solveWithDet(A, b)
		if errors.Is(err, NotSolvableError) {
			continue
		}
		if err != nil {
			panic(err)
		}
		var resultInt [2]int
		for i, value := range result {
			if math.Abs(math.Round(value)-value) > 1e-9 {
				log.Info("No solution in ℤ", "A", A, "b", b, "x", "result", result)
				continue Loop
			}
			resultInt[i] = int(math.Round(value))
		}
		log.Info("Solved", "A", A, "b", b, "x", result)
		tokensSpent += resultInt[0]*3 + resultInt[1]
	}

	log.Warnf("(Part 2) %d Tokens spent to win all Prizes \n", tokensSpent)
	return 0
}

func solveWithDet(A [2][2]float64, b [2]float64) (x [2]float64, err error) {
	det := A[0][0]*A[1][1] - A[0][1]*A[1][0]
	if det == 0 {
		return x, fmt.Errorf("det == 0: %w", NotSolvableError)
	}

	var inv [2][2]big.Rat

	detRat := (&big.Rat{}).SetFloat64(det)
	detRat.Inv(detRat)
	inv[0][1].Mul(inv[0][1].SetFloat64(-A[0][1]), detRat)
	inv[1][0].Mul(inv[1][0].SetFloat64(-A[1][0]), detRat)
	inv[0][0].Mul(inv[0][0].SetFloat64(A[1][1]), detRat)
	inv[1][1].Mul(inv[1][1].SetFloat64(A[0][0]), detRat)

	var xRat [2]big.Rat

	for i := range A {
		for j := range A[0] {
			xRat[i].Add(
				&xRat[i],
				inv[i][j].Mul(
					&inv[i][j],
					big.NewRat(int64(b[j]), 1),
				),
			)
		}
		x[i], _ = xRat[i].Float64()
	}

	return x, nil
}

func gauss(A [][]float64, b []float64) (x []int, err error) {
	log.Debug("Before ", "A", strings.ReplaceAll(fmt.Sprintf("%v", A), "] [", "]\n ["), "b", b)
	if len(A) != len(A[0]) {
		return nil, fmt.Errorf("Only square matrices supported")
	}
	if len(A) != len(b) {
		return nil, fmt.Errorf("Dimensions of A and b must match")
	}

	for i := range A {
		if A[i][i] == 0 {
			for row := i + 1; row < len(A); row++ {
				if math.Abs(A[row][i]) > math.Abs(A[i][i]) {
					A[row], A[i] = A[i], A[row]
					b[row], b[i] = b[i], b[row]
					break
				}
			}
		}
	}

	log.Debug("After Pivot", "A", strings.ReplaceAll(fmt.Sprintf("%v", A), "] [", "]\n ["), "b", b)

	for j := range A[0] {
		for i := range A[j:] {
			i = i + j
			switch {
			case i == j:
				divisor := A[i][i]
				if divisor == 1 {
					continue
				}
				for col := range A[i] {
					A[i][col] /= divisor
				}
				b[i] /= divisor
			case i > j:
				multiplier := A[i][j]
				if multiplier == 0 {
					continue
				}
				for col := range A[i] {
					A[i][col] -= multiplier * A[j][col]
				}
				b[i] -= multiplier * b[j]
			default:
				log.Error("Somehow ended in default case", "j", j, "i", i)
				panic("whut")
			}
		}
	}

	log.Debug("After Forward Elimination", "A", strings.ReplaceAll(fmt.Sprintf("%v", A), "] [", "]\n ["), "b", b)

	for i := range A {
		var allZero bool
		for _, val := range A[i] {
			if val != 0 {
				allZero = false
			}
		}
		if allZero && b[i] != 0 {
			return nil, fmt.Errorf("There stepless rows with b_i != 0: %w", NotSolvableError)
		} else if allZero {
			A[i] = nil
		}
	}

	for j := range slices.Backward(A[0]) {
		if j == 0 || A[j] == nil {
			continue
		}
		for i := range slices.Backward(A[j][:j]) {
			log.Debug("Rev El", "i", i, "j", j)
			multiplier := A[i][j]
			for col := range A[i] {
				A[i][col] -= multiplier * A[j][col]
			}
			b[i] -= multiplier * b[j]
		}
	}

	log.Debug("After Backwards Elimination", "A", strings.ReplaceAll(fmt.Sprintf("%v", A), "] [", "]\n ["), "b", b)

	for _, value := range b {
		if math.Abs(math.Round(value)-value) < 1e-9 {
			x = append(x, int(math.Round(value)))
		} else {
			return nil, fmt.Errorf("No solution in ℤ, result is %v: %w", b, NotSolvableError)
		}
	}

	return x, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
