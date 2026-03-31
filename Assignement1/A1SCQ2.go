package main

import "fmt"

func main() {
	var str1, str2 string

	fmt.Print("Enter first string: ")
	fmt.Scan(&str1)

	fmt.Print("Enter second string: ")
	fmt.Scan(&str2)

	if str1 == str2 {
		fmt.Println("Strings are Equal")
	} else {
		fmt.Println("Strings are Not Equal")
	}
}
