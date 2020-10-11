package main


func Solution(a []int, b []int) bool {
	firstVertex := a[0]
	secondVertex := b[0]
	a, b = remove(a, b, 0)
	lastVertex := findNext(a, b, secondVertex)
	if lastVertex == firstVertex {
		return true
	}
	return false
}

func findNext(a []int, b []int, vertex int) int {
	i := findIndex(a, vertex)
	if i == -1 {
		return -1
	}
	if len(a) == 1 {
		return b[0]
	}
	v := b[i]
	a, b = remove(a, b, i)
	return findNext(a, b, v)
}

func findIndex(arr []int, n int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return i
		}
	}
	return -1
}

func remove(a []int, b []int, i int) ([]int, []int) {
	copy(a[i:], a[i+1:])
	copy(b[i:], b[i+1:])
	return a[:len(a)-1], b[:len(b)-1]
}
