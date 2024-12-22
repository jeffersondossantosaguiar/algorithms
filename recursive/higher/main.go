package main

import "fmt"

func main() {
	list := []int{1, 3, 7, 0, 2}

	higherNumber := Higher(list)

	fmt.Printf("Maior numero: %d \n", higherNumber)
}

func Higher(arr []int) int {
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		} else {
			return arr[1]
		}
	}
	subMax := Higher(arr[1:])

	if arr[0] > subMax {
		return arr[0]
	} else {
		return subMax
	}
}
