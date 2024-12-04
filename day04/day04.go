package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"iter"
	"slices"
	"strings"
)

//go:embed input
var input string

func main() {
	wordSalad := parse()
	part1(wordSalad)
	part2(wordSalad)
}

func parse() wordGrid {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var wordSalad wordGrid
	for scanner.Scan() {
		line := scanner.Text()
		wordSalad = append(wordSalad, []byte(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return wordSalad
}

func part1(wordSalad wordGrid) {
	var count int

	// horizontal
	for _, line := range wordSalad {
		count += bytes.Count(line, []byte("XMAS"))
		count += bytes.Count(line, []byte("SAMX"))
	}

	for line := range wordSalad.Vertically() {
		count += bytes.Count(line, []byte("XMAS"))
		count += bytes.Count(line, []byte("SAMX"))
	}

	for line := range wordSalad.Diagonally() {
		// fmt.Printf("↘ %s\n", line)
		count += bytes.Count(line, []byte("XMAS"))
		count += bytes.Count(line, []byte("SAMX"))
	}

	for line := range wordSalad.DiagonallyReverse() {
		// fmt.Printf("↗ %s\n", line)
		count += bytes.Count(line, []byte("XMAS"))
		count += bytes.Count(line, []byte("SAMX"))
	}

	fmt.Printf("(Part 1) XMAS found %d times \n", count)
}

type wordGrid [][]byte

func (w wordGrid) Vertically() iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		for i := range w[0] {
			var col []byte
			for _, row := range w {
				col = append(col, row[i])
			}
			if !yield(col) {
				return
			}
		}
	}
}

func (w wordGrid) Diagonally() iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		for i := range w[0] {
			var col []byte
			for j, row := range w {
				if i+j == len(row) {
					if !yield(col) {
						return
					}
					col = nil
				}
				col = append(col, row[(i+j)%len(row)])
			}
			if len(col) > 0 {
				if !yield(col) {
					return
				}
			}
		}
	}
}

func (w wordGrid) DiagonallyReverse() iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		for i := range w[0] {
			var col []byte
			for j, row := range slices.Backward(w) {
				if i+(len(w)-j-1) == len(row) {
					if !yield(col) {
						return
					}
					col = nil
				}
				col = append(col, row[(i+len(w)-j-1)%len(row)])
			}
			if len(col) > 0 {
				if !yield(col) {
					return
				}
			}
		}
	}
}

func part2(wordSalad wordGrid) {
	var count int

	// horizontal
	for line := range wordSalad.Xly() {
		count += bytes.Count(line, []byte("MASMAS"))
		count += bytes.Count(line, []byte("MASSAM"))
		count += bytes.Count(line, []byte("SAMSAM"))
		count += bytes.Count(line, []byte("SAMMAS"))
	}

	fmt.Printf("(Part 2) X-MAS found %d times \n", count)
}

func (w wordGrid) Xly() iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		for y := range w[:len(w)-2] {
			for x := range w[0][:len(w[0])-2] {
				value := []byte{w[y][x], w[y+1][x+1], w[y+2][x+2], w[y+2][x], w[y+1][x+1], w[y][x+2]}
				if !yield(value) {
					return
				}
			}
		}
	}
}
