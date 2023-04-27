//for循环练习5.5
package main

import "fmt"

func main() {
	var str string = ""
	for i := 0; i < 25; i++ {
		str = str + "G"
		fmt.Printf("%s\n", str)
	}

	fmt.Printf("%s\n\n\n", "---------换行符-------------")

	for m := 0; m < 25; m++ {
		for n := 0; n < m; n++ {
			fmt.Printf("%s", "G")
		}
		fmt.Printf("%s\n", "")
	}
}
