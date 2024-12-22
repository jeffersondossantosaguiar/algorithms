package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5}

	count := Count(list)

	fmt.Printf("Total de items: %d \n", count)
}

func Count(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + Count(arr[1:])
}
