package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{50, 20, 40, 10, 30}

	fmt.Println("Before Sorting:", arr)

	sort.Ints(arr)

	fmt.Println("After Sorting (Ascending):", arr)
}
