//练习5.7
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Printf("%s\n", "FizzBuzz")
		case i%3 == 0:
			fmt.Printf("%s\n", "Fizz")
		case i%5 == 0:
			fmt.Printf("%s\n", "Buzz")
		default:
			fmt.Printf("%d\n", i)
		}
	}
}
