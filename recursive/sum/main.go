package main

import "fmt"

func main() {
	list := []int{2, 4, 6}

	sum := Sum(list)

	fmt.Printf("Soma: %d \n", sum)
}

func Sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return arr[0] + Sum(arr[1:])
}
