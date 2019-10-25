package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, p string) int {
	for i, v := range vs {
		if v == p {
			return i
		}
	}
	return -1
}

func Include(vs []string, p string) bool {
	if Index(vs, p) >= 0 {
		return true
	}
	return false
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	var result []string
	for _, v := range vs {
		if f(v) {
			result = append(result, v)	
		}
	}
	return result
}

func Map(vs []string, f func(string) string) []string {
	result := make([]string, len(vs))
	for i, v := range vs {
		result[i] = f(v)
	}
	return result
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(p string) bool {
		return strings.HasPrefix(p, "p")
	}))
	fmt.Println(All(strs, func(p string) bool {
		return strings.HasPrefix(p, "p")
	}))
	fmt.Println(Filter(strs, func(p string) bool {
		return strings.Contains(p, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))
}
