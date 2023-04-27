package main

import "fmt"

func main() {
	strSlice := []string{"a", "a", "b", "c", "c", "d", "e", "e", "f", "g", "g"}
	slice := uniqString(strSlice)
	fmt.Println(slice)
}

func uniqString(strSlice []string) (returnSlice []string) {
	slice := make([]string, 0)
	for key, val := range strSlice {
		if (key + 1) == len(strSlice) {
			slice = append(slice, val)
			continue
		}

		if strSlice[key] != strSlice[key+1] {
			slice = append(slice, val)
		}
	}
	return slice
}
