//递归函数（阶乘）
package main

import "fmt"

func main() {
	result := factorial(10)
	fmt.Printf("%d\n", result)
}

func factorial(num int) (res int) {
	if num == 1 {
		return 1
	}

	return factorial(num-1) * num
}
