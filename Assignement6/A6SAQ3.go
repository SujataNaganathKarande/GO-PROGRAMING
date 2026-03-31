package main

import (
	"Assignement6/rectangle"
	"fmt"
)

func main() {

	var l, w float64

	fmt.Println("Enter length and width of rectangle:")
	fmt.Scan(&l, &w)

	area := rectangle.Area(l, w)

	fmt.Println("Area of Rectangle:", area)
}
