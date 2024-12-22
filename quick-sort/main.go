package main

import "fmt"

func main() {
	list := []int{2, 4, 1, 7, 10}

	orderedList := QuickSort(list)

	fmt.Printf("List Ordered: %v \n", orderedList)
}

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	pivot := list[0]
	var smaller []int
	var bigger []int

	for i := range list[1:] {
		number := list[i+1]
		if number <= pivot {
			smaller = append(smaller, number)
		} else {
			bigger = append(bigger, number)
		}
	}

	return append(append(QuickSort(smaller), pivot), QuickSort(bigger)...)
}
