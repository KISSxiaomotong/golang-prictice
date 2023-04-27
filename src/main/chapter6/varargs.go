//函数不定长传参
package main

import "fmt"

func main() {
	printStr("a", "b", "c", "d", "e", "f", "g")
}

func printStr(str ...string) {
	for _, v := range str {
		fmt.Printf("%s\n", v)
	}
}
