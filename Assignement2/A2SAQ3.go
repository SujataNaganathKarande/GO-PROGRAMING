package main

import "fmt"

// function to check palindrome
func isPalindrome(num int) bool {
	original := num
	reverse := 0

	for num != 0 {
		digit := num % 10
		reverse = reverse*10 + digit
		num = num / 10
	}

	return original == reverse
}

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	if isPalindrome(number) {
		fmt.Println("Number is Palindrome")
	} else {
		fmt.Println("Number is Not Palindrome")
	}
}
