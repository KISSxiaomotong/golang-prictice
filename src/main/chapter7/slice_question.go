package main

import "fmt"

func main() {
	a := make([]byte, 5)
	fmt.Printf("切片a的len是%d\n", len(a))
	fmt.Printf("切片a的cap是%d\n", cap(a))

	b := []byte{'p', 'o', 'e', 'm'}
	c := b[2:]
	fmt.Printf("切片c的len是%d\n", len(c))
	fmt.Printf("切片c的cap是%d\n", cap(c))
}
