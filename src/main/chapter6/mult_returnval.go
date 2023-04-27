//练习6.1
package main

import "fmt"

func main() {
	add, mul, sub := SumProductDiff(6, 3)
	fmt.Printf("%d\t\t\t", add)
	fmt.Printf("%d\t\t\t", mul)
	fmt.Printf("%d\t\t\t", sub)

	fmt.Printf("%s\n", "")

	addN, mulN, subN := SumProductDiffN(6, 3)
	fmt.Printf("%d\t\t\t", addN)
	fmt.Printf("%d\t\t\t", mulN)
	fmt.Printf("%d\t\t\t", subN)
}

//命名返回值
func SumProductDiff(a int, b int) (add int, mul int, sub int) {
	add = a + b
	mul = a * b
	sub = a - b
	return
}

//非命名返回值
func SumProductDiffN(a int, b int) (int, int, int) {
	return a + b, a * b, a - b
}
