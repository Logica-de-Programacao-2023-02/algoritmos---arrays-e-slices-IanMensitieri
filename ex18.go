package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Informe um número inteiro positivo n: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Por favor, informe um número inteiro positivo maior que zero.")
		return
	}

	numerosPrimos := gerarPrimos(n)
	fmt.Printf("Os %d primeiros números primos são: %v\n", n, numerosPrimos)
}

func gerarPrimos(n int) []int {
	numerosPrimos := make([]int, 0, n)
	numeroAtual := 2

	for len(numerosPrimos) < n {
		if ehPrimo(numeroAtual) {
			numerosPrimos = append(numerosPrimos, numeroAtual)
		}
		numeroAtual++
	}

	return numerosPrimos
}

func ehPrimo(numero int) bool {
	if numero < 2 {
		return false
	}
	limite := int(math.Sqrt(float64(numero)))
	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}
