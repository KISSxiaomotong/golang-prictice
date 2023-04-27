//练习6.2
package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := MySqrt(-6.6)
	resultN, errN := MySqrtN(6.6)

	fmt.Printf("%f\n", result)
	fmt.Printf("%s\n", err)
	fmt.Printf("%f\n", resultN)
	fmt.Printf("error:", errN)
}

func MySqrt(num float64) (sqrt float64, err error) {
	if num < 0 {
		return math.NaN(), errors.New("不合法的传参！")
	}

	return math.Sqrt(num), nil
}

func MySqrtN(num float64) (float64, error) {
	if num < 0 {
		return math.NaN(), errors.New("不合法的传参！")
	}

	return math.Sqrt(num), nil
}
