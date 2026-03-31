package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string

	fmt.Print("Enter main string: ")
	fmt.Scan(&str2)

	fmt.Print("Enter substring: ")
	fmt.Scan(&str1)

	if strings.Contains(str2, str1) {
		fmt.Println("Substring Found")
	} else {
		fmt.Println("Substring Not Found")
	}
}
