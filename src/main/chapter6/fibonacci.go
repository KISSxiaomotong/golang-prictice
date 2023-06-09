//递归函数（斐波拉切数列）
package main

import "fmt"

func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
