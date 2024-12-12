package set

import (
	"fmt"
	"iter"
	"maps"
)

type Set[T comparable] map[T]bool

func New[T comparable]() Set[T] {
	return make(Set[T])
}

func (s *Set[T]) Add(value T) {
	(*s)[value] = true
}

func (s *Set[T]) AddAll(values []T) {
	for _, v := range values {
		s.Add(v)
	}
}

func (s *Set[T]) Delete(value T) {
	delete(*s, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, ok := (*s)[value]
	return ok
}

func (s *Set[T]) Clear() {
	clear(*s)
}

func (s *Set[T]) Cleanup() {
	maps.DeleteFunc(*s, func(key T, value bool) bool { return !value })
}

func (s *Set[T]) Join(s2 Set[T]) {
	s2.Cleanup()
	maps.Copy(*s, s2)
}

func (s *Set[T]) Intersect(s2 *Set[T]) Set[T] {
	intersection := New[T]()
	smaller, larger := s, s2
	if s.Cardinality() > s2.Cardinality() {
		smaller, larger = s2, s
	}
	for key := range *smaller {
		if larger.Contains(key) {
			intersection.Add(key)
		}
	}
	return intersection
}

func (s *Set[T]) Cardinality() int {
	return len(*s)
}

func (s *Set[T]) All() iter.Seq[T] {
	return maps.Keys(*s)
}

func (s Set[T]) String() string {
	str := "Set(\n"
	for v := range maps.Keys(s) {
		str += fmt.Sprintf("    %v; \n", v)
	}
	return str + ")"
}
