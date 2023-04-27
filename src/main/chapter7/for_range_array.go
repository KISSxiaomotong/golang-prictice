package main

import "fmt"

func main() {
	var arr [5]int

	for i, _ := range arr {
		arr[i] = i * i
	}

	for key, value := range arr {
		fmt.Printf("Array at index %d is %d\n", key, value)
	}
}
