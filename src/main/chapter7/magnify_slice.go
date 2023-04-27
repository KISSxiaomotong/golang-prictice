package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}

	copySlice := enlargeSlice(slice, 5)

	for i, value := range copySlice {
		fmt.Printf("Array at index %d is %d\n", i, value)
	}
}

func enlargeSlice(slice []int, factor int) (result []int) {
	sliceA := make([]int, len(slice)*factor)
	copy(sliceA, slice)
	return sliceA
}
