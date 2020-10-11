package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"3, 2, 1, 1, 2, 3, 1",
			args{
				[]int{3, 2, 1, 1, 2, 3, 1},
			},
			5,
		},
		{
			"4, 1, 4, 1",
			args{
				[]int{4, 1, 4, 1},
			},
			6,
		},
		{
			"3, 3, 3",
			args{
				[]int{3, 3, 3},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.a); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
