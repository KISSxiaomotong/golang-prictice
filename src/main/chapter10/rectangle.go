package main

import "fmt"

func main() {
	type Rectangle struct {
		length float32
		width  float32
	}

	rectangle := Rectangle{10.8, 3.6}

	fmt.Println("长方形的面积是：", Area(rectangle.length, rectangle.width))
	fmt.Println("长方形的周长是：", Perimeter(rectangle.length, rectangle.width))
}

func Area(length float32, width float32) float32 {
	return length * width
}

func Perimeter(length float32, width float32) float32 {
	return 2 * (length + width)
}
