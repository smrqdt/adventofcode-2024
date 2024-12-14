package main

import (
	"reflect"
	"testing"

	_ "github.com/charmbracelet/log"
)

func Test_solveWithDet(t *testing.T) {
	type args struct {
		A [2][2]float64
		b [2]float64
	}
	tests := []struct {
		name    string
		args    args
		wantX   [2]float64
		wantErr bool
	}{
		{
			"example 1",
			args{[2][2]float64{{94, 22}, {34, 67}}, [2]float64{8400, 5400}},
			[2]float64{80, 40},
			false,
		},
		{
			"skript",
			args{[2][2]float64{{1, 1}, {2, 4}}, [2]float64{150, 440}},
			[2]float64{80, 70},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, err := solveWithDet(tt.args.A, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("solveWithDet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotX, tt.wantX) {
				t.Errorf("solveWithDet() = %v, want %v", gotX, tt.wantX)
			}
		})
	}
}

func Test_gauss(t *testing.T) {
	type args struct {
		A [][]float64
		b []float64
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			"example 1",
			args{[][]float64{{94, 22}, {34, 67}}, []float64{8400, 5400}},
			[]int{80, 40},
			false,
		},
		{
			"skript n=3",
			args{[][]float64{{0, 0, 1}, {4, 2, 1}, {9, 3, 1}}, []float64{15, 15, 0}},
			[]int{-5, 10, 15},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gauss(tt.args.A, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("gauss() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gauss() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
