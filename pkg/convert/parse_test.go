package convert

import (
	"fmt"
	"testing"
)

func TestRuneToInt(t *testing.T) {
	var tests = []struct {
		char rune
		want int
	}{
		{'0', 0},
		{'1', 1},
		{'2', 2},
		{'3', 3},
		{'4', 4},
		{'5', 5},
		{'6', 6},
		{'7', 7},
		{'8', 8},
		{'9', 9},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("convert %v to %d", tt.char, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans, err := RuneToInt(tt.char)
			if err != nil {
				t.Errorf("Received error: %v", err)
			}
			if ans != tt.want {
				t.Errorf("got %v, wanted %v", ans, tt.want)
			}
		})
	}

}
