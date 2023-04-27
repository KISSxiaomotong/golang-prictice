package main

import "fmt"

func main() {
	type Anonymous struct {
		money float32
		int
		string
	}

	anonymous := new(Anonymous)

	anonymous.money = 300000
	anonymous.int = 30
	anonymous.string = "30"

	fmt.Println(anonymous)
}

func function(A float32, B float32) float32 {
	return A * B
}
