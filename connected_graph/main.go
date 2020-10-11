package main

func Solution(n int, a, b []int) bool {
	for i := 1; i < n; i++ {
		if hasNext(a, b, i) == false {
			return false
		}
	}
	return true
}

func hasNext(a, b []int, n int) bool {
	for i, v := range a {
		if v == n && b[i] == n+1 || v == n+1 && b[i] == n {
			return true
		}
	}
	return false
}
