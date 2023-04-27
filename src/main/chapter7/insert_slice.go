package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sliceB := insertStringSlice(sliceA, 4, 888)

	fmt.Println(sliceB)
}

func insertStringSlice(sliceA []int, insertIndex int, number int) (sliceB []int) {
	slice := make([]int, 0)
	for key, val := range sliceA {
		slice = append(slice, val)
		if key == insertIndex {
			slice = append(slice, number)
		}
	}
	return slice
}
