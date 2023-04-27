package main

import "fmt"

func main() {
	var arr [5]int

	//数组赋值
	for i := 0; i < len(arr); i++ {
		arr[i] = i * i
	}

	//数组输出
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr[i])
	}
}
