package grid

import (
	"reflect"
	"testing"

	"github.com/smrqdt/adventofcode-2024/pkg/convert"
)

func TestConvertGridType(t *testing.T) {
	type args struct {
		g           Grid[rune]
		convertFunc convert.ConvertFunc[rune, int]
	}
	tests := []struct {
		name    string
		args    args
		want    Grid[int]
		wantErr bool
	}{
		{
			"Rune grid to int grid",
			args{Grid[rune]{[]rune("0123456789"), []rune("9876543210")}, convert.RuneToInt},
			Grid[int]{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
			false,
		},
		{
			"Rune grid to int grid with invalid input",
			args{Grid[rune]{[]rune("abcdefg"), []rune("9876543210")}, convert.RuneToInt},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertGridType(tt.args.g, tt.args.convertFunc)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertGridType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertGridType() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
