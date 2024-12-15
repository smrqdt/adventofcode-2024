package vector

import (
	"math"
	"testing"
)

func TestVector_Angle(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"North", fields{NORTH.X, NORTH.Y}, 0},
		{"North-East", fields{NORTH_EAST.X, NORTH_EAST.Y}, 0.25 * math.Pi},
		{"East", fields{EAST.X, EAST.Y}, 0.5 * math.Pi},
		{"South-East", fields{SOUTH_EAST.X, SOUTH_EAST.Y}, 0.75 * math.Pi},
		{"South", fields{SOUTH.X, SOUTH.Y}, 1 * math.Pi},
		{"South-West", fields{SOUTH_WEST.X, SOUTH_WEST.Y}, 1.25 * math.Pi},
		{"West", fields{WEST.X, WEST.Y}, 1.5 * math.Pi},
		{"North-West", fields{NORTH_WEST.X, NORTH_WEST.Y}, 1.75 * math.Pi},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Angle(); got != tt.want {
				t.Errorf("Vector.Angle() = %#v, want %#v", got/math.Pi, tt.want/math.Pi)
			}
		})
	}
}

func TestVector_Arrow(t *testing.T) {
	type fields struct {
		X int
		Y int
	}
	tests := []struct {
		name   string
		fields fields
		want   rune
	}{
		{"North", fields{NORTH.X, NORTH.Y}, '⬆'},
		{"North-East", fields{NORTH_EAST.X, NORTH_EAST.Y}, '⬈'},
		{"East", fields{EAST.X, EAST.Y}, '➡'},
		{"South-East", fields{SOUTH_EAST.X, SOUTH_EAST.Y}, '⬊'},
		{"South", fields{SOUTH.X, SOUTH.Y}, '⬇'},
		{"South-West", fields{SOUTH_WEST.X, SOUTH_WEST.Y}, '⬋'},
		{"West", fields{WEST.X, WEST.Y}, '⬅'},
		{"North-West", fields{NORTH_WEST.X, NORTH_WEST.Y}, '⬉'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Arrow(); got != tt.want {
				t.Errorf("Vector.Angle() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
