package main

import "fmt"

func main() {
	fmt.Println(createArray())
	fmt.Println(createPointArray())
}

func createArray() (array any) {
	return new([5]int)
}

func createPointArray() (array any) {
	var arr [5]int
	return arr
}
