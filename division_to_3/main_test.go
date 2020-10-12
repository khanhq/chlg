package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{

			"23",
			args{
				"23",
			},
			7,
		},
		{

			"0081",
			args{
				"0081",
			},
			11,
		},
		{

			"022",
			args{
				"022",
			},
			9,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Solution(tc.args.s); got != tc.want {
				t.Errorf("Solution() = %d, want = %d", got, tc.want)
			}
		})
	}
}
