package main

import (
	"fmt"
	"math"
)

//func Solution(n int, a, b []int) bool {
//	for i := 1; i < n; i++ {
//		if hasNext(a, b, i) == false {
//			return false
//		}
//	}
//	return true
//}
//
//func hasNext(a, b []int, n int) bool {
//	for i, v := range a {
//		if v == n && b[i] == n+1 || v == n+1 && b[i] == n {
//			return true
//		}
//	}
//	return false
//}

func Solution(n int, a, b []int) bool {
	result := make(map[int]int)
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
		fmt.Println(b[i])

		if math.Abs(float64(a[i]-b[i])) == 1 {
			key := a[i]
			if a[i] < b[i] {
				key = b[i]
			}
			result[key] = 1
		}
	}
	if len(result) < n-1 {
		return false
	}
	return true
}
