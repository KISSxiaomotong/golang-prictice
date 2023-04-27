package main

import "fmt"

func main() {
	numberSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := mapFunc(multiplyTen, numberSlice)
	fmt.Println(slice)
}

func multiplyTen(num int) int {
	return num * 10
}

func mapFunc(multiplyTen func(int) int, slice []int) []int {
	multiplySlice := make([]int, len(slice))
	for key, val := range slice {
		multiplySlice[key] = multiplyTen(val)
	}
	return multiplySlice
}
