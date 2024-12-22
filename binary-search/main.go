package main

import "fmt"

func main() {
	nums := make([]int, 100) // cria uma slice com 100 posições

	for i := range nums { // percorre o slice
		nums[i] = i + 1 // adiciona a posição atual um valor, no caso o index + 1
	}

	hasNumber := BinarySearch(nums, 1) // chama a função de busca binária e aloca a uma variável

	fmt.Print(hasNumber) // exibe o resultado
}

func BinarySearch(nums []int, item int) bool { // função que faz a busca binária, recebe uma lista e um numero
	low := 0              // primeira posição da lista
	high := len(nums) - 1 // ultima posição da lista

	for low <= high { // esse for funciona igual um while enquanto o low for menor igual a high.
		fmt.Printf("Low: %d \n", low)
		fmt.Printf("High: %d \n", high)
		middle := (low + high) / 2 // posição do item do meio da lista
		guess := nums[middle]      // recebe o valor do item do meio da lista

		if guess == item { // se o chute for igual ao item que está sendo buscado retorna true
			return true
		}
		if guess > item { // se for maior atualiza a variável high (ultima posição) para o valor do meio da lista - 1
			high = middle - 1
		} else { // se for menor atualiza a variável low (primeira posição) para o valor do meio da lista + 1
			low = middle + 1
		}
	}

	return false
}
