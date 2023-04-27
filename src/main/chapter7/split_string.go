package main

import "fmt"

func main() {
	str := "HelloWorld"

	strA, strB := divisionStr(str, 5)

	fmt.Println(strA, strB)
}

func divisionStr(str string, index int) (strA string, strB string) {
	strOne := str[:5]
	strTwo := str[5:]
	return strOne, strTwo
}
