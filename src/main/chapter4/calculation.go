//运算符
package main

import "fmt"

func main() {
	var a float32 = 1.6
	var b float32 = 3.2

	var c = a + b
	var d = a * b
	var e = b / a

	fmt.Println(c)
	fmt.Println(fmt.Sprintf("%.2f", d))
	fmt.Println(e)
}
