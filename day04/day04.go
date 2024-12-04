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
