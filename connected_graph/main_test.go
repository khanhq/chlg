package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		n int
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"[1, 2, 4, 4, 3] vs [2, 3, 1, 3, 1]",
			args{
				4,
				[]int{1, 2, 4, 4, 3},
				[]int{2, 3, 1, 3, 1},
			},
			true,
		},
		{
			"[1, 2, 1, 3] vs [2, 4, 3, 4]",
			args{
				4,
				[]int{1, 2, 1, 3},
				[]int{2, 4, 3, 4},
			},
			false,
		},
		{
			"[2, 4, 5, 3] vs [3, 5, 6, 4]",
			args{
				6,
				[]int{2, 4, 5, 3},
				[]int{3, 5, 6, 4},
			},
			false,
		},
		{
			"[1, 3] vs [2, 2]",
			args{
				3,
				[]int{1, 3},
				[]int{2, 2},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.n, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
