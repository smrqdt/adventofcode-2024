package convert

import "fmt"

type ConvertFunc[T, R any] func(T) (R, error)

// mapFunc for NewFromInput()
func RuneToInt(r rune) (int, error) {
	if r < '0' || r > '9' {
		return 0, fmt.Errorf("Rune is not a latin numeral")
	}
	return int(r - '0'), nil
}
