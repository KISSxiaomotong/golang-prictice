//练习5.8
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Printf("%s", "*")
		}
		fmt.Printf("%s\n", "")
	}
}
