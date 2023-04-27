package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sliceB := removeStringSlice(sliceA, 2, 4)

	fmt.Println(sliceB)
}

func removeStringSlice(sliceA []int, start int, end int) (sliceB []int) {
	slice := make([]int, 0)
	for key, val := range sliceA {
		if key < start || key > end {
			slice = append(slice, val)
		}
	}

	return slice
}
