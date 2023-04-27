//指针
package main

import "fmt"

func main() {
	var pointA = "pointsA"
	var pointB *string
	pointB = &pointA
	fmt.Printf("%p\n", pointB)
	fmt.Printf("%s\n", *pointB)

	pointA = "this is pointsA"
	fmt.Printf("%s\n", *pointB)
}
