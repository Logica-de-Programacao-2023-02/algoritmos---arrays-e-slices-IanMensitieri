package main

import (
	"fmt"
	"math"
)

func main() {
	slice := []int{10, 5, 8, 20, 15, 3, 12, 7, 18, 1}

	minimo, maximo := encontrarMinMax(slice)

	fmt.Printf("Valor mínimo no slice: %d\n", minimo)
	fmt.Printf("Valor máximo no slice: %d\n", maximo)
}

func encontrarMinMax(slice []int) (int, int) {
	if len(slice) == 0 {
		return 0, 0
	}

	minimo, maximo := math.MaxInt, math.MinInt

	for _, valor := range slice {
		if valor < minimo {
			minimo = valor
		}
		if valor > maximo {
			maximo = valor
		}
	}

	return minimo, maximo
}
