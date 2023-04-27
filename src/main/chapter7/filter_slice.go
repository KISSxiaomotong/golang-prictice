package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 89, 10}
	result := Filter(s, even)

	for i, value := range result {
		fmt.Printf("Slice at index %d is %d\n", i, value)
	}
}

func Filter(s []int, fn func(n int) (k bool)) []int {
	var slice []int
	for _, value := range s {
		if fn(value) {
			slice = append(slice, value)
		}
	}
	return slice
}

func even(n int) (k bool) {
	if n%2 == 0 {
		return true
	}
	return false
}
