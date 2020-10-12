package main

import (
	"strconv"
)

func main() {

}
func Solution(s string) int {
	result := make(map[string]int)
	for i := 0; i < len(s); i++ {
		for j := 0; j < 10; j++ {
			sum, digits := getSum(s, i, j)
			if sum%3 == 0 {
				j += 2
				result[string(digits)] = 1
			}
		}
	}
	return len(result)
}

func getSum(s string, i, j int) (int, []byte) {
	r := []byte(s)
	r[i] = strconv.Itoa(j)[0]
	total := 0
	for _, d := range r {
		sum, _ := strconv.Atoi(string(d))
		total += sum
	}
	return total, r
}
