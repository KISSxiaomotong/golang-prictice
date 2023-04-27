package main

import "fmt"

func main() {
	week := map[string]string{"day1": "Monday", "day2": "Tuesday", "day3": "Wednesday", "day4": "Thursday", "day5": "friday", "day6": "Saturday", "day7": "Sunday"}
	for _, value := range week {
		fmt.Println(value)
	}
}
