package grid

import (
	"fmt"

	"github.com/smrqdt/adventofcode-2024/pkg/convert"
)

func ConvertGridType[T, V comparable](g Grid[T], convertFunc convert.ConvertFunc[T, V]) (Grid[V], error) {
	newGrid := New[V](g.Count())
	for vec, value := range g.All() {
		converted, err := convertFunc(value)
		if err != nil {
			return nil, fmt.Errorf("Could not convert value %#v to %T: %w", value, converted, err)
		}
		err = newGrid.SetValue(vec, converted)
		if err != nil {
			return nil, fmt.Errorf("Could not set %v to %v: %w", vec, converted, err)
		}
	}
	return newGrid, nil
}
