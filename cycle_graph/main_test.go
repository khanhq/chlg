package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1, 3, 2, 4 vs 4, 1, 3, 2",
			args{
				[]int{1, 3, 2, 4},
				[]int{4, 1, 3, 2},
			},
			true,
		},
		{"1, 2, 3, 4 vs 2, 1, 4, 3",
			args{
				[]int{1, 2, 3, 4},
				[]int{2, 1, 4, 3},
			},
			false,
		},
		{"3, 1, 2 vs 2, 3, 1",
			args{
				[]int{3, 1, 2},
				[]int{2, 3, 1},
			},
			true,
		},
		{"1, 2, 1 vs 2, 3, 3",
			args{
				[]int{1, 2, 1},
				[]int{2, 3, 3},
			},
			false,
		},
		{"1, 2, 3, 4 vs 2, 1, 4, 4",
			args{
				[]int{1, 2, 3, 4},
				[]int{2, 1, 4, 4},
			},
			false,
		},
		{"1, 2, 3, 4 vs 2, 1, 4, 3",
			args{
				[]int{1, 2, 3, 4},
				[]int{2, 1, 4, 3},
			},
			false,
		},
		{"1, 2, 2, 3, 3 vs 2, 3, 3, 4, 5",
			args{
				[]int{1, 2, 2, 3, 3},
				[]int{2, 3, 3, 4, 5},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
