package main

import "fmt"

func main() {
	var array = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliceA := array[5:7]
	sliceB := array[0:4]
	for i, value := range sliceA {
		fmt.Printf("SliceA at %d is %d\n", i, value)
	}

	for j, value := range sliceB {
		fmt.Printf("SliceB at %d is %d\n", j, value)
	}
}
