package main

import "fmt"

func main() {
	var stringArray = [5]string{2: "A", 3: "B", 4: "C"}

	for key, value := range stringArray {
		fmt.Printf("Person at %d is %s\n", key, value)
	}
}
