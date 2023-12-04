package main

import "fmt"

func main() {
	var n int
	fmt.Print("Informe o tamanho dos arrays (n): ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Por favor, informe um valor válido para n.")
		return
	}

	array1 := lerArray("Informe os elementos do primeiro array:", n)
	array2 := lerArray("Informe os elementos do segundo array:", n)

	soma := somarArrays(array1, array2)

	fmt.Println("Resultado da soma dos arrays:", soma)
}

func lerArray(mensagem string, n int) []int {
	fmt.Println(mensagem)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scan(&array[i])
	}
	return array
}

func somarArrays(array1, array2 []int) []int {
	n := len(array1)
	soma := make([]int, n)

	for i := 0; i < n; i++ {
		soma[i] = array1[i] + array2[i]
	}

	return soma
}
