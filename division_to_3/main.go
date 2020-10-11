package main

import (
	"strconv"
	"strings"
)

func Solution(s string) int {
	var result []int
	arr1 := ToArray(s)
	for i := 0; i < len(arr1); i++ {
		arr := make([]string, len(arr1))
		copy(arr, arr1)
		for j := 0; j < 10; j++ {
			arr[i] = strconv.Itoa(j)
			if n,_ := strconv.Atoi(strings.Join(arr, "")) ; n%3 == 0 && isExist(result, n) == false {
				result = append(result, n)
			}
		}
	}
	return len(result)
}

func ToArray(s string)[]string{
	var result []string
	for _, r := range s{
		result = append(result, string(r))
	}
	return  result
}

func isExist(arr []int, n int) bool{
	for _, v := range arr{
		if n == v{
			return true
		}
	}
	return false
}