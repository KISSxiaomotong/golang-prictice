//控制结构
package main

import "fmt"

func main() {
	//找出三个数中最大的那个
	var a int = 9
	var b int = 15
	var c int = 10
	var max int

	fmt.Printf("三个数中最大的数字是：")

	if a > b {
		max = a
	} else {
		max = b
	}

	if max > c {
		fmt.Printf("%d\n", max)
	} else {
		fmt.Printf("%d\n", c)
	}
}
