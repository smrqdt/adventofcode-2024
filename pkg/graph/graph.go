package graph

import (
	g "github.com/smrqdt/adventofcode-2024/pkg/grid"
	"github.com/smrqdt/adventofcode-2024/pkg/vector"
)

type Graph[T comparable] struct {
	NodesMap map[T][]*Node[T]
	Nodes    []*Node[T]
}

func New[T comparable]() (g Graph[T]) {
	return Graph[T]{
		NodesMap: make(map[T][]*Node[T]),
	}
}

func NewFromGridNeighbors[T comparable](grid g.Grid[*Node[T]], linkFunc func(n1, n2 *Node[T]) bool) (graph Graph[T], err error) {
	for y, row := range grid {
		for x, cell := range row {
			neighs, err := grid.GetNeighbourValues(vector.Vector{X: x, Y: y})
			if err != nil {
				return New[T](), err
			}
			for _, neigh := range neighs {
				cell.Link(neigh)
			}
		}
	}
	return
}

func (g *Graph[T]) Add(value T) (node *Node[T]) {
	node = NewNode(value)
	g.NodesMap[value] = append(g.NodesMap[value], node)
	g.Nodes = append(g.Nodes, node)
	return node
}

func (g *Graph[T]) AddNode(node *Node[T]) {
	g.NodesMap[node.Value] = append(g.NodesMap[node.Value], node)
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph[T]) Find(value T) (node []*Node[T]) {
	return g.NodesMap[value]
}
