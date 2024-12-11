package graph

import (
	"fmt"
	"iter"
	"maps"
)

var lastNodeID int

type Node[T comparable] struct {
	ID            int
	Value         T
	edges         map[*Node[T]]bool
	incomingEdges map[*Node[T]]bool
}

func NewNode[T comparable](value T) *Node[T] {
	lastNodeID++
	n := &Node[T]{ID: lastNodeID}
	n.Value = value
	n.edges = make(map[*Node[T]]bool)
	n.incomingEdges = make(map[*Node[T]]bool)
	return n
}

func (n *Node[T]) String() (str string) {
	return n.ExtendedString(0)
}

func (n *Node[T]) ExtendedString(verbosity int) (str string) {
	switch {
	case verbosity > 2:
		str = fmt.Sprintf("[%v]← %v", n.incomingEdges, str)
		fallthrough
	case verbosity > 1:
		str = fmt.Sprintf("[%v]→ %v", n.edges, str)
		fallthrough
	default:
		str = fmt.Sprintf("Node%d(%v) %v", n.ID, n.Value, str)
	}
	return str
}

func (n *Node[T]) Link(n2 *Node[T]) bool {
	if n.edges[n2] {
		return false
	}
	n.edges[n2] = true
	n2.incomingEdges[n] = true
	return true
}

func (n *Node[T]) Unlink(n2 *Node[T]) bool {
	if !n.edges[n2] {
		return false
	}
	delete(n.edges, n2)
	delete(n2.edges, n)
	return true
}

func (n *Node[T]) Edges() iter.Seq[*Node[T]] {
	return maps.Keys(n.edges)
}

func (n *Node[T]) InEdges() iter.Seq[*Node[T]] {
	return maps.Keys(n.incomingEdges)
}

func (n *Node[T]) AllEdges() iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		for key := range maps.Keys(n.edges) {
			if !yield(key) {
				return
			}
		}
		for key := range maps.Keys(n.incomingEdges) {
			if !yield(key) {
				return
			}
		}
	}
}
