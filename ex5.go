package main

import (
	"fmt"
)

func main() {
	matriz := make([][]int, 3)
	for i := range matriz {
		matriz[i] = make([]int, 2)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Informe o valor para matriz[%d][%d]: ", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}

	fmt.Println("Matriz Resultante:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d\t", matriz[i][j])
		}
		fmt.Println()
	}
}
