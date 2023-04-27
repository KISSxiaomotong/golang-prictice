package main

import "fmt"

func main() {
	slice := make([]int, 10, 15)
	for key, _ := range slice {
		slice[key] = key * 5
	}

	for i, value := range slice {
		fmt.Printf("Slice at %d is %d\n", i, value)
	}

}
