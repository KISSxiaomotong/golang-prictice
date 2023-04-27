package main

import "fmt"

func main() {
	slice := make([]float32, 50, 100)

	for key, _ := range slice {
		slice[key] = 0.87
	}

	num := sumArray(slice)

	fmt.Printf("总数是 %f\n", num)
}

func sumArray(slice []float32) (sum float32) {
	var num float32 = 0
	for _, value := range slice {
		num = num + value
	}

	return num
}
