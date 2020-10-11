package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		a []string
		b []string
		p string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"pim, pom",
			args{
				[]string{"pim", "pom"},
				[]string{"999999999", "777888999"},
				"8899",
			},
			"pom",
		},
		{"sander, amy, ann, michael",
			args{
				[]string{"sander", "amy", "ann", "michael"},
				[]string{"123456789", "234567890", "789123456", "123123123"},
				"1",
			},
			"ann",
		},
		{"adam, eva, leo",
			args{
				[]string{"adam", "eva", "leo"},
				[]string{"121212121", "111111111", "444555666"},
				"112",
			},
			"NO CONTACT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.a, tt.args.b, tt.args.p); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
