package main

import "fmt"

func main() {
	slice := make([]int, 3, 10)
	for key, value := range slice {
		value = key + value
	}

	for i, value := range slice {
		fmt.Printf("Slice at %d is %d\n", i, value)
	}

	array := [5]int{1, 2, 3, 4, 5}
	for key, value := range array {
		value = key + value
	}

	for i, value := range array {
		fmt.Printf("Array at %d is %d\n", i, value)
	}
}
