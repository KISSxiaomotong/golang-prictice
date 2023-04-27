package main

import "fmt"

func main() {
	bubbling := []int{2, 4, 7, 9, 7, 5, 2, 35, 6, 7, 2, 9}
	bubblingSort(bubbling)
	fmt.Println(bubbling)
}

func bubblingSort(bubbling []int) {
	for pass := 1; pass < len(bubbling); pass++ {
		// one pass:
		for i := 0; i < len(bubbling)-pass; i++ { // the bigger value 'bubbles up' to the last position
			if bubbling[i] > bubbling[i+1] {
				bubbling[i], bubbling[i+1] = bubbling[i+1], bubbling[i]
			}
		}
	}
}
