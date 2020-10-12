package main

import (
	"strconv"
)

func Solution(n int) int {
	num := 0
	for _, v := range strconv.Itoa(n) {
		x, _ := strconv.Atoi(string(v))
		num += x
	}
	sum := num * 2
	if sum > n {
		return sum
	}
	tmp := ""
	for i := 9; i > 0; i-- {
		p := 0
		if len(tmp) > 0 {
			p = -1
		}
		if sum-i > p {
			sum -= i
			tmp = strconv.Itoa(i) + tmp
			i++
		}
	}
	result, _ := strconv.Atoi(tmp)
	return result
}
