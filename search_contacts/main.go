package main

import (
	"sort"
	"strings"
)

func Solution(a []string, b []string, p string) string {
	var names []string
	for i := 0; i < len(b); i++ {
		if strings.Contains(b[i], p) {
			names = append(names, a[i])
		}
	}
	if len(names) < 1 {
		return "NO CONTACT"
	}
	sort.Strings(names)
	return names[0]
}
