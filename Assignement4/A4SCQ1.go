package main

import "fmt"

type Data interface{}

func main() {
	var d Data

	d = "Hello Go"
	if val, ok := d.(string); ok {
		fmt.Println("String value:", val)
	}

	d = 42
	if val, ok := d.(int); ok {
		fmt.Println("Integer value:", val)
	}

	d = 3.14
	if val, ok := d.(float64); ok {
		fmt.Println("Float value:", val)
	}
}
