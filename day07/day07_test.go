package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkConcatByMath(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ConcatInt(r.Intn(999999999), r.Intn(999999999))
	}
}

func BenchmarkConcatByString(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ConcatIntByString(r.Intn(999999999), r.Intn(999999999))
	}
}

func TestConcatInt(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{15, 6, 156},
		{6, 15, 615},
		{123, 456, 123456},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("ConcatInt(%d, %d) → %d", tt.a, tt.b, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := ConcatInt(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("go %d, wanted %d", ans, tt.want)
			}
		})
	}
}
