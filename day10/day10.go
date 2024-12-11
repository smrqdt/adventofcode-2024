package main

import (
	_ "embed"
	"fmt"
	"log"
	"slices"

	"github.com/smrqdt/adventofcode-2024/pkg/convert"
	gr "github.com/smrqdt/adventofcode-2024/pkg/graph"
	g "github.com/smrqdt/adventofcode-2024/pkg/grid"
	"github.com/smrqdt/adventofcode-2024/pkg/set"
)

//go:embed input
var input string

func main() {
	graph := parse()
	solve(graph)
}

func parse() (graph gr.Graph[int]) {
	graph = gr.New[int]()
	grid, err := g.NewFromInput(input, func(r rune) (*gr.Node[int], error) {
		num, err := convert.RuneToInt(r)
		if err != nil {
			return nil, err
		}
		return graph.Add(num), nil
	})

	for vec, node := range grid.All() {
		neighs, err := grid.GetNeighbourValues(vec)
		if err != nil {
			panic(err)
		}
		for _, neigh := range neighs {
			if neigh.Value == node.Value+1 {
				node.Link(neigh)
			}
		}
	}
	if err != nil {
		panic(err)
	}
	return graph

}

func solve(graph gr.Graph[int]) {
	startNodes := graph.Find(0)
	var paths [][]*gr.Node[int]
	var score int
	for _, node := range startNodes {
		newPaths, found := findPathsToValue(node, 9, nil)
		if len(newPaths) > 0 && !found {
			log.Fatalf("Found false but %d new paths", len(newPaths))
		}
		paths = append(paths, newPaths...)
		distinctTargets := set.New[*gr.Node[int]]()
		for _, path := range newPaths {
			distinctTargets.Add(path[9])
		}
		newScore := distinctTargets.Cardinality()
		score += newScore
		// fmt.Println(node, newScore, len(newPaths), distinctTargets, score)
	}
	fmt.Printf("(Part 1) Sum of trailhead scores: %d \n", score)
	fmt.Printf("(Part 2) Sum of trailhead trails: %d \n", len(paths))
}

func findPathsToValue(node *gr.Node[int], target int, path []*gr.Node[int]) (paths [][]*gr.Node[int], found bool) {
	path = append(path, node)
	if node.Value == target {
		return append(paths, path), true
	}
	var foundAny bool
	for edge := range node.Edges() {
		edgePaths, found := findPathsToValue(edge, target, slices.Clone(path))
		if found {
			paths = append(paths, edgePaths...)
			foundAny = true
		}
	}
	return paths, foundAny
}
