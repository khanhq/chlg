package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"10", 10, 11},
		{"12", 12, 15},
		{"14", 14, 19},
		{"99", 99, 9999},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.n); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
