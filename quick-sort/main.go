package main

import (
	"fmt"
	"math/rand"
)

func main() {
	list := []int{2, 4, 1, 7, 10, 12, 15, 16, 20, 21, 3}

	orderedList := QuickSort(list)

	fmt.Printf("List Ordered: %v \n", orderedList)
}

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	// Escolhendo um índice aleatório como pivô
	pivotIndex := rand.Intn(len(list))
	pivot := list[pivotIndex]

	// Remover o pivô da lista
	list = append(list[:pivotIndex], list[pivotIndex+1:]...)

	var smaller []int
	var bigger []int

	// Particionando os elementos em relação ao pivô
	for _, number := range list {
		if number <= pivot {
			smaller = append(smaller, number)
		} else {
			bigger = append(bigger, number)
		}
	}

	// Recursão para ordenar as partições
	return append(append(QuickSort(smaller), pivot), QuickSort(bigger)...)
}
